package main

import (
	"go.uber.org/zap"
	netHttp "net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"veric-backend/internal/log"
	"veric-backend/logic/blockchain/contract"
	"veric-backend/logic/blockchain/events"
	"veric-backend/logic/config"
	"veric-backend/logic/cron"
	"veric-backend/logic/db"
	"veric-backend/logic/http"
	"veric-backend/logic/tasks"
)

func main() {
	if config.Get().Debug.Enable {
		time.Local = time.FixedZone("UTC", 0)
	}

	go func() {
		err := http.StartAndServe()
		if err != nil {
			if err == netHttp.ErrServerClosed {
				log.GetLogger().Info("http stopped")
				return
			} else {
				panic(err)
			}
		}
	}()
	defer http.Stop()

	cron.Start()
	tasks.DefaultManage.Start(true)

	allListen := make([]*events.Events, 0)
	blockchains, err := db.AllBlockchain()
	if err != nil {
		log.GetLogger().Panic("db find all blockchain fail", zap.Error(err))
		return
	}
	for _, blockchain := range blockchains {
		eventsListen, err := events.NewEvents(blockchain)
		if err != nil {
			panic(err)
		}

		err = contract.DefaultContract.InitDynamicChain(blockchain.Name, blockchain.Endpoint)
		if err != nil {
			log.GetLogger().Panic("contract init dynamic chain fail", zap.Error(err))
		}

		allListen = append(allListen, eventsListen)
	}

	defer func() {
		for _, listen := range allListen {
			listen.Close()
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-done

	log.GetLogger().Info("gracefully shutdown...")
}
