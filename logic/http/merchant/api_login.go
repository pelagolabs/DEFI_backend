package merchant

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/multiformats/go-multibase"
	"log"
	"math/rand"
	"net/http"
	"time"
	"veric-backend/internal/util"
	"veric-backend/logic/blockchain/did"
	"veric-backend/logic/blockchain/eth"
	"veric-backend/logic/config"
	"veric-backend/logic/db"
	"veric-backend/logic/http/http_util"
)

var (
	userSignContentSet = new(util.SyncedMap[string, string])
)

func GetMerchantSignContent(typ interface{}, r *http_util.HTTPContext) (resp interface{}, respErr error) {
	address := eth.ToEthAddress(r.QueryWithDefault("address", ""))

	if !eth.IsValidEthAddress(address) {
		return nil, http_util.NewHttpError(http.StatusBadRequest, "eth address is invalid.")
	}

	nonce := rand.Intn(899999) + 100000
	content := fmt.Sprintf("Welcome to Airswift merchant platform. nonce: %d", nonce)
	userSignContentSet.Store(address, content)

	return &merchantSignContentResponse{
		Content: content,
	}, nil
}

func GetUserRelatedMerchant(typ interface{}, r *http_util.HTTPContext) (resp interface{}, respErr error) {
	address := eth.ToEthAddress(r.QueryWithDefault("address", ""))

	if !eth.IsValidEthAddress(address) {
		return nil, http_util.NewHttpError(http.StatusBadRequest, "eth address is invalid.")
	}

	user, userErr := db.FindUserByAddress(address)
	if userErr != nil || user.ID == 0 {
		return nil, http_util.NewHttpError(http.StatusNotFound, "user not found.")
	}

	relations, relationErr := db.FindAvailableMerchantUserWithMerchantByUser(user.ID)
	if relationErr != nil {
		return nil, http_util.NewHttpError(http.StatusNotFound, "relation not found.")
	}

	merchantUsers := make([]merchantUserItem, 0)
	for _, relation := range relations {
		merchantUsers = append(merchantUsers, merchantUserItem{
			MerchantId:   relation.Merchant.ID,
			MerchantName: relation.Merchant.Name,
			Role:         relation.Role,
		})
	}

	return &userRelatedMerchantResponse{
		MerchantUser: merchantUsers,
	}, nil
}

func CheckUserExist(typ interface{}, r *http_util.HTTPContext) (resp interface{}, respErr error) {
	address := eth.ToEthAddress(r.QueryWithDefault("address", ""))

	if !eth.IsValidEthAddress(address) {
		return nil, http_util.NewHttpError(http.StatusBadRequest, "eth address is invalid.")
	}

	user, userErr := db.FindUserByAddress(address)
	if userErr != nil || user.ID == 0 {
		return false, nil
	}

	return true, nil
}

func LoginUseEthSignature(typ interface{}, r *http_util.HTTPContext) (resp interface{}, err error) {
	req := typ.(*loginUseEthSignatureRequest)
	address := eth.ToEthAddress(req.EthAddress)

	if !eth.IsValidEthAddress(address) {
		return nil, http_util.NewHttpError(http.StatusBadRequest, "eth address is invalid.")
	}

	if val, ok := userSignContentSet.Load(address); !ok {
		return nil, http_util.NewHttpError(http.StatusBadRequest, "sign content is invalid.")
	} else {
		pk, parseErr := http_util.GetUserPublicKeyFromSignWithCustomContent(val, req.SignData)
		if parseErr != nil {
			return nil, http_util.NewHttpError(http.StatusBadRequest, "sign data is invalid.")
		}

		if pk.Address().String() != address {
			log.Println(pk.Address().String(), address)
			return nil, http_util.NewHttpError(http.StatusBadRequest, "sign address is invalid.")
		}

		user, userErr := db.FindUserByAddress(address)
		if userErr != nil || user.ID == 0 {
			return nil, http_util.NewHttpError(http.StatusNotFound, "user not found.")
		}

		if user.Status == "unavailable" {
			return nil, http_util.NewHttpError(http.StatusForbidden, "user was suspend. please connect to admin")
		}

		// get permission of merchant
		merchantUser, merchantUserErr := db.FindMerchantUserByUserMerchant(user.ID, req.MerchantId)
		if merchantUserErr != nil || merchantUser.ID == 0 {
			return nil, http_util.NewHttpError(http.StatusForbidden, "current user has no permission for this store")
		}

		// check merchant
		merchant, merchantErr := db.FindMerchantById(req.MerchantId)
		if merchantErr != nil {
			return nil, merchantErr
		}

		if merchant.ID == 0 || merchant.Status != "available" {
			return nil, http_util.NewHttpError(http.StatusForbidden, "this merchant was suspend. please contact the admin")
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, http_util.MerchantClaims{
			UserId:     user.ID,
			MerchantId: req.MerchantId,
			Role:       "admin",
		})
		expires := time.Now().Add(12 * time.Hour)

		tokenStr, respErr := token.SignedString([]byte(config.Get().HTTP.JwtEncryptSecret))
		if respErr != nil {
			return nil, respErr
		}

		r.SetCookies(&http.Cookie{
			Name:     "user-auth-token",
			Value:    tokenStr,
			Path:     "/",
			Expires:  expires,
			HttpOnly: true,
			SameSite: 0,
		})

		return true, err
	}
}

