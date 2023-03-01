package http_util

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/schema"
	"net/http"
	"reflect"
)

var validate = validator.New()
var urlDecode = schema.NewDecoder()

type SimpleWrapFunc func(r *HTTPContext) (resp interface{}, err error)
type SimpleInterfaceWrapFunc func(typ interface{}, r *HTTPContext) (resp interface{}, err error)
type AutoSimpleInterfaceWrapFunc[T any] func(typ *T, r *HTTPContext) (resp interface{}, err error)
type RespFunc func(writer http.ResponseWriter, request *http.Request)

type Empty struct{}

func (r RespFunc) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	r(writer, request)
}

type successJson struct {
	Success bool        `json:"success"`
	Msg     interface{} `json:"msg"`
}

type failJson struct {
	Success bool   `json:"success"`
	ErrCode uint32 `json:"err_code"`
	ErrMsg  string `json:"err_msg"`
}

func SimpleWrap(api SimpleWrapFunc) RespFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var successJson = &successJson{Success: true}
		var failJson = &failJson{Success: false}

		jsonWriter := json.NewEncoder(w)
		successJson.Msg, err = api(NewHTTPContext(r, w))
		if err != nil {
			if httpErr, ok := err.(*HTTPError); ok {
				failJson.ErrCode = httpErr.errorCode
				failJson.ErrMsg = httpErr.error.Error()
			} else {
				failJson.ErrCode = 500
				failJson.ErrMsg = err.Error()
			}

			if failJson.ErrCode >= 400 && failJson.ErrCode < 600 {
				w.WriteHeader(int(failJson.ErrCode))
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}

			_ = jsonWriter.Encode(failJson)
		} else {
			w.Header().Add("Content-Type", "application/json; charset=utf-8")

			_ = jsonWriter.Encode(successJson)
		}
	}
}

func AutoSimpleJsonBodyWrap[T any](api AutoSimpleInterfaceWrapFunc[T]) RespFunc {
	typ := reflect.TypeOf(new(T)).Elem()
	return SimpleWrap(func(r *HTTPContext) (resp interface{}, err error) {
		if r.ContentLength <= 0 {
			return nil, fmt.Errorf("need a `%s` body", typ.String())
		}

		val := reflect.New(typ).Interface()

		err = json.NewDecoder(r.Body).Decode(val)
		if err != nil {
			return nil, fmt.Errorf("decode `%s` body err: %v", typ.String(), err)
		}

		err = validate.Struct(val)
		if err != nil {
			return nil, err
		}

		return api(val.(*T), r)
	})
}

func SimpleJsonBodyWrap(typ reflect.Type, api SimpleInterfaceWrapFunc) RespFunc {
	return SimpleWrap(func(r *HTTPContext) (resp interface{}, err error) {
		if r.ContentLength <= 0 {
			return nil, fmt.Errorf("need a `%s` body", typ.String())
		}

		val := reflect.New(typ).Interface()

		err = json.NewDecoder(r.Body).Decode(&val)
		if err != nil {
			return nil, fmt.Errorf("decode `%s` body err: %v", typ.String(), err)
		}

		err = validate.Struct(val)
		if err != nil {
			return nil, err
		}

		return api(val, r)
	})
}

func SimpleUrlQueryWrap(typ reflect.Type, api SimpleInterfaceWrapFunc) RespFunc {
	return SimpleWrap(func(r *HTTPContext) (resp interface{}, err error) {
		val := reflect.New(typ).Interface()

		err = urlDecode.Decode(val, r.URL.Query())
		if err != nil {
			return nil, fmt.Errorf("decode `%s` query err: %v", typ.String(), err)
		}

		err = validate.Struct(val)
		if err != nil {
			return nil, err
		}

		return api(val, r)
	})
}
