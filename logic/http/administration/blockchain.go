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

func GetBlockchainList(typ interface{}, r *http_util.HTTPContext) (resp interface{}, respErr error) {
	status := r.QueryWithDefault("status", "all")
	page := r.QueryWithDefaultInt("page", 1)
	size := r.QueryWithDefaultInt("size", 10)

	// get admin
	jwtClaims, jwtErr := r.GetAdminJwt("admin-auth-token")
	if jwtErr != nil {
		return nil, http_util.NewHttpError(http.StatusForbidden, "need permission for this operate")
	}

	if !r.CheckPermissionAny(jwtClaims.Privileges, []string{"currency"}) {
		return nil, http_util.NewHttpError(http.StatusForbidden, "no permission for this operation")
	}

	// calc pagination
	offset := (page - 1) * size

	// get transaction list
	var (
		blockchains []db.Blockchain
		total       uint
	)

	total = db.GetBlockchainCountByStatus(status)
	blockchains, findErr := db.FindPaginationBlockchainByStatus(status, offset, size)
	if findErr != nil {
		log.GetLogger().Error("get blockchain list failed", zap.String("error", findErr.Error()))
		return nil, http_util.NewHttpError(http.StatusInternalServerError, "get blockchain list failed")
	}

	return &blockchainListResponse{
		Total:       total,
		Blockchains: blockchains,
	}, nil
}

func ChangeBlockchainStatus(typ interface{}, r *http_util.HTTPContext) (resp interface{}, saveErr error) {
	req := typ.(*changeBlockchainStatusRequest)

	// get admin
	jwtClaims, jwtErr := r.GetAdminJwt("admin-auth-token")
	if jwtErr != nil {
		return nil, jwtErr
	}

	if !r.CheckPermissionAny(jwtClaims.Privileges, []string{"currency"}) {
		return nil, http_util.NewHttpError(http.StatusForbidden, "no permission for this operation")
	}

	// get blockchain
	blockchain, saveErr := db.FindBlockchainById(req.ChainId)
	if saveErr != nil || blockchain.ID == 0 {
		return nil, http_util.NewHttpError(http.StatusNotFound, "blockchain not exist")
	}

	if blockchain.Status == req.NewStatus {
		return nil, http_util.NewHttpError(http.StatusBadRequest, "new status can not same as old")
	}

	// update status
	blockchain.Status = req.NewStatus

	saveErr = db.SaveBlockchain(blockchain)
	if saveErr != nil {
		log.GetLogger().Error("save blockchain failed", zap.String("error", saveErr.Error()))
		return nil, http_util.NewHttpError(http.StatusInternalServerError, "save blockchain failed")
	}

	taskErr := tasks.DefaultManage.NewTask(tasks.TaskTypeChainSwitch, &tasks.ChainSwitchTaskConfig{
		ChainId:   req.ChainId,
		NewStatus: req.NewStatus,
	})
	if taskErr != nil {
		log.GetLogger().Error("conduct chain switch task failed", zap.Error(taskErr), zap.Any("chain_id", req.ChainId), zap.String("new_status", req.NewStatus))
		return nil, http_util.NewHttpError(http.StatusInternalServerError, "change pool status failed. "+taskErr.Error())
	}

	log.GetLogger().Info("change blockchain status succeed by admin:" + strconv.Itoa(int(jwtClaims.AdminId)))

	return true, nil
}
