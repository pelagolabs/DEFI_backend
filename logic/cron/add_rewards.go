package cron

import (
	"go.uber.org/zap"
	"veric-backend/internal/log"
	"veric-backend/logic/blockchain/exchange"
	"veric-backend/logic/db"
	"veric-backend/logic/db/types"
	"veric-backend/logic/tasks"
)

// in usd cents
var minCollection = types.NewBigFloatFast(100)

func init() {
	forceAddFunc("0 4 * * *", DayAddRewards)
}

func DayAddRewards() {
	err := addRewards()
	if err != nil {
		log.GetLogger().Warn("[cron addRewards] addRewards failed", zap.Error(err))
	}
}

func addRewards() error {
	blockchains, err := db.AllBlockchain()
	if err != nil {
		return err
	}

	for _, blockchain := range blockchains {
		currencies, err := db.FindAllCurrencyByChain(blockchain.ID, db.WithPreload("Blockchain"))
		if err != nil {
			return err
		}

		for _, currency := range currencies {
			err = addRewardsByCurrency(currency)
			if err != nil {
				log.GetLogger().Warn("[cron addRewards] addRewards failed", zap.String("name", currency.Name), zap.Error(err))
			}
		}
	}

	return nil
}

func addRewardsByCurrency(currency *db.Currency) error {
	fees, err := db.FindAvailablePaymentFeesByFeeTypeAndCurrencyId(db.PaymentFeeFeeTypePool, currency.ID)
	if err != nil {
		return err
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
		return nil
	}

	usdAmount, err := exchange.DefaultManage.ExchangeCoinToUSD(totalClaimAmount.RawBigInt(), currency)
	if err == nil && types.NewBigFloat(usdAmount).Mul(types.NewBigFloatFast(100)).Cmp(minCollection) < 0 {
		return nil
	}

	conf := &tasks.AddRewardsTaskConfig{
		ManagerId:   0,
		Remark:      "System Auto Add Rewards",
		FeeIds:      feeIds,
		Chain:       currency.Blockchain.ContractName,
		Currency:    currency.Symbol,
		TotalAmount: totalClaimAmount.RawBigInt(),
	}

	return tasks.DefaultManage.NewTaskWithTx(tasks.TaskTypeAddRewards, conf, func(tx db.Options) error {
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