func RegisterUseEthSignature(typ interface{}, r *http_util.HTTPContext) (resp interface{}, err error) {
	req := typ.(*registerUseEthSignatureRequest)
	address := eth.ToEthAddress(req.EthAddress)

	if !eth.IsValidEthAddress(address) {
		return nil, http_util.NewHttpError(http.StatusBadRequest, "eth address is invalid.")
	}

	if val, ok := userSignContentSet.Load(address); !ok {
		return nil, http_util.NewHttpError(http.StatusBadRequest, "sign content is invalid.")
	} else {
		pk, parseErr := http_util.GetUserPublicKeyFromSignWithCustomContent(val, req.SignData)
		if parseErr != nil {
			return nil, http_util.NewHttpError(http.StatusBadRequest, "sign data is invalid.")
		}

		if pk.Address().String() != address {
			return nil, http_util.NewHttpError(http.StatusBadRequest, "sign address is invalid.")
		}

		// get did pub key
		didDoc, didErr := did.ParseUserDIDFromJsonStr([]byte(req.DidDocument))
		if didErr != nil {
			return nil, http_util.NewHttpError(http.StatusBadRequest, "did document content is invalid.")
		}

		_, pubRaw, decodeErr := multibase.Decode(didDoc.VerificationMethod[0].MultibaseKey)
		if decodeErr != nil {
			return nil, http_util.NewHttpError(http.StatusBadRequest, "did document multibase key is invalid.")
		}
		pubKey, pubErr := eth.NewPublicKeyFromByte(pubRaw)
		if pubErr != nil {
			return nil, http_util.NewHttpError(http.StatusBadRequest, "load public key failed.")
		}

		if pubKey.Address().String() != address {
			return nil, http_util.NewHttpError(http.StatusBadRequest, "did document address is invalid.")
		}

		// check user
		user, userErr := db.FindUserByAddress(address)
		if userErr != nil || user.ID == 0 {
			// auto register new user
			newUser := db.User{
				Address:     address,
				Did:         "did:veric:" + address,
				Status:      "available",
				DidPubKey:   didDoc.VerificationMethod[0].MultibaseKey,
				DidUpStatus: db.DidUpStatusCreated,
			}

			newUserErr := db.SaveUser(&newUser)
			if newUserErr != nil {
				return nil, http_util.NewHttpError(http.StatusInternalServerError, "create user failed. please connect to admin")
			}
			user = &newUser
		} else {
			if user.DidPubKey == "" {
				// update did pub key
				user.Did = "did:veric:" + address
				user.DidPubKey = didDoc.VerificationMethod[0].MultibaseKey
				user.DidUpStatus = db.DidUpStatusCreated

				saveErr := db.SaveUser(user)
				if saveErr != nil {
					return nil, http_util.NewHttpError(http.StatusInternalServerError, "update user did failed. please connect to admin")
				}
			}
		}

		if user.Status == "unavailable" {
			return nil, http_util.NewHttpError(http.StatusForbidden, "user was suspend. please connect to admin")
		}

		// check merchant
		merchant, checkErr := db.FindMerchantByOwnerId(user.ID)
		if checkErr != nil {
			return nil, http_util.NewHttpError(http.StatusInternalServerError, "check merchant failed. please connect to admin")
		}

		if merchant != nil && merchant.ID > 0 {
			return nil, http_util.NewHttpError(http.StatusBadRequest, "already register merchant. can not register again")
		}

		// create merchant
		nowTime := time.Now()
		newMerchant := db.Merchant{
			Name: req.StoreName,
			Applications: []db.Application{{
				Name:            req.StoreName,
				Link:            req.StoreLink,
				CallbackUrl:     req.CallbackUrl,
				Slippage:        1,
				ApiKey:          util.RandString(32),
				ApiKeyCreatedAt: nowTime,
				IpnKey:          util.RandString(32),
				IpnKeyCreatedAt: nowTime,
				Status:          "available",
			}},
			Did:     user.Did,
			OwnerId: user.ID,
			Status:  "available",
		}

		createMerchantErr := db.CreateMerchant(&newMerchant)
		if createMerchantErr != nil {
			return nil, http_util.NewHttpError(http.StatusInternalServerError, "create merchant failed. please connect to admin")
		}

		// create permission relation
		newMerchantUser := db.MerchantUser{
			MerchantId: newMerchant.ID,
			UserId:     user.ID,
			Role:       "admin",
			Status:     "available",
		}
		merchantUserErr := db.SaveMerchantUser(&newMerchantUser)
		if merchantUserErr != nil {
			return nil, http_util.NewHttpError(http.StatusInternalServerError, "grant merchant access. please connect to admin")
		}

		return true, err
	}
}

