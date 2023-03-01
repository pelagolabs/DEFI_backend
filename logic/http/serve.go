package http

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"time"
	"veric-backend/internal/log"
	"veric-backend/logic/config"
)

var srv *http.Server

func StartAndServe() error {
	srv = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.Get().HTTP.Listen, config.Get().HTTP.Port),
		Handler: registerRouter(),
	}

	log.GetLogger().Info("http server start", zap.String("address", srv.Addr))

	return srv.ListenAndServe()
}

func Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_ = srv.Shutdown(ctx)
}
