package http

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"veric-backend/logic/http/administration"
	"veric-backend/logic/http/merchant"
	"veric-backend/logic/http/open_api"
)

func registerRouter() http.Handler {
	r := mux.NewRouter()

	merchant.RegisterRouter(r.PathPrefix("/api/merchant").Subrouter())
	administration.RegisterRouter(r.PathPrefix("/api/admin").Subrouter())
	open_api.RegisterRouter(r.PathPrefix("/open_api").Subrouter())

	return handlers.LoggingHandler(os.Stdout, r)
}
