package did

import (
	"veric-backend/internal/log"
	"veric-backend/logic/blockchain/eth"
	"veric-backend/logic/config"
)

var IssuePriKey *eth.PrivateKey
var IssueDID *DIDDocument

func init() {
	log.GetLogger().Info("init did issuer...")

	var err error
	IssuePriKey, err = eth.NewPrivateKey(config.Get().AccountPriKey.DIDIssuer)
	if err != nil {
		panic(err)
	}

	IssueDID, err = CreateUserDID(IssuePriKey.Address().String(), IssuePriKey)
}
