package cron

import (
	"fmt"
	"go.uber.org/zap"
	"veric-backend/internal/log"
	"veric-backend/internal/util"
	"veric-backend/logic/config"
	"veric-backend/logic/db"
)

func init() {
	forceAddFunc("@every 5m", MonitorWalletPool)
}

func MonitorWalletPool() {
	// get all waller
	wallets, getWalletErr := db.AllWalletFromPool()
	if getWalletErr != nil {
		log.GetLogger().Error("[cron MonitorWalletPool] get all waller failed", zap.Error(getWalletErr))
		return
	}

	if len(wallets) == 0 {
		return
	}

	walletSet := make(map[string]*walletPoolItem, 0)
	walletStatSet := make(map[string]*walletPoolStatItem, 0)

	for _, wallet := range wallets {
		if _, ok := walletStatSet[wallet.ChainName]; !ok {
			walletStatSet[wallet.ChainName] = &walletPoolStatItem{
				Total: 0,
				Used:  0,
			}
		}

		// skip duplicate wallet
		if _, ok := walletSet[wallet.Addr]; ok {
			continue
		}

		walletSet[wallet.Addr] = &walletPoolItem{
			Address: wallet.Addr,
			Chain:   wallet.ChainName,
			IsStat:  false,
		}

		walletStatSet[wallet.ChainName].Total++
	}

	// get all handling payment
	payments, getPaymentErr := db.FindPaymentsByStatus([]string{db.PaymentStatusCreated, db.PaymentStatusPending})
	if getPaymentErr != nil {
		log.GetLogger().Error("[cron MonitorWalletPool] get all handling payment failed", zap.Error(getPaymentErr))
		return
	}

	for _, payment := range payments {
		if stat, ok := walletSet[payment.CollectionAddress]; ok && !stat.IsStat {
			walletStatSet[walletSet[payment.CollectionAddress].Chain].Used++
			walletSet[payment.CollectionAddress].IsStat = true
		}
	}

	for chain, statItem := range walletStatSet {
		usedRate := statItem.Used / statItem.Total
		fmt.Printf("chain: %s usedRate: %.2f statItem: %+v", chain, usedRate, statItem)

		if usedRate >= config.Get().Logic.WalletPoolLackThreshold {
			incidentKey := fmt.Sprintf("wallet not enough! - %s", chain)
			message := fmt.Sprintf(util.WalletPoolLackMessage, chain, chain, statItem.Total, statItem.Used, usedRate*100)
			severity := "error"

			pagerErr := util.SendPagerNotify(incidentKey, "trigger", message, severity)
			if pagerErr != nil {
				log.GetLogger().Error("[cron MonitorWalletPool] send pagerduty incident failed", zap.Error(pagerErr))
			}
		}
	}
}
