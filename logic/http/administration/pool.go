package administration

import (
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"veric-backend/internal/log"
	"veric-backend/logic/db"
	"veric-backend/logic/http/http_util"
	"veric-backend/logic/tasks"
)

func GetPoolList(typ interface{}, r *http_util.HTTPContext) (resp interface{}, respErr error) {
	status := r.QueryWithDefault("status", "all")
	page := r.QueryWithDefaultInt("page", 1)
	size := r.QueryWithDefaultInt("size", 10)

	// get admin
	jwtClaims, jwtErr := r.GetAdminJwt("admin-auth-token")
	if jwtErr != nil {
		return nil, http_util.NewHttpError(http.StatusForbidden, "need permission for this operate")
	}

	if !r.CheckPermissionAny(jwtClaims.Privileges, []string{"liquidity-pool"}) {
		return nil, http_util.NewHttpError(http.StatusForbidden, "no permission for this operation")
	}

	// calc pagination
	offset := (page - 1) * size

	// get transaction list
	var (
		pools []db.Pool
		total uint
	)

	total = db.GetPoolCountByStatus(status)
	pools, findErr := db.FindPaginationPoolByStatus(status, offset, size)
	if findErr != nil {
		log.GetLogger().Error("get pool list failed", zap.String("error", findErr.Error()))
		return nil, http_util.NewHttpError(http.StatusInternalServerError, "get pool list failed")
	}

	return &poolListResponse{
		Total: total,
		Pools: pools,
	}, nil
}

func ChangePoolStatus(typ interface{}, r *http_util.HTTPContext) (resp interface{}, saveErr error) {
	req := typ.(*changePoolStatusRequest)

	// get admin
	jwtClaims, jwtErr := r.GetAdminJwt("admin-auth-token")
	if jwtErr != nil {
		return nil, jwtErr
	}

	if !r.CheckPermissionAny(jwtClaims.Privileges, []string{"liquidity-pool"}) {
		return nil, http_util.NewHttpError(http.StatusForbidden, "no permission for this operation")
	}

	// get pool
	pool, findErr := db.FindPoolById(req.PoolId)
	if findErr != nil || pool.ID == 0 {
		return nil, http_util.NewHttpError(http.StatusNotFound, "pool not exist")
	}

	if pool.Status == req.NewStatus {
		return nil, http_util.NewHttpError(http.StatusBadRequest, "new status can not same as old")
	}

	taskErr := tasks.DefaultManage.NewTask(tasks.TaskTypePoolSwitch, &tasks.PoolSwitchTaskConfig{
		PoolId:    req.PoolId,
		NewStatus: req.NewStatus,
	})
	if taskErr != nil {
		log.GetLogger().Error("conduct pool switch task failed", zap.Error(taskErr), zap.Any("pool_id", req.PoolId), zap.String("new_status", req.NewStatus))
		return nil, http_util.NewHttpError(http.StatusInternalServerError, "change pool status failed. "+taskErr.Error())
	}

	log.GetLogger().Info("change pool status succeed by admin:" + strconv.Itoa(int(jwtClaims.AdminId)))

	return true, nil
}
