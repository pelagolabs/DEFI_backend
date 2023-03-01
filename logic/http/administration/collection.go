package administration

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
	"net/http"
	"veric-backend/internal/log"
	"veric-backend/logic/blockchain/exchange"
	"veric-backend/logic/config"
	"veric-backend/logic/db"
	"veric-backend/logic/db/types"
	"veric-backend/logic/http/http_util"
	"veric-backend/logic/tasks"
)

func GetAllRemainingFeeList(r *http_util.HTTPContext) (resp interface{}, respErr error) {
	jwtClaims, jwtErr := r.GetAdminJwt("admin-auth-token")
	if jwtErr != nil {
		return nil, http_util.NewHttpError(http.StatusForbidden, "need permission for this operate")
	}

	if !r.CheckPermissionAny(jwtClaims.Privileges, []string{"withdraw"}) {
		return nil, http_util.NewHttpError(http.StatusForbidden, "no permission for this operation")
	}

	currencies, currencyErr := db.FindAvailableCurrencies(db.WithPreload("Blockchain"))
	if currencyErr != nil {
		return nil, http_util.NewHttpError(http.StatusInternalServerError, "get currency currency failed. please try again later")
	}

	response := &allRemainingFeeListResponse{
		Total: uint(len(currencies)),
		Fees:  make([]*remainingFeeItem, 0),
	}

	feeAmountSet := make(map[uint]*remainingFeeAmountItem, 0)
	for _, currency := range currencies {
		feeAmountSet[currency.ID] = &remainingFeeAmountItem{
			TotalFeeRemainingAmount:    types.NewBigIntZero(),
			PlatformFeeRemainingAmount: types.NewBigIntZero(),
			PoolFeeRemainingAmount:     types.NewBigIntZero(),
		}
	}

	fees, err := db.FindAllAvailablePaymentFeesByFeeType([]string{"Platform", "Pool"}, db.WithPreload("Currency"))
	if err != nil {
		return nil, err
	}

	for _, fee := range fees {
		canClaimedAmount := fee.TotalAmount.Copy().Sub(fee.ClaimedAmount).Sub(fee.FrozenAmount)
		if canClaimedAmount.IsZero() {
			continue
		}

		if _, ok := feeAmountSet[fee.CurrencyId]; !ok {
			continue
		}

		switch fee.FeeType {
		case "Platform":
			feeAmountSet[fee.CurrencyId].PlatformFeeRemainingAmount.Add(canClaimedAmount)
		case "Pool":
			feeAmountSet[fee.CurrencyId].PoolFeeRemainingAmount.Add(canClaimedAmount)
		}

		feeAmountSet[fee.CurrencyId].TotalFeeRemainingAmount.Add(canClaimedAmount)
	}

	for _, currency := range currencies {
		platformAmountRaw, exchangePlatformErr := exchange.DefaultManage.ExchangeCoinToUSD(feeAmountSet[currency.ID].PlatformFeeRemainingAmount.RawBigInt(), &currency)
		poolAmountRaw, exchangePoolErr := exchange.DefaultManage.ExchangeCoinToUSD(feeAmountSet[currency.ID].PoolFeeRemainingAmount.RawBigInt(), &currency)

		//balanceRaw, exchangeErr := exchange.DefaultManage.ExchangeCoinToUSD(feeAmountSet[currency.ID].TotalFeeRemainingAmount.RawBigInt(), &currency)
		//if exchangeErr != nil {
		//	continue
		//}
		//
		//balance, _ := balanceRaw.Float64()

		if exchangePlatformErr != nil || exchangePoolErr != nil {
			continue
		}

		platformBalance, _ := platformAmountRaw.Float64()
		poolBalance, _ := poolAmountRaw.Float64()

		feeItem := &remainingFeeItem{
			ChainName:                currency.Blockchain.Name,
			ChainImageUrl:            currency.Blockchain.ImageUrl,
			CurrencyId:               currency.ID,
			CurrencyName:             currency.Name,
			CurrencySymbol:           currency.Symbol,
			CurrencyImageUrl:         currency.ImageUrl,
			AvailablePlatformBalance: platformBalance,
			AvailablePoolBalance:     poolBalance,
		}

		response.Fees = append(response.Fees, feeItem)
	}

	return response, nil
}

