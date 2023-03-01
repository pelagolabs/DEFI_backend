package http_util

import (
	"encoding/json"
	"net/http"
)

type HTTPMethod string
type MethodMap map[HTTPMethod]RespFunc

const (
	MethodGet     HTTPMethod = "GET"
	MethodCopy    HTTPMethod = "COPY"
	MethodHead    HTTPMethod = "HEAD"
	MethodPost    HTTPMethod = "POST"
	MethodPut     HTTPMethod = "PUT"
	MethodPatch   HTTPMethod = "PATCH"
	MethodDelete  HTTPMethod = "DELETE"
	MethodConnect HTTPMethod = "CONNECT"
	MethodOptions HTTPMethod = "OPTIONS"
	MethodTrace   HTTPMethod = "TRACE"
	MethodStatus  HTTPMethod = "STATUS"
)

func (m MethodMap) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Credentials", "true")

	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Add("Access-Control-Allow-Origin", origin)
	} else {
		w.Header().Add("Access-Control-Allow-Origin", "*")
	}

	if exposeHeaders := r.Header.Get("X-Expose-Headers"); exposeHeaders != "" {
		w.Header().Add("Access-Control-Expose-Headers", exposeHeaders)
	}

	if r.Method == "OPTIONS" {
		if header := r.Header.Get("Access-Control-Request-Headers"); header != "" {
			w.Header().Add("Access-Control-Allow-Headers", header)
		}

		if header := r.Header.Get("Access-Control-Request-Origin"); header != "" {
			w.Header().Add("Access-Control-Allow-Origin", header)
		}

		if header := r.Header.Get("Access-Control-Request-Method"); header != "" {
			w.Header().Add("Access-Control-Allow-Methods", header)
		}

		_, _ = w.Write([]byte("Allow"))
		return
	}

	var failJson = &failJson{Success: false, ErrMsg: "Method Not Allowed"}
	jsonWriter := json.NewEncoder(w)
	method := r.Method

	if xMethod := r.Header.Get("X-Method"); xMethod != "" {
		method = xMethod
	}

	if next, ok := m[HTTPMethod(method)]; !ok {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_ = jsonWriter.Encode(failJson)
		return
	} else {
		next(w, r)
	}
}
