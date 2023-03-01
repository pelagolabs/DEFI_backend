package administration

import (
	"github.com/golang-jwt/jwt"
	"github.com/kataras/hcaptcha"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
	"veric-backend/internal/log"
	"veric-backend/logic/config"
	"veric-backend/logic/db"
	"veric-backend/logic/http/http_util"
)

var (
	captchaStore = base64Captcha.DefaultMemStore
)

func LoginUsePwd(typ interface{}, r *http_util.HTTPContext) (resp interface{}, respErr error) {
	req := typ.(*loginUsePwdRequest)

	// check captcha
	_, captchaCheckOk := hcaptcha.Get(r.Request)
	if !captchaCheckOk {
		return nil, http_util.NewHttpError(http.StatusForbidden, "captcha code not correct")
	}

	manager, findErr := db.FindManagerWithPermissionByUserName(req.Username)
	if findErr != nil || manager.ID == 0 {
		return nil, http_util.NewHttpError(http.StatusForbidden, "manager not found or password not correct")
	}

	compErr := bcrypt.CompareHashAndPassword([]byte(manager.Password), []byte(req.Password))
	if compErr != nil {
		return nil, http_util.NewHttpError(http.StatusForbidden, "manager not found or password not correct")
	}

	// load manager permission
	privileges := make([]string, 0)
	for _, permission := range manager.Permissions {
		privileges = append(privileges, permission.Identity)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, http_util.AdminClaims{
		AdminId:    manager.ID,
		Privileges: privileges,
	})

	tokenStr, respErr := token.SignedString([]byte(config.Get().HTTP.JwtEncryptSecret))
	if respErr != nil {
		return nil, respErr
	}

	r.SetCookies(&http.Cookie{
		Name:     "admin-auth-token",
		Value:    tokenStr,
		Path:     "/",
		Expires:  time.Now().Add(12 * time.Hour),
		HttpOnly: true,
		SameSite: 0,
	})

	return &loginUsePwdResponse{
		Id:         manager.ID,
		Username:   manager.Username,
		Nickname:   manager.Nickname,
		Privileges: privileges,
		Token:      tokenStr,
	}, respErr
}

func CreateNewCaptcha(r *http_util.HTTPContext) (resp interface{}, err error) {
	driver := &base64Captcha.DriverString{
		Height:          60,
		Width:           240,
		NoiseCount:      0,
		Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
		Length:          6,
		ShowLineOptions: base64Captcha.OptionShowSlimeLine,
	}

	c := base64Captcha.NewCaptcha(driver, captchaStore)
	captchaId, captchaContent, captchaErr := c.Generate()

	if captchaErr != nil {
		return nil, http_util.NewHttpError(http.StatusInternalServerError, "captcha create failed")
	}

	return &createNewCaptchaResponse{
		CaptchaId:      captchaId,
		CaptchaContent: captchaContent,
	}, nil
}

func ChangePwd(typ interface{}, r *http_util.HTTPContext) (resp interface{}, respErr error) {
	req := typ.(*changePwdRequest)

	// get admin
	jwtClaims, jwtErr := r.GetAdminJwt("admin-auth-token")
	if jwtErr != nil {
		return nil, jwtErr
	}

	manager, findErr := db.FindManagerById(jwtClaims.AdminId)
	if findErr != nil || manager.ID == 0 {
		return nil, http_util.NewHttpError(http.StatusForbidden, "manager not found or not login")
	}

	if config.Get().Debug.Enable {
		if manager.ID == 1 {
			return nil, http_util.NewHttpError(http.StatusForbidden, "can not change admin account")
		}
	}

	// check old password
	compErr := bcrypt.CompareHashAndPassword([]byte(manager.Password), []byte(req.OldPassword))
	if compErr != nil {
		return nil, http_util.NewHttpError(http.StatusForbidden, "old password not correct")
	}

	// calc new password
	hash, generateErr := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.MinCost)
	if generateErr != nil {
		log.GetLogger().Error("calc new password failed.", zap.String("error", generateErr.Error()))
		return nil, http_util.NewHttpError(http.StatusInternalServerError, "change password failed.  please connect to admin")
	}

	manager.Password = string(hash)

	saveErr := db.SaveManager(manager)
	if saveErr != nil {
		log.GetLogger().Error("save manager failed.", zap.String("error", saveErr.Error()))
		return nil, http_util.NewHttpError(http.StatusInternalServerError, "save manager failed.  please connect to admin")
	}

	return true, nil
}

func FastLoginForTest(typ interface{}, r *http_util.HTTPContext) (resp interface{}, respErr error) {
	req := typ.(*loginUsePwdRequest)

	manager, findErr := db.FindManagerWithPermissionByUserName(req.Username)
	if findErr != nil || manager.ID == 0 {
		return nil, http_util.NewHttpError(http.StatusForbidden, "manager not found or password not correct")
	}

	compErr := bcrypt.CompareHashAndPassword([]byte(manager.Password), []byte(req.Password))
	if compErr != nil {
		return nil, http_util.NewHttpError(http.StatusForbidden, "manager not found or password not correct")
	}

	// load manager permission
	privileges := make([]string, 0)
	for _, permission := range manager.Permissions {
		privileges = append(privileges, permission.Identity)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, http_util.AdminClaims{
		AdminId:    manager.ID,
		Privileges: privileges,
	})

	tokenStr, respErr := token.SignedString([]byte(config.Get().HTTP.JwtEncryptSecret))
	if respErr != nil {
		return nil, respErr
	}

	r.SetCookies(&http.Cookie{
		Name:     "admin-auth-token",
		Value:    tokenStr,
		Path:     "/",
		Expires:  time.Now().Add(12 * time.Hour),
		HttpOnly: true,
		SameSite: 0,
	})

	return &loginUsePwdResponse{
		Id:         manager.ID,
		Username:   manager.Username,
		Nickname:   manager.Nickname,
		Privileges: privileges,
		Token:      tokenStr,
	}, respErr
}
