package administration

import (
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"veric-backend/internal/log"
	"veric-backend/logic/db"
	"veric-backend/logic/http/http_util"
)

func GetManagerList(typ interface{}, r *http_util.HTTPContext) (resp interface{}, respErr error) {
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
		managers []db.Manager
		total    uint
	)

	total = db.GetManagerCountByStatus(status)
	managers, findErr := db.FindPaginationManagerWithPermissionByStatus(status, offset, size)
	if findErr != nil {
		log.GetLogger().Error("get manager list failed", zap.String("error", findErr.Error()))
		return nil, http_util.NewHttpError(http.StatusInternalServerError, "get manager list failed")
	}

	return &managerListResponse{
		Total:    total,
		Managers: managers,
	}, nil
}

func CreateManager(typ interface{}, r *http_util.HTTPContext) (resp interface{}, respErr error) {
	req := typ.(*createManagerRequest)

	// get admin
	jwtClaims, jwtErr := r.GetAdminJwt("admin-auth-token")
	if jwtErr != nil {
		return nil, jwtErr
	}

	if !r.CheckPermissionAny(jwtClaims.Privileges, []string{"sub-account"}) {
		return nil, http_util.NewHttpError(http.StatusForbidden, "no permission for this operation")
	}

	// check manager name
	existManager, nameErr := db.FindManagerByUserName(req.Name)
	if nameErr != nil {
		return nil, nameErr
	}

	if existManager.ID > 0 {
		return nil, http_util.NewHttpError(http.StatusBadRequest, "manager name already exist")
	}

	// check permission
	permissions, permissionErr := db.FindAvailablePermissionByIdGroup(req.Privileges, 1)
	if permissionErr != nil {
		return nil, http_util.NewHttpError(http.StatusInternalServerError, "check permission failed. please try again later")
	}

	if len(permissions) != len(req.Privileges) {
		return nil, http_util.NewHttpError(http.StatusBadRequest, "invalid permission set")
	}

	// calc new password
	hash, generateErr := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)
	if generateErr != nil {
		log.GetLogger().Error("calc new password failed.", zap.String("error", generateErr.Error()))
		return nil, http_util.NewHttpError(http.StatusInternalServerError, "change password failed.  please connect to admin")
	}

	newManager := &db.Manager{
		Username: req.Name,
		Password: string(hash),
		Nickname: req.Name,
		Email:    req.Email,
		Status:   "available",
	}

	saveErr := db.SaveManager(newManager)
	if saveErr != nil {
		log.GetLogger().Error("save manager failed.", zap.String("error", saveErr.Error()))
		return nil, http_util.NewHttpError(http.StatusInternalServerError, "save manager failed.  please connect to admin")
	}

	replaceErr := db.ReplaceManagerPermission(newManager, req.Privileges)
	if replaceErr != nil {
		return nil, http_util.NewHttpError(http.StatusInternalServerError, "replace manager permission failed. please connect to admin")
	}

	return true, nil
}

func ChangeManager(typ interface{}, r *http_util.HTTPContext) (resp interface{}, respErr error) {
	req := typ.(*changeManagerRequest)

	// get admin
	jwtClaims, jwtErr := r.GetAdminJwt("admin-auth-token")
	if jwtErr != nil {
		return nil, jwtErr
	}

	if !r.CheckPermissionAny(jwtClaims.Privileges, []string{"sub-account"}) {
		return nil, http_util.NewHttpError(http.StatusForbidden, "no permission for this operation")
	}

	// get manager
	manager, getErr := db.FindManagerById(req.ManagerId)
	if getErr != nil {
		return nil, getErr
	}

	if manager.ID == 0 {
		return nil, http_util.NewHttpError(http.StatusNotFound, "manager not found")
	}

	if manager.ID == 1 {
		return nil, http_util.NewHttpError(http.StatusForbidden, "can not change admin account")
	}

	// update name
	if req.Name != "" && req.Name != manager.Username {
		// check manager name
		existManager, nameErr := db.FindManagerByUserName(req.Name)
		if nameErr != nil {
			return nil, nameErr
		}

		if existManager.ID > 0 {
			return nil, http_util.NewHttpError(http.StatusBadRequest, "manager name already exist")
		}

		manager.Username = req.Name
		manager.Nickname = req.Name
	}

	// update password
	if req.Password != "" {
		// calc new password
		hash, generateErr := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)
		if generateErr != nil {
			log.GetLogger().Error("calc new password failed.", zap.String("error", generateErr.Error()))
			return nil, http_util.NewHttpError(http.StatusInternalServerError, "change password failed.  please connect to admin")
		}

		manager.Password = string(hash)
	}

	// update email
	if req.Email != "" {
		manager.Email = req.Email
	}

	// update permission
	if len(req.Privileges) > 0 {
		// check permission
		permissions, permissionErr := db.FindAvailablePermissionByIdGroup(req.Privileges, 1)
		if permissionErr != nil {
			return nil, http_util.NewHttpError(http.StatusInternalServerError, "check permission failed. please try again later")
		}

		if len(permissions) != len(req.Privileges) {
			return nil, http_util.NewHttpError(http.StatusBadRequest, "invalid permission set")
		}

		replaceErr := db.ReplaceManagerPermission(manager, req.Privileges)
		if replaceErr != nil {
			return nil, http_util.NewHttpError(http.StatusInternalServerError, "replace manager permission failed. please connect to admin")
		}
	}

	saveErr := db.SaveManager(manager)
	if saveErr != nil {
		log.GetLogger().Error("save manager failed.", zap.String("error", saveErr.Error()))
		return nil, http_util.NewHttpError(http.StatusInternalServerError, "save manager failed.  please connect to admin")
	}

	return true, nil
}

func DeleteManager(typ interface{}, r *http_util.HTTPContext) (resp interface{}, respErr error) {
	req := typ.(*deleteManagerRequest)

	// get admin
	jwtClaims, jwtErr := r.GetAdminJwt("admin-auth-token")
	if jwtErr != nil {
		return nil, jwtErr
	}

	if !r.CheckPermissionAny(jwtClaims.Privileges, []string{"sub-account"}) {
		return nil, http_util.NewHttpError(http.StatusForbidden, "no permission for this operation")
	}

	// get manager
	manager, getErr := db.FindManagerById(req.ManagerId)
	if getErr != nil {
		return nil, getErr
	}

	if manager.ID == 0 {
		return nil, http_util.NewHttpError(http.StatusNotFound, "manager not found")
	}

	if manager.ID == 1 {
		return nil, http_util.NewHttpError(http.StatusForbidden, "can not delete admin account")
	}

	// delete permission
	replaceErr := db.ReplaceManagerPermission(manager, []uint{})
	if replaceErr != nil {
		return nil, http_util.NewHttpError(http.StatusInternalServerError, "replace manager permission failed. please connect to admin")
	}

	deleteErr := db.DeleteManagerById(manager.ID)
	if deleteErr != nil {
		log.GetLogger().Error("delete manager failed.", zap.String("error", deleteErr.Error()))
		return nil, http_util.NewHttpError(http.StatusInternalServerError, "delete manager failed.  please connect to admin")
	}

	return true, nil
}
