package administration

import (
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"veric-backend/internal/log"
	"veric-backend/logic/db"
	"veric-backend/logic/http/http_util"
)

func GetApplicationList(typ interface{}, r *http_util.HTTPContext) (resp interface{}, respErr error) {
	status := r.QueryWithDefault("status", "all")
	page := r.QueryWithDefaultInt("page", 1)
	size := r.QueryWithDefaultInt("size", 10)
	merchantId := uint(r.QueryWithDefaultInt("merchant_id", 0))

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
		applications []db.Application
		total        uint
	)

	total = db.GetApplicationCountByStatusMerchant(status, merchantId)
	applications, findErr := db.FindPaginationApplicationByStatusMerchant(status, offset, size, merchantId)
	if findErr != nil {
		log.GetLogger().Error("get application list failed", zap.String("error", findErr.Error()))
		return nil, http_util.NewHttpError(http.StatusInternalServerError, "get application list failed")
	}

	return &applicationListResponse{
		Total:        total,
		Applications: applications,
	}, nil
}

func ChangeApplicationStatus(typ interface{}, r *http_util.HTTPContext) (resp interface{}, saveErr error) {
	req := typ.(*changeApplicationStatusRequest)

	// get admin
	jwtClaims, jwtErr := r.GetAdminJwt("admin-auth-token")
	if jwtErr != nil {
		return nil, jwtErr
	}

	if !r.CheckPermissionAny(jwtClaims.Privileges, []string{"merchant"}) {
		return nil, http_util.NewHttpError(http.StatusForbidden, "no permission for this operation")
	}

	// get merchant
	application, saveErr := db.FindApplicationById(req.ApplicationId)
	if saveErr != nil || application.ID == 0 {
		return nil, http_util.NewHttpError(http.StatusNotFound, "application not exist")
	}

	if application.Status == req.NewStatus {
		return nil, http_util.NewHttpError(http.StatusBadRequest, "new status can not same as old")
	}

	// todo
	// do sht when changing application status

	// update status
	application.Status = req.NewStatus

	saveErr = db.SaveApplication(&application)
	if saveErr != nil {
		log.GetLogger().Error("save application failed", zap.String("error", saveErr.Error()))
		return nil, http_util.NewHttpError(http.StatusInternalServerError, "save application failed")
	}

	log.GetLogger().Info("change application status succeed by admin:" + strconv.Itoa(int(jwtClaims.AdminId)))

	return true, nil
}
