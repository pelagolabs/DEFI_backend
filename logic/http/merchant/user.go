package merchant

import (
	"go.uber.org/zap"
	"net/http"
	"strings"
	"veric-backend/internal/log"
	"veric-backend/logic/blockchain/eth"
	"veric-backend/logic/db"
	"veric-backend/logic/http/http_util"
)

func GetMerchantUserList(typ interface{}, r *http_util.HTTPContext) (resp interface{}, respErr error) {
	status := r.QueryWithDefault("status", "all")
	page := r.QueryWithDefaultInt("page", 1)
	size := r.QueryWithDefaultInt("size", 10)

	// get user
	jwtClaims, jwtErr := r.GetMerchantJwt("user-auth-token")
	if jwtErr != nil {
		return nil, http_util.NewHttpError(http.StatusUnauthorized, "please login first")
	}

	if !r.CheckPermissionAny([]string{jwtClaims.Role}, []string{"admin", "shop_manager"}) {
		return nil, http_util.NewHttpError(http.StatusForbidden, "no permission for this operation")
	}

	// calc pagination
	offset := (page - 1) * size

	// get transaction list
	var (
		users []db.MerchantUser
		total uint
	)

	total = db.GetMerchantUserCountByMerchantStatus(jwtClaims.MerchantId, status)
	users, findErr := db.FindPaginationMerchantUserWithUserByStatus(jwtClaims.MerchantId, status, offset, size)
	if findErr != nil {
		log.GetLogger().Error("get merchant user list failed", zap.String("error", findErr.Error()))
		return nil, http_util.NewHttpError(http.StatusInternalServerError, "get user list failed")
	}

	return &merchantUserListResponse{
		Total:        total,
		MerchantUser: users,
	}, nil
}

func GrantUserMerchantRole(typ interface{}, r *http_util.HTTPContext) (resp interface{}, respErr error) {
	req := typ.(*grantUserMerchantRoleRequest)

	// get user
	jwtClaims, jwtErr := r.GetMerchantJwt("user-auth-token")
	if jwtErr != nil {
		return nil, http_util.NewHttpError(http.StatusUnauthorized, "please login first")
	}

	// can not grant admin
	if !r.CheckPermissionAny([]string{jwtClaims.Role}, []string{"admin"}) {
		return nil, http_util.NewHttpError(http.StatusForbidden, "no permission for this operation")
	}

	if req.Role == "admin" {
		return nil, http_util.NewHttpError(http.StatusBadRequest, "can not grant admin to others")
	}

	var (
		targetSet = strings.Split(req.TargetUser, ":")
		user      *db.User
		userErr   error
		address   string
	)

	// check user
	if len(targetSet) == 3 && strings.ToLower(targetSet[0]) == "did" {
		address = eth.ToEthAddress(targetSet[2])

		user, userErr = db.FindUserByDid(req.TargetUser)
	} else {
		address = eth.ToEthAddress(req.TargetUser)

		user, userErr = db.FindUserByAddress(req.TargetUser)
	}

	if userErr != nil {
		return nil, http_util.NewHttpError(http.StatusInternalServerError, "check user failed. please try again later")
	}

	if user.ID == 0 {
		if !eth.IsValidEthAddress(address) {
			return nil, http_util.NewHttpError(http.StatusBadRequest, "target wallet address is invalid.")
		}

		// create a new user
		newUser := &db.User{
			Address:  address,
			Did:      "did:veric:" + address,
			Nickname: "",
			Status:   "available",
		}

		newUserErr := db.SaveUser(newUser)
		if newUserErr != nil {
			return nil, http_util.NewHttpError(http.StatusInternalServerError, "create user failed. please try again later")
		}

		user.ID = newUser.ID
	}

	// check exist relation
	relation, relationErr := db.FindMerchantUserByUserMerchant(user.ID, jwtClaims.MerchantId)
	if relationErr != nil {
		return nil, http_util.NewHttpError(http.StatusInternalServerError, "check exist relation failed. please try again later")
	}

	if relation.ID > 0 {
		return nil, http_util.NewHttpError(http.StatusBadRequest, "user already be granted permission")
	}

	// create permission relation
	newMerchantUser := db.MerchantUser{
		MerchantId: jwtClaims.MerchantId,
		UserId:     user.ID,
		Role:       req.Role,
		Status:     "available",
	}
	merchantUserErr := db.SaveMerchantUser(&newMerchantUser)
	if merchantUserErr != nil {
		return nil, http_util.NewHttpError(http.StatusInternalServerError, "grant merchant access failed. please connect to admin")
	}

	return true, nil
}

func RevokeUserMerchantRole(typ interface{}, r *http_util.HTTPContext) (resp interface{}, respErr error) {
	req := typ.(*revokeUserMerchantRoleRequest)

	// get user
	jwtClaims, jwtErr := r.GetMerchantJwt("user-auth-token")
	if jwtErr != nil {
		return nil, http_util.NewHttpError(http.StatusUnauthorized, "please login first")
	}

	if !r.CheckPermissionAny([]string{jwtClaims.Role}, []string{"admin"}) {
		return nil, http_util.NewHttpError(http.StatusForbidden, "no permission for this operation")
	}

	// get exist relation
	relation, relationErr := db.FindMerchantUserById(req.RelationId)
	if relationErr != nil {
		return nil, http_util.NewHttpError(http.StatusInternalServerError, "check exist relation failed. please try again later")
	}

	if relation.ID == 0 || relation.MerchantId != jwtClaims.MerchantId {
		return nil, http_util.NewHttpError(http.StatusForbidden, "no permission for this operation")
	}

	if relation.Role == "admin" {
		return nil, http_util.NewHttpError(http.StatusBadRequest, "can not revoke yourself")
	}

	// delete relation
	deleteErr := db.DeleteMerchantUserById(relation.ID)
	if deleteErr != nil {
		return nil, http_util.NewHttpError(http.StatusInternalServerError, "revoke user grant failed. please connect to admin")
	}

	return true, nil
}

func ChangeUserMerchantRole(typ interface{}, r *http_util.HTTPContext) (resp interface{}, respErr error) {
	req := typ.(*changeUserMerchantRoleRequest)

	// get user
	jwtClaims, jwtErr := r.GetMerchantJwt("user-auth-token")
	if jwtErr != nil {
		return nil, http_util.NewHttpError(http.StatusUnauthorized, "please login first")
	}

	if !r.CheckPermissionAny([]string{jwtClaims.Role}, []string{"admin"}) {
		return nil, http_util.NewHttpError(http.StatusForbidden, "no permission for this operation")
	}

	// can not grant admin
	if req.Role == "admin" {
		return nil, http_util.NewHttpError(http.StatusBadRequest, "can not grant admin to others")
	}

	// get exist relation
	relations, relationErr := db.FindMerchantUserByIds(req.RelationIds)
	if relationErr != nil {
		return nil, http_util.NewHttpError(http.StatusInternalServerError, "check exist relation failed. please try again later")
	}

	for _, relation := range relations {
		if relation.ID == 0 || relation.MerchantId != jwtClaims.MerchantId {
			return nil, http_util.NewHttpError(http.StatusForbidden, "no permission for this operation")
		}
	}

	// update relation
	saveErr := db.UpdateMerchantUserRoleByIds(req.RelationIds, req.Role)
	if saveErr != nil {
		return nil, http_util.NewHttpError(http.StatusInternalServerError, "change merchant access failed. please connect to admin")
	}

	return true, nil
}
