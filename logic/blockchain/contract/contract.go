package contract

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
	"math/big"
	"math/rand"
	"sync"
	"time"
	"veric-backend/internal/log"
	"veric-backend/internal/util"
	"veric-backend/logic/blockchain/contract/chain/dynamic"
	"veric-backend/logic/blockchain/contract/chain/ethereum"
	"veric-backend/logic/blockchain/contract/chain/harmony"
	"veric-backend/logic/config"
	"veric-backend/logic/db"
)

var DefaultContract *Contract

func init() {
	log.GetLogger().Info("init contract info...")

	var err error
	DefaultContract, err = NewContract()
	if err != nil {
		panic(err)
	}
}

type Contract struct {
	conf      *config.ContractConf
	walletMap util.SyncedMap[string, bool]

	dynamicLock sync.Mutex

	ethereumChain *ethereum.EthereumChain
	harmonyChain  *harmony.HarmonyChain
	dynamicChain  util.SyncedMap[string, *dynamic.DynamicChain]
}

func NewContract() (*Contract, error) {
	c := &Contract{
		conf: config.Get().Contract,
	}

	err := c.init()
	if err != nil {
		return nil, err
	}

	c.updateWalletMapOnce()
	go c.updateWalletMap()

	return c, nil
}

func (c *Contract) init() (err error) {
	if c.conf == nil {
		log.GetLogger().Panic("contract config is nil")
	}

	c.harmonyChain, err = harmony.NewHarmonyChain(c.conf)
	if err != nil {
		return err
	}

	c.ethereumChain, err = ethereum.NewEthereumChain(c.conf)
	if err != nil {
		return err
	}

	return nil
}

func (c *Contract) InitDynamicChain(name string, endpoint string) error {
	_, err := c.getDynamicChain(name, endpoint)
	return err
}

func (c *Contract) getDynamicChain(name string, endpoint string) (chain *dynamic.DynamicChain, err error) {
	if chain, ok := c.dynamicChain.Load(name); ok {
		return chain, err
	} else {
		c.dynamicLock.Lock()
		defer c.dynamicLock.Unlock()

		if chain, ok := c.dynamicChain.Load(name); ok {
			return chain, err
		}

		dynamicChain, err := dynamic.NewDynamicChain(c.conf, name, endpoint)
		if err != nil {
			return nil, err
		}

		c.dynamicChain.Store(name, dynamicChain)
		return dynamicChain, nil
	}
}

func (c *Contract) updateWalletMap() {
	for range time.Tick(time.Minute) {
		c.updateWalletMapOnce()
	}
}

func (c *Contract) updateWalletMapOnce() {
	wallets, err := db.AllWalletFromPool()
	if err != nil {
		log.GetLogger().Warn("refresh wallet fail", zap.Error(err))
	}
	for _, wallet := range wallets {
		cacheKey := fmt.Sprintf("%s-%s", wallet.ChainName, common.HexToAddress(wallet.Addr))
		c.walletMap.Store(cacheKey, true)
	}
}

func (c *Contract) IsContractWalletAddress(chainName string, addr common.Address) bool {
	cacheKey := fmt.Sprintf("%s-%s", chainName, addr.String())
	result, _ := c.walletMap.Load(cacheKey)
	return result
}

func (c *Contract) getCurrencyBlockchainName(currency *db.Currency) string {
	if currency.Blockchain.ID > 0 {
		return currency.Blockchain.ContractName
	} else {
		blockchain, err := db.FindBlockchainById(currency.ChainId)
		if err != nil {
			return ""
		}

		return blockchain.ContractName
	}
}

func (c *Contract) ReturnAddressToPool(currency *db.Currency, address string) error {
	return c.harmonyChain.ReleaseVaultProcess().RunAsync(&harmony.VaultInfo{
		ChainName: c.getCurrencyBlockchainName(currency),
		Address:   address,
	})
}

func (c *Contract) GetAddressFromPool(currency *db.Currency) (string, error) {
	result, err := c.harmonyChain.UseVaultProcess().RunSyncResultWithTimeout(5*time.Minute, &harmony.VaultInfo{
		ChainName: c.getCurrencyBlockchainName(currency),
		RandomId:  big.NewInt(rand.Int63()).String(),
	})
	if err != nil {
		return "", err
	}

	if result.Address == "" {
		return "", errors.New("no available vault")
	}

	return result.Address, nil
}

func (c *Contract) QueryEnoughWalletsMaxAmount(wallets []*db.WalletPool, allowLP bool, currency *db.Currency, needAmount *big.Int) (wallet string, err error) {
	return c.ethereumChain.QueryEnoughWalletsMaxAmount(wallets, allowLP, currency, needAmount)
}

func (c *Contract) QueryEnoughWalletsMaxNativeAmount(wallets []*db.WalletPool, allowLP bool, currency *db.Currency, needAmount *big.Int) (wallet string, err error) {
	return c.ethereumChain.QueryEnoughWalletsMaxNativeAmount(wallets, allowLP, currency, needAmount)
}

func (c *Contract) WithdrawNative(info *ethereum.WithdrawInfo) error {
	return c.ethereumChain.WithdrawNativeProcess().RunAsync(info)
}

func (c *Contract) Withdraw(info *ethereum.WithdrawInfo) error {
	return c.ethereumChain.WithdrawProcess().RunAsync(info)
}

func (c *Contract) DidRegistry(did, pubKey string) error {
	return c.harmonyChain.DidRegistryProcess().RunAsync(&harmony.DidInfo{
		Did:    did,
		PubKey: pubKey,
	})
}

func (c *Contract) LpSwitch(lpAddress, currencyAddress, action string, isNative, isGlobal bool) error {
	if action != "pause" && action != "unpause" {
		return errors.New("invalid action")
	}

	return c.ethereumChain.LpSwitchProcess().RunAsync(&ethereum.LpSwitchInfo{
		LpAddress:        common.HexToAddress(lpAddress),
		CurrencyAddress:  common.HexToAddress(currencyAddress),
		IsNativeCurrency: isNative,
		IsGlobal:         isGlobal,
		Action:           action,
	})
}

func (c *Contract) FarmSwitch(farmAddress, action string) error {
	if action != "pause" && action != "unpause" {
		return errors.New("invalid action")
	}

	return c.ethereumChain.FarmSwitchProcess().RunAsync(&ethereum.FarmSwitchInfo{
		FarmAddress: common.HexToAddress(farmAddress),
		Action:      action,
	})
}

func (c *Contract) AddRewards(info *dynamic.AddRewardsInfo) error {
	chain, err := c.getDynamicChain(info.BlockchainName, info.BlockchainEndpoint)
	if err != nil {
		return err
	}

	return chain.AddRewardsProcess().RunAsync(info)
}
