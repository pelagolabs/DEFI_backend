package administration

import (
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"veric-backend/internal/log"
	"veric-backend/logic/config"
	"veric-backend/logic/db"
	"veric-backend/logic/http/http_util"
)

func GetMerchantList(typ interface{}, r *http_util.HTTPContext) (resp interface{}, respErr error) {
	status := r.QueryWithDefault("status", "all")
	page := r.QueryWithDefaultInt("page", 1)
	size := r.QueryWithDefaultInt("size", 10)
	did := r.QueryWithDefault("did", "")

	// get admin
	jwtClaims, jwtErr := r.GetAdminJwt("admin-auth-token")
	if jwtErr != nil {
		return nil, http_util.NewHttpError(http.StatusForbidden, "need permission for this operate")
	}

	if !r.CheckPermissionAny(jwtClaims.Privileges, []string{"merchant"}) {
		return nil, http_util.NewHttpError(http.StatusForbidden, "no permission for this operation")
	}

	// calc pagination
	offset := (page - 1) * size

	// get transaction list
	var (
		merchants = make([]*db.Merchant, 0)
		merchant  *db.Merchant
		total     uint
		findErr   error
	)

	if did != "" {
		merchant, findErr = db.FindMerchantWithApplicationsByDid(did)
		if merchant.ID != 0 {
			total = 1
			merchants = append(merchants, merchant)
		}
	} else {
		total = db.GetMerchantCountByStatus(status)
		merchants, findErr = db.FindPaginationMerchantByStatus(status, offset, size)
	}

	if findErr != nil {
		log.GetLogger().Error("get merchant list failed", zap.String("error", findErr.Error()))
		return nil, http_util.NewHttpError(http.StatusInternalServerError, "get merchant list failed")
	}

	return &merchantListResponse{
		Total:      total,
		Merchants:  merchants,
		SuspendAll: config.Get().Logic.SuspendAllMerchant,
	}, nil
}

func ChangeMerchantStatus(typ interface{}, r *http_util.HTTPContext) (resp interface{}, saveErr error) {
	req := typ.(*changeMerchantStatusRequest)

	// get admin
	jwtClaims, jwtErr := r.GetAdminJwt("admin-auth-token")
	if jwtErr != nil {
		return nil, jwtErr
	}

	if !r.CheckPermissionAny(jwtClaims.Privileges, []string{"merchant"}) {
		return nil, http_util.NewHttpError(http.StatusForbidden, "no permission for this operation")
	}

	// get merchant
	merchant, saveErr := db.FindMerchantById(req.MerchantId)
	if saveErr != nil || merchant.ID == 0 {
		return nil, http_util.NewHttpError(http.StatusNotFound, "merchant not exist")
	}

	if merchant.Status == req.NewStatus {
		return nil, http_util.NewHttpError(http.StatusBadRequest, "new status can not same as old")
	}

	// todo
	// do sth when changing merchant status

	// update status
	merchant.Status = req.NewStatus

	saveErr = db.SaveMerchant(merchant)
	if saveErr != nil {
		log.GetLogger().Error("save merchant failed", zap.String("error", saveErr.Error()))
		return nil, http_util.NewHttpError(http.StatusInternalServerError, "save merchant failed")
	}

	log.GetLogger().Info("change merchant status succeed by admin:" + strconv.Itoa(int(jwtClaims.AdminId)))

	return true, nil
}

func SuspendAllMerchant(typ interface{}, r *http_util.HTTPContext) (resp interface{}, saveErr error) {
	req := typ.(*suspendAllMerchantRequest)

	// get admin
	jwtClaims, jwtErr := r.GetAdminJwt("admin-auth-token")
	if jwtErr != nil {
		return nil, jwtErr
	}

	if !r.CheckPermissionAny(jwtClaims.Privileges, []string{"merchant"}) {
		return nil, http_util.NewHttpError(http.StatusForbidden, "no permission for this operation")
	}

	wantToSuspendAll := false
	if req.NewStatus == "unavailable" {
		wantToSuspendAll = true
	}

	if config.Get().Logic.SuspendAllMerchant == wantToSuspendAll {
		return nil, http_util.NewHttpError(http.StatusBadRequest, "new status can not same as old")
	}

	// todo
	// do sth when changing merchant status

	// update status
	config.SetSuspendAllMerchant(wantToSuspendAll)

	log.GetLogger().Info("suspend all merchant succeed by admin:" + strconv.Itoa(int(jwtClaims.AdminId)))

	return true, nil
}
