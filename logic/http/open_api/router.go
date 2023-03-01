package open_api

import (
	"github.com/gorilla/mux"
	"veric-backend/logic/http/http_util"
)

func RegisterRouter(r *mux.Router) {
	r.Handle("/payment", http_util.MethodMap{
		http_util.MethodGet:  http_util.SimpleWrap(FindFullPayment),
		http_util.MethodPost: http_util.AutoSimpleJsonBodyWrap(NewPayment),
	})
}
