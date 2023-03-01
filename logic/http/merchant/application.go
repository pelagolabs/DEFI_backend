package merchant

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
	"veric-backend/internal/util"
	"veric-backend/logic/db"
	"veric-backend/logic/http/http_util"
)

func GetApplicationDetail(r *http_util.HTTPContext) (resp interface{}, respErr error) {
	params := mux.Vars(r.Request)

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

	return application, nil
}

func ModifyApplicationBase(typ interface{}, r *http_util.HTTPContext) (resp interface{}, respErr error) {
	req := typ.(*modifyApplicationBaseRequest)

	// get user
	jwtClaims, jwtErr := r.GetMerchantJwt("user-auth-token")
	if jwtErr != nil {
		return nil, http_util.NewHttpError(http.StatusUnauthorized, "please login first")
	}

	if !r.CheckPermissionAny([]string{jwtClaims.Role}, []string{"admin"}) {
		return nil, http_util.NewHttpError(http.StatusForbidden, "no permission for this operation")
	}

	// get application
	application, err := getMerchantDefaultApplication(jwtClaims.MerchantId, req.AppId)
	if err != nil {
		return nil, err
	}

	// update application base info
	application.Name = req.Name
	application.Link = req.Link
	application.CallbackUrl = req.CallbackUrl
	application.Describe = req.Describe
	application.LegalTender = req.LegalTender

	saveErr := db.SaveApplication(application)
	if saveErr != nil {
		return nil, http_util.NewHttpError(http.StatusInternalServerError, "update application base failed. please connect to admin")
	}

	// update merchant name
	merchant, err := db.FindMerchantById(jwtClaims.MerchantId)
	if err != nil {
		return nil, err
	}

	merchant.Name = req.Name
	saveMerchantErr := db.SaveMerchant(merchant)
	if saveMerchantErr != nil {
		return nil, http_util.NewHttpError(http.StatusInternalServerError, "update merchant base failed. please connect to admin")
	}

	return true, nil
}

func ModifyApplicationCurrency(typ interface{}, r *http_util.HTTPContext) (resp interface{}, respErr error) {
	req := typ.(*modifyApplicationCurrencyRequest)

	// get user
	jwtClaims, jwtErr := r.GetMerchantJwt("user-auth-token")
	if jwtErr != nil {
		return nil, http_util.NewHttpError(http.StatusUnauthorized, "please login first")
	}

	if !r.CheckPermissionAny([]string{jwtClaims.Role}, []string{"admin"}) {
		return nil, http_util.NewHttpError(http.StatusForbidden, "no permission for this operation")
	}

	// check currency
	currencies, currencyErr := db.FindAvailableCurrenciesById(req.Currencies)
	if currencyErr != nil {
		return nil, http_util.NewHttpError(http.StatusInternalServerError, "check currency currency failed. please try again later")
	}

	if len(currencies) != len(req.Currencies) {
		return nil, http_util.NewHttpError(http.StatusBadRequest, "invalid currency set")
	}

	// get application
	application, err := getMerchantDefaultApplication(jwtClaims.MerchantId, req.AppId)
	if err != nil {
		return nil, err
	}

	replaceErr := db.ReplaceApplicationCurrency(application, req.Currencies)
	if replaceErr != nil {
		return nil, http_util.NewHttpError(http.StatusInternalServerError, "replace application currency failed. please connect to admin")
	}

	return true, nil
}

func ModifyApplicationApiKey(typ interface{}, r *http_util.HTTPContext) (resp interface{}, respErr error) {
	req := typ.(*modifyApplicationApiKeyRequest)

	// get user
	jwtClaims, jwtErr := r.GetMerchantJwt("user-auth-token")
	if jwtErr != nil {
		return nil, http_util.NewHttpError(http.StatusUnauthorized, "please login first")
	}

	if !r.CheckPermissionAny([]string{jwtClaims.Role}, []string{"admin"}) {
		return nil, http_util.NewHttpError(http.StatusForbidden, "no permission for this operation")
	}

	// get application
	application, err := getMerchantDefaultApplication(jwtClaims.MerchantId, req.AppId)
	if err != nil {
		return nil, err
	}

	// update application base info
	application.ApiKey = util.RandString(32)
	application.ApiKeyCreatedAt = time.Now()

	saveErr := db.SaveApplication(application)
	if saveErr != nil {
		return nil, http_util.NewHttpError(http.StatusInternalServerError, "update application api key failed. please connect to admin")
	}

	return &modifyApplicationApiKeyResponse{
		Key:       application.ApiKey,
		CreatedAt: application.ApiKeyCreatedAt.Format(util.TimeLayer),
	}, nil
}

func ModifyApplicationIpnKey(typ interface{}, r *http_util.HTTPContext) (resp interface{}, respErr error) {
	req := typ.(*modifyApplicationIpnKeyRequest)

	// get user
	jwtClaims, jwtErr := r.GetMerchantJwt("user-auth-token")
	if jwtErr != nil {
		return nil, http_util.NewHttpError(http.StatusUnauthorized, "please login first")
	}

	if !r.CheckPermissionAny([]string{jwtClaims.Role}, []string{"admin"}) {
		return nil, http_util.NewHttpError(http.StatusForbidden, "no permission for this operation")
	}

	// get application
	application, err := getMerchantDefaultApplication(jwtClaims.MerchantId, req.AppId)
	if err != nil {
		return nil, err
	}

	// update application base info
	application.IpnKey = util.RandString(32)
	application.IpnKeyCreatedAt = time.Now()

	saveErr := db.SaveApplication(application)
	if saveErr != nil {
		return nil, http_util.NewHttpError(http.StatusInternalServerError, "update application ipn key failed. please connect to admin")
	}

	return &modifyApplicationIpnKeyResponse{
		Key:       application.IpnKey,
		CreatedAt: application.IpnKeyCreatedAt.Format(util.TimeLayer),
	}, nil
}

func ModifyApplicationSlippage(typ interface{}, r *http_util.HTTPContext) (resp interface{}, respErr error) {
	req := typ.(*modifyApplicationSlippageRequest)

	// get user
	jwtClaims, jwtErr := r.GetMerchantJwt("user-auth-token")
	if jwtErr != nil {
		return nil, http_util.NewHttpError(http.StatusUnauthorized, "please login first")
	}

	if !r.CheckPermissionAny([]string{jwtClaims.Role}, []string{"admin"}) {
		return nil, http_util.NewHttpError(http.StatusForbidden, "no permission for this operation")
	}

	// get application
	application, err := getMerchantDefaultApplication(jwtClaims.MerchantId, req.AppId)
	if err != nil {
		return nil, err
	}

	if req.Slippage < 0 || req.Slippage > 100 {
		return nil, http_util.NewHttpError(http.StatusBadRequest, "invalid slippage. must between 0 and 100")
	}

	// update application slippage
	application.Slippage = req.Slippage

	saveErr := db.SaveApplication(application)
	if saveErr != nil {
		return nil, http_util.NewHttpError(http.StatusInternalServerError, "update application slippage failed. please connect to admin")
	}

	return true, nil
}
