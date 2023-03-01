package merchant

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gorilla/mux"
	"github.com/multiformats/go-multibase"
	"go.uber.org/zap"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"
	"veric-backend/internal/log"
	"veric-backend/internal/util"
	"veric-backend/logic/blockchain/did"
	"veric-backend/logic/blockchain/eth"
	"veric-backend/logic/blockchain/exchange"
	"veric-backend/logic/config"
	"veric-backend/logic/db"
	"veric-backend/logic/db/types"
	"veric-backend/logic/http/http_util"
	"veric-backend/logic/tasks"
)

func GetWithdrawSummary(typ interface{}, r *http_util.HTTPContext) (resp interface{}, respErr error) {
	params := mux.Vars(r.Request)

	tz := r.QueryWithDefaultInt("tz", 0)

	// check app id
	if _, ok := params["app_id"]; !ok {
		return nil, http_util.NewHttpError(http.StatusBadRequest, "invalid application id")
	}

	appId, atoiErr := strconv.Atoi(params["app_id"])
	if atoiErr != nil {
		return nil, http_util.NewHttpError(http.StatusBadRequest, "invalid application id")
	}

	// get user
	jwtClaims, jwtErr := r.GetMerchantJwt("user-auth-token")
	if jwtErr != nil {
		return nil, http_util.NewHttpError(http.StatusUnauthorized, "please login first")
	}

	if !r.CheckPermissionAny([]string{jwtClaims.Role}, []string{"admin", "shop_manager"}) {
		return nil, http_util.NewHttpError(http.StatusForbidden, "no permission for this operation")
	}

	// get application
	application, err := getMerchantDefaultApplication(jwtClaims.MerchantId, uint(appId))
	if err != nil {
		return nil, err
	}

	nowTime := time.Now().Add(time.Duration(tz) * time.Hour)
	dateBegin := nowTime.AddDate(0, 0, -90).Format(util.DateLayer)
	dateToday := nowTime.Format(util.DateLayer)
	withdrawSummary := &withdrawSummaryResponse{
		Latest90DaysTotalWithdraw: 0,
		TodayTotalWithdraw:        0,
	}

	// get latest 90 days summary
	latest90Stats, sum90Err := db.FindApplicationFinishWithdrawWithCurrencyByTime(application.ID, util.GetTimeStrFixedByTZ(dateBegin+" 00:00:00", tz), util.GetTimeStrFixedByTZ(dateToday+" 23:59:59", tz))
	if sum90Err != nil {
		return nil, sum90Err
	}

	for _, withdrawItem := range latest90Stats {
		if amountRaw, ok := big.NewInt(0).SetString(withdrawItem.Amount, 10); ok {
			withdrawAmountRaw, exchangeErr := exchange.DefaultManage.ExchangeCoinToUSD(amountRaw, withdrawItem.Currency)
			if exchangeErr != nil {
				continue
			}

			withdrawAmount, _ := withdrawAmountRaw.Float64()
			withdrawSummary.Latest90DaysTotalWithdraw += withdrawAmount
		}
	}

	// get today summary
	todayStats, sumErr := db.FindApplicationFinishWithdrawWithCurrencyByTime(application.ID, util.GetTimeStrFixedByTZ(dateToday+" 00:00:00", tz), util.GetTimeStrFixedByTZ(dateToday+" 23:59:59", tz))
	if sumErr != nil {
		return nil, sumErr
	}

	for _, withdrawItem := range todayStats {
		if amountRaw, ok := big.NewInt(0).SetString(withdrawItem.Amount, 10); ok {
			withdrawAmountRaw, exchangeErr := exchange.DefaultManage.ExchangeCoinToUSD(amountRaw, withdrawItem.Currency)
			if exchangeErr != nil {
				continue
			}

			withdrawAmount, _ := withdrawAmountRaw.Float64()
			withdrawSummary.TodayTotalWithdraw += withdrawAmount
		}
	}

	return withdrawSummary, nil
}

