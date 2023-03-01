package http_util

import (
	"encoding/json"
	lru "github.com/hashicorp/golang-lru"
	"github.com/kataras/hcaptcha"
	"net/http"
	"strings"
	"veric-backend/logic/blockchain/eth"
	"veric-backend/logic/config"
)

const SignContent = "sign in"

var (
	userCache, _   = lru.New(100)
	hCaptchaClient *hcaptcha.Client
)

func GetUserPublicKeyFromSign(sign string) (*eth.PublicKey, error) {
	if cached, ok := userCache.Get(sign); ok {
		return cached.(*eth.PublicKey), nil
	} else {
		publicKey, err := eth.GetPublicKeyUseEthSign(SignContent, sign)
		if err != nil {
			return nil, err
		}

		userCache.Add(sign, publicKey)
		return publicKey, nil
	}
}

func GetUserPublicKeyFromSignWithCustomContent(content, sign string) (*eth.PublicKey, error) {
	if cached, ok := userCache.Get(sign); ok {
		return cached.(*eth.PublicKey), nil
	} else {
		publicKey, err := eth.GetPublicKeyUseEthSign(content, sign)
		if err != nil {
			return nil, err
		}

		userCache.Add(sign, publicKey)
		return publicKey, nil
	}
}

func GetHCaptchaClient() *hcaptcha.Client {
	if hCaptchaClient == nil {
		hCaptchaClient = hcaptcha.New(config.Get().HCaptcha.SecretKey)
		hCaptchaClient.FailureHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			hcaptchaRes, _ := hcaptcha.Get(r)
			resJson := &failJson{
				Success: false,
				ErrCode: http.StatusForbidden,
				ErrMsg:  strings.Join(hcaptchaRes.ErrorCodes, " / "),
			}
			jsonWriter := json.NewEncoder(w)

			w.WriteHeader(int(resJson.ErrCode))
			_ = jsonWriter.Encode(resJson)
		})
	}

	return hCaptchaClient
}
