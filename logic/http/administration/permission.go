package administration

import (
	"go.uber.org/zap"
	"net/http"
	"veric-backend/internal/log"
	"veric-backend/logic/db"
	"veric-backend/logic/http/http_util"
)

func GetAllPermission(r *http_util.HTTPContext) (resp interface{}, respErr error) {
	// get admin
	_, jwtErr := r.GetAdminJwt("admin-auth-token")
	if jwtErr != nil {
		return nil, http_util.NewHttpError(http.StatusForbidden, "need permission for this operate")
	}

	// get permission list
	var (
		permissions []db.Permission
		total       uint
	)

	total = db.GetAllPermissionCountByGroup(1)
	permissions, findErr := db.FindAllPermissionByGroup(1)
	if findErr != nil {
		log.GetLogger().Error("get permission list failed", zap.String("error", findErr.Error()))
		return nil, http_util.NewHttpError(http.StatusInternalServerError, "get permission list failed")
	}

	return &permissionListResponse{
		Total:       total,
		Permissions: permissions,
	}, nil
}
