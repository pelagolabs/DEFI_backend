package administration

import (
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"veric-backend/internal/log"
	"veric-backend/logic/db"
	"veric-backend/logic/http/http_util"
)

func GetCurrencyList(typ interface{}, r *http_util.HTTPContext) (resp interface{}, respErr error) {
	status := r.QueryWithDefault("status", "all")
	page := r.QueryWithDefaultInt("page", 1)
	size := r.QueryWithDefaultInt("size", 10)

	// get admin
	_, jwtErr := r.GetAdminJwt("admin-auth-token")
	if jwtErr != nil {
		return nil, http_util.NewHttpError(http.StatusForbidden, "need permission for this operate")
	}

	// calc pagination
	offset := (page - 1) * size

	// get transaction list
	var (
		currencies []db.Currency
		total      uint
	)

	total = db.GetCurrencyCountByStatus(status)
	currencies, findErr := db.FindPaginationCurrencyByStatus(status, offset, size)
	if findErr != nil {
		log.GetLogger().Error("get currency list failed", zap.String("error", findErr.Error()))
		return nil, http_util.NewHttpError(http.StatusInternalServerError, "get currency list failed")
	}

	return &currencyListResponse{
		Total:      total,
		Currencies: currencies,
	}, nil
}

func ChangeCurrencyStatus(typ interface{}, r *http_util.HTTPContext) (resp interface{}, saveErr error) {
	req := typ.(*changeCurrencyStatusRequest)

	// get admin
	jwtClaims, jwtErr := r.GetAdminJwt("admin-auth-token")
	if jwtErr != nil {
		return nil, jwtErr
	}

	if !r.CheckPermissionAny(jwtClaims.Privileges, []string{"currency"}) {
		return nil, http_util.NewHttpError(http.StatusForbidden, "no permission for this operation")
	}

	// get currency
	currency, saveErr := db.FindCurrencyById(req.CurrencyId)
	if saveErr != nil || currency.ID == 0 {
		return nil, http_util.NewHttpError(http.StatusNotFound, "currency not exist")
	}

	if currency.Status == req.NewStatus {
		return nil, http_util.NewHttpError(http.StatusBadRequest, "new status can not same as old")
	}

	// todo
	// do sht when changing currency status

	// update status
	currency.Status = req.NewStatus

	saveErr = db.SaveCurrency(currency)
	if saveErr != nil {
		log.GetLogger().Error("save currency failed", zap.String("error", saveErr.Error()))
		return nil, http_util.NewHttpError(http.StatusInternalServerError, "save currency failed")
	}

	log.GetLogger().Info("change currency status succeed by admin:" + strconv.Itoa(int(jwtClaims.AdminId)))

	return true, nil
}
