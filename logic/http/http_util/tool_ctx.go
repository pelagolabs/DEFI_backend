package http_util

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strconv"
	"veric-backend/logic/blockchain/eth"
	"veric-backend/logic/config"
)

type HTTPContext struct {
	*http.Request

	w http.ResponseWriter
}

type MerchantClaims struct {
	UserId     uint   `json:"user_id"`
	MerchantId uint   `json:"merchant_id"`
	Role       string `json:"role"`
	jwt.StandardClaims
}

type AdminClaims struct {
	AdminId    uint     `json:"admin_id"`
	Privileges []string `json:"privileges"`
	jwt.StandardClaims
}

func NewHTTPContext(r *http.Request, w http.ResponseWriter) *HTTPContext {
	return &HTTPContext{Request: r, w: w}
}

func (c *HTTPContext) GetHeaderPublicKey() (*eth.PublicKey, error) {
	signData := c.Header.Get("X-Token")
	if signData == "" {
		return nil, NewHttpError(0xE000001, "token not exists")
	}

	return GetUserPublicKeyFromSign(signData)
}

func (c *HTTPContext) QueryWithDefault(key, def string) string {
	query := c.URL.Query()
	if query.Has(key) {
		return query.Get(key)
	}

	return def
}

func (c *HTTPContext) QueryWithDefaultInt(key string, def int) int {
	query := c.URL.Query()
	if query.Has(key) {
		queryNum, err := strconv.Atoi(query.Get(key))
		if err != nil {
			return def
		}

		return queryNum
	}

	return def
}

func (c *HTTPContext) SetCookies(cookie *http.Cookie) {
	http.SetCookie(c.w, cookie)
}

func (c *HTTPContext) GetMerchantJwt(cookieName string) (*MerchantClaims, error) {
	cookie, err := c.Cookie(cookieName)
	if err != nil {
		return nil, err
	}

	token, err := jwt.ParseWithClaims(cookie.Value, &MerchantClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.Get().HTTP.JwtEncryptSecret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*MerchantClaims); ok && token.Valid {
		return claims, err
	}

	return nil, errors.New("auth error")
}

func (c *HTTPContext) GetAdminJwt(cookieName string) (*AdminClaims, error) {
	cookie, err := c.Cookie(cookieName)
	if err != nil {
		return nil, err
	}

	token, err := jwt.ParseWithClaims(cookie.Value, &AdminClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.Get().HTTP.JwtEncryptSecret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*AdminClaims); ok && token.Valid {
		return claims, err
	}

	return nil, errors.New("auth error")
}

func (c *HTTPContext) CheckPermissionAny(hasPermission, needPermission []string) bool {
	m := make(map[string]int)
	nn := make([]string, 0)
	for _, v := range hasPermission {
		m[v]++
	}

	for _, v := range needPermission {
		times, _ := m[v]
		if times >= 1 {
			nn = append(nn, v)
		}
	}

	return len(nn) > 0
}