func GetCollectionHistoryList(typ interface{}, r *http_util.HTTPContext) (resp interface{}, respErr error) {
	page := r.QueryWithDefaultInt("page", 1)
	size := r.QueryWithDefaultInt("size", 10)
	currencyId := uint(r.QueryWithDefaultInt("currency_id", 0))
	date := r.QueryWithDefault("date", "")

	// check permission
	jwtClaims, jwtErr := r.GetAdminJwt("admin-auth-token")
	if jwtErr != nil {
		return nil, http_util.NewHttpError(http.StatusForbidden, "need permission for this operate")
	}

	if !r.CheckPermissionAny(jwtClaims.Privileges, []string{"withdraw"}) {
		return nil, http_util.NewHttpError(http.StatusForbidden, "no permission for this operation")
	}

	// calc pagination
	offset := (page - 1) * size

	// get transaction list
	var (
		collections = make([]collectionHistoryListItem, 0)
		total       uint
	)

	total = db.GetFeeWithdrawLogCountByStatusCurrencyDate(currencyId, date)
	collectionSet, findErr := db.FindPaginationFeeWithdrawLogByStatusCurrencyDate(currencyId, date, offset, size)
	if findErr != nil {
		log.GetLogger().Error("get collection log list failed", zap.String("error", findErr.Error()))
		return nil, http_util.NewHttpError(http.StatusInternalServerError, "get withdraw history list failed")
	}

	for _, collectionItem := range collectionSet {
		amount, _ := exchange.DefaultManage.ExchangeCoinToNaturalAmount(collectionItem.Amount.RawBigInt(), collectionItem.Currency).Float64()

		collections = append(collections, collectionHistoryListItem{
			ChainName:      collectionItem.Currency.Blockchain.Name,
			CurrencySymbol: collectionItem.Currency.Symbol,
			Amount:         amount,
			CreatedAt:      collectionItem.CreatedAt,
			TxHash:         collectionItem.TxHash,
		})
	}

	return &collectionHistoryListResponse{
		Total:       total,
		Collections: collections,
	}, nil
}

func DoCollection(req *doCollectionRequest, r *http_util.HTTPContext) (resp interface{}, err error) {
	jwtClaims, jwtErr := r.GetAdminJwt("admin-auth-token")
	if jwtErr != nil {
		return nil, http_util.NewHttpError(http.StatusForbidden, "need permission for this operate")
	}

	if !r.CheckPermissionAny(jwtClaims.Privileges, []string{"withdraw"}) {
		return nil, http_util.NewHttpError(http.StatusForbidden, "no permission for this operation")
	}

	currency, err := db.FindAvailableCurrencyById(req.CurrencyId, db.WithPreload("Blockchain"))
	if err != nil {
		return nil, err
	} else if currency.ID != req.CurrencyId {
		return nil, errors.New("currency invalid")
	}

	fees, err := db.FindAvailablePaymentFeesByFeeTypeAndCurrencyId(req.FeeType, req.CurrencyId)
	if err != nil {
		return nil, err
	}

	totalClaimAmount := types.NewBigIntZero()
	updateFees := make([]*db.PaymentFee, 0)
	feeIds := make([]uint, 0)
	for _, fee := range fees {
		if !fee.FrozenAmount.IsZero() {
			continue
		}

		canClaimedAmount := fee.TotalAmount.Copy().Sub(fee.ClaimedAmount).Sub(fee.FrozenAmount)
		if canClaimedAmount.IsZero() {
			continue
		}

		totalClaimAmount.Add(canClaimedAmount)
		updateFees = append(updateFees, fee)
		feeIds = append(feeIds, fee.ID)
	}

	if totalClaimAmount.IsZero() {
		return nil, errors.New("there is no available claim amount")
	}

	conf := &tasks.CollectionTaskConfig{
		ManagerId:   jwtClaims.AdminId,
		Remark:      req.Remark,
		FeeType:     req.FeeType,
		FeeIds:      feeIds,
		ToAddress:   common.HexToAddress(config.Get().Logic.AdministrationCollectionAddr),
		Chain:       currency.Blockchain.ContractName,
		Currency:    currency.Symbol,
		TotalAmount: totalClaimAmount.RawBigInt(),
	}

	return nil, tasks.DefaultManage.NewTaskWithTx(tasks.TaskTypeCollection, conf, func(tx db.Options) error {
		for _, fee := range updateFees {
			canClaimedAmount := fee.TotalAmount.Copy().Sub(fee.ClaimedAmount).Sub(fee.FrozenAmount)
			if canClaimedAmount.IsZero() {
				continue
			}

			err = db.UpdatePaymentFeeFrozenAmount(fee, fee.FrozenAmount.Copy().Add(canClaimedAmount), tx)
			if err != nil {
				return err
			}
		}

		return nil
	})
}