func GetWithdrawList(typ interface{}, r *http_util.HTTPContext) (resp interface{}, respErr error) {
	status := r.QueryWithDefault("status", "all")
	page := r.QueryWithDefaultInt("page", 1)
	size := r.QueryWithDefaultInt("size", 10)
	withdrawNum := r.QueryWithDefault("withdraw_num", "")
	currencyId := uint(r.QueryWithDefaultInt("currency_id", 0))
	date := r.QueryWithDefault("date", "")

	// check app id
	params := mux.Vars(r.Request)
	if _, ok := params["app_id"]; !ok {
		return nil, http_util.NewHttpError(http.StatusBadRequest, "invalid application id")
	}

	appId, atoiErr := strconv.Atoi(params["app_id"])
	if atoiErr != nil {
		return nil, http_util.NewHttpError(http.StatusBadRequest, "invalid application id")
	}

	// get user
	jwtClaims, jwtErr := r.GetMerchantJwt("user-auth-token")
	if jwtErr != nil {
		return nil, http_util.NewHttpError(http.StatusUnauthorized, "please login first")
	}

	if !r.CheckPermissionAny([]string{jwtClaims.Role}, []string{"admin", "shop_manager"}) {
		return nil, http_util.NewHttpError(http.StatusForbidden, "no permission for this operation")
	}

	// get application
	application, err := getMerchantDefaultApplication(jwtClaims.MerchantId, uint(appId))
	if err != nil {
		return nil, err
	}

	// calc pagination
	offset := (page - 1) * size

	// get transaction list
	var (
		withdraws = make([]withdrawListItem, 0)
		total     uint
	)

	if withdrawNum != "" {
		withdraw, findErr := db.FindApplicationWithdrawWithCurrencyByPaymentNum(application.ID, withdrawNum)
		if findErr != nil {
			log.GetLogger().Error("get withdraw list failed", zap.String("error", findErr.Error()))
			return nil, http_util.NewHttpError(http.StatusInternalServerError, "get withdraw list failed")
		}

		if withdraw.ID > 0 {
			total = 1

			amount := float64(0)
			if amountRaw, ok := big.NewInt(0).SetString(withdraw.Amount, 10); ok {
				amount, _ = exchange.DefaultManage.ExchangeCoinToNaturalAmount(amountRaw, withdraw.Currency).Float64()
			}

			withdraws = append(withdraws, withdrawListItem{
				WithdrawNum:    withdraw.WithdrawNum,
				CurrencySymbol: withdraw.Currency.Symbol,
				Amount:         amount,
				Status:         withdraw.Status,
				CreatedAt:      withdraw.CreatedAt,
				TxHash:         withdraw.TxHash,
			})
		}
	} else {
		total = db.GetApplicationWithdrawCountByStatusCurrencyDate(application.ID, status, currencyId, date)
		withdrawSet, findErr := db.FindPaginationApplicationWithdrawByStatusCurrencyDate(application.ID, status, currencyId, date, offset, size)
		if findErr != nil {
			log.GetLogger().Error("get withdraw list failed", zap.String("error", findErr.Error()))
			return nil, http_util.NewHttpError(http.StatusInternalServerError, "get withdraw list failed")
		}

		for _, withdraw := range withdrawSet {
			amount := float64(0)
			if amountRaw, ok := big.NewInt(0).SetString(withdraw.Amount, 10); ok {
				amount, _ = exchange.DefaultManage.ExchangeCoinToNaturalAmount(amountRaw, withdraw.Currency).Float64()
			}

			withdraws = append(withdraws, withdrawListItem{
				WithdrawNum:    withdraw.WithdrawNum,
				CurrencySymbol: withdraw.Currency.Symbol,
				Amount:         amount,
				Status:         withdraw.Status,
				CreatedAt:      withdraw.CreatedAt,
				TxHash:         withdraw.TxHash,
			})
		}
	}

	return &withdrawListResponse{
		Total:     total,
		Withdraws: withdraws,
	}, nil
}

func GetWithdrawDetail(r *http_util.HTTPContext) (resp interface{}, respErr error) {
	params := mux.Vars(r.Request)

	// check app id
	if _, ok := params["app_id"]; !ok {
		return nil, http_util.NewHttpError(http.StatusBadRequest, "invalid application id")
	}

	appId, atoiErr := strconv.Atoi(params["app_id"])
	if atoiErr != nil {
		return nil, http_util.NewHttpError(http.StatusBadRequest, "invalid application id")
	}

	// check withdraw id
	if _, ok := params["withdraw_id"]; !ok {
		return nil, http_util.NewHttpError(http.StatusBadRequest, "invalid withdraw id")
	}

	withdrawId, atoiErr := strconv.Atoi(params["withdraw_id"])
	if atoiErr != nil {
		return nil, http_util.NewHttpError(http.StatusBadRequest, "invalid withdraw id")
	}

	// get user
	jwtClaims, jwtErr := r.GetMerchantJwt("user-auth-token")
	if jwtErr != nil {
		return nil, http_util.NewHttpError(http.StatusUnauthorized, "please login first")
	}

	if !r.CheckPermissionAny([]string{jwtClaims.Role}, []string{"admin", "shop_manager"}) {
		return nil, http_util.NewHttpError(http.StatusForbidden, "no permission for this operation")
	}

	// get application
	application, err := getMerchantDefaultApplication(jwtClaims.MerchantId, uint(appId))
	if err != nil {
		return nil, err
	}

	withdraw, withdrawErr := db.FindApplicationWithdrawById(application.ID, uint(withdrawId))
	if withdrawErr != nil {
		return nil, withdrawErr
	}

	if withdraw == nil || withdraw.ID == 0 {
		return nil, http_util.NewHttpError(http.StatusNotFound, "withdraw not found")
	}

	return withdraw, nil
}