func GetUserNickname(typ interface{}, r *http_util.HTTPContext) (resp interface{}, respErr error) {
	address := eth.ToEthAddress(r.QueryWithDefault("address", ""))

	if !eth.IsValidEthAddress(address) {
		return nil, http_util.NewHttpError(http.StatusBadRequest, "eth address is invalid.")
	}

	user, userErr := db.FindUserByAddress(address)
	if userErr != nil || user.ID == 0 {
		return nil, http_util.NewHttpError(http.StatusNotFound, "user not found.")
	}

	return &userNicknameResponse{
		Nickname: user.Nickname,
	}, nil
}

func SetNicknameUseEthSignature(typ interface{}, r *http_util.HTTPContext) (resp interface{}, err error) {
	req := typ.(*setNicknameUseEthSignatureRequest)
	address := eth.ToEthAddress(req.EthAddress)

	if !eth.IsValidEthAddress(address) {
		return nil, http_util.NewHttpError(http.StatusBadRequest, "eth address is invalid.")
	}

	if val, ok := userSignContentSet.Load(address); !ok {
		return nil, http_util.NewHttpError(http.StatusBadRequest, "sign content is invalid.")
	} else {
		pk, parseErr := http_util.GetUserPublicKeyFromSignWithCustomContent(val, req.SignData)
		if parseErr != nil {
			return nil, http_util.NewHttpError(http.StatusBadRequest, "sign data is invalid.")
		}

		if pk.Address().String() != address {
			log.Println(pk.Address().String(), address)
			return nil, http_util.NewHttpError(http.StatusBadRequest, "sign address is invalid.")
		}

		user, userErr := db.FindUserByAddress(address)
		if userErr != nil || user.ID == 0 {
			return nil, http_util.NewHttpError(http.StatusNotFound, "user not found.")
		}

		if user.Status == "unavailable" {
			return nil, http_util.NewHttpError(http.StatusForbidden, "user was suspend. please connect to admin")
		}

		user.Nickname = req.Nickname

		saveErr := db.SaveUser(user)
		if saveErr != nil {
			return nil, http_util.NewHttpError(http.StatusInternalServerError, "set nickname failed. please connect to admin")
		}

		return true, err
	}
}

// todo
func FastLoginForTest(typ interface{}, r *http_util.HTTPContext) (resp interface{}, err error) {
	req := typ.(*loginUseEthSignatureRequest)

	user, userErr := db.FindUserByAddress(req.EthAddress)
	if userErr != nil || user.ID == 0 {
		return nil, http_util.NewHttpError(http.StatusNotFound, "user not found.")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, http_util.MerchantClaims{
		UserId:     user.ID,
		MerchantId: req.MerchantId,
		Role:       "admin",
	})
	expires := time.Now().Add(12 * time.Hour)

	tokenStr, respErr := token.SignedString([]byte(config.Get().HTTP.JwtEncryptSecret))
	if respErr != nil {
		return nil, respErr
	}

	r.SetCookies(&http.Cookie{
		Name:     "user-auth-token",
		Value:    tokenStr,
		Path:     "/",
		Expires:  expires,
		HttpOnly: true,
		SameSite: 0,
	})

	return &loginUseEthSignatureResponse{
		Token:     tokenStr,
		ExpiredAt: expires.Format(util.TimeLayer),
	}, err
}