type DoWithdrawRequest struct {
	VP        string `json:"vp"`
	ToAddress string `json:"to_address"`
}

func findPublicKeyByVP(vp *did.VerifiablePresentation) (*eth.PublicKey, error) {
	verificationMethod := vp.Proof.VerificationMethod
	didKeyPos := strings.Index(verificationMethod, "#")
	if didKeyPos == -1 {
		return nil, errors.New("unknown verification method")
	}
	didKey := verificationMethod[:didKeyPos]

	user, err := db.FindUserByDid(didKey)
	if err != nil {
		return nil, err
	}

	if user.ID == 0 {
		return nil, errors.New("unknown verification method")
	}

	_, pbKeyByte, err := multibase.Decode(user.DidPubKey)
	if err != nil {
		return nil, err
	}

	return eth.NewPublicKeyFromByte(pbKeyByte)
}

func DoWithdraw(req *DoWithdrawRequest, r *http_util.HTTPContext) (resp interface{}, err error) {
	// get user
	jwtClaims, jwtErr := r.GetMerchantJwt("user-auth-token")
	if jwtErr != nil {
		return nil, http_util.NewHttpError(http.StatusUnauthorized, "please login first")
	}

	if !r.CheckPermissionAny([]string{jwtClaims.Role}, []string{"admin"}) {
		return nil, http_util.NewHttpError(http.StatusForbidden, "no permission for this operation")
	}

	// check global merchant status
	if config.Get().Logic.SuspendAllMerchant {
		return nil, errors.New("system operate paused. please connect to admin")
	}

	// check merchant status
	merchant, err := db.FindMerchantById(jwtClaims.MerchantId)
	if err != nil {
		return nil, err
	}

	if merchant.Status == "unavailable" {
		return nil, http_util.NewHttpError(http.StatusForbidden, "merchant was suspend. please connect to admin")
	}

	vpDocument, err := did.ParseVerifiablePresentationFromJsonStr([]byte(req.VP))
	if err != nil {
		return nil, err
	}

	if !common.IsHexAddress(req.ToAddress) {
		return nil, errors.New("invalid address")
	}

	pbKey, err := findPublicKeyByVP(vpDocument)
	if err != nil {
		return nil, err
	}

	verify, err := vpDocument.Verify(pbKey, did.IssuePriKey.PublicKey())
	if err != nil {
		return nil, err
	}

	if !verify {
		return nil, errors.New("vp invalid")
	}

	allId := make([]string, 0, 10)
	vcMap := make(map[string]*did.VerifiableCredential)
	for _, vcDocument := range vpDocument.VerifiableCredential {
		allId = append(allId, vcDocument.ID)
		vcCopy := vcDocument
		vcMap[vcDocument.ID] = &vcCopy
	}

	// todo need check merchant?
	vcs, err := db.FindVCsByVCIdAndStatus(allId, db.VCStatusActive, db.WithPreload("Payment"))
	if err != nil {
		return nil, err
	}

	currentChain := ""
	currentCurrency := ""
	totalAmount := types.NewBigIntZero()
	totalVCID := make([]uint, 0)
	for _, vc := range vcs {
		vcSubjectDeposit, err := did.ParseVCSubjectDepositFromVerifiableCredential(vcMap[vc.VCID])
		if err != nil {
			return nil, err
		}

		if currentChain == "" {
			currentChain = vcSubjectDeposit.Chain
			currentCurrency = vcSubjectDeposit.Currency
		} else if currentChain != vcSubjectDeposit.Chain || currentCurrency != vcSubjectDeposit.Currency {
			return nil, errors.New("current not support multi chain or currency")
		}

		if vcAmount, ok := big.NewInt(0).SetString(vcSubjectDeposit.MerchantAmount, 10); ok {
			totalAmount.Add(types.NewBigInt(vcAmount))
			totalVCID = append(totalVCID, vc.ID)
		}
	}

	if len(totalVCID) == 0 {
		return nil, errors.New("invalid vcs in vp")
	}

	conf := &tasks.WithdrawTaskConfig{
		VCId:        totalVCID,
		ToAddress:   common.HexToAddress(req.ToAddress),
		Chain:       currentChain,
		Currency:    currentCurrency,
		TotalAmount: totalAmount.RawBigInt(),
	}
	err = tasks.DefaultManage.NewTaskWithTx(tasks.TaskTypeWithdraw, conf, func(tx db.Options) error {
		for _, vc := range vcs {
			vc.VCStatus = db.VCStatusProcessing
			err = db.UpdateVC(vc, db.VCStatusActive, tx)
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}
