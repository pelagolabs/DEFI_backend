package merchant

import (
	"github.com/gorilla/mux"
	"reflect"
	"veric-backend/logic/config"
	"veric-backend/logic/http/http_util"
)

func RegisterRouter(r *mux.Router) {
	r.Handle("/sign_content", http_util.MethodMap{
		http_util.MethodGet: http_util.SimpleUrlQueryWrap(reflect.TypeOf(merchantSignContentRequest{}), GetMerchantSignContent),
	})
	r.Handle("/related_merchant", http_util.MethodMap{
		http_util.MethodGet: http_util.SimpleUrlQueryWrap(reflect.TypeOf(userRelatedMerchantRequest{}), GetUserRelatedMerchant),
	})
	r.Handle("/check_user", http_util.MethodMap{
		http_util.MethodGet: http_util.SimpleUrlQueryWrap(reflect.TypeOf(checkUserExistRequest{}), CheckUserExist),
	})

	r.Handle("/login", http_util.MethodMap{
		http_util.MethodPost: http_util.SimpleJsonBodyWrap(reflect.TypeOf(loginUseEthSignatureRequest{}), LoginUseEthSignature),
	})
	if config.Get().Debug.Enable {
		r.Handle("/login/fast_test", http_util.MethodMap{
			http_util.MethodPost: http_util.SimpleJsonBodyWrap(reflect.TypeOf(loginUseEthSignatureRequest{}), FastLoginForTest),
		})
	}

	r.Handle("/register", http_util.MethodMap{
		http_util.MethodPost: http_util.SimpleJsonBodyWrap(reflect.TypeOf(registerUseEthSignatureRequest{}), RegisterUseEthSignature),
	})

	r.Handle("/nickname", http_util.MethodMap{
		http_util.MethodGet:  http_util.SimpleUrlQueryWrap(reflect.TypeOf(userNicknameRequest{}), GetUserNickname),
		http_util.MethodPost: http_util.SimpleJsonBodyWrap(reflect.TypeOf(setNicknameUseEthSignatureRequest{}), SetNicknameUseEthSignature),
	})

	r.Handle("/config/all_currency", http_util.MethodMap{
		http_util.MethodGet: http_util.SimpleWrap(GetAvailableCurrency),
	})
	r.Handle("/config/exchange_rate", http_util.MethodMap{
		http_util.MethodGet: http_util.SimpleUrlQueryWrap(reflect.TypeOf(exchangeRateRequest{}), GetExchangeRate),
	})

	r.Handle("/summary/base", http_util.MethodMap{
		http_util.MethodGet: http_util.SimpleUrlQueryWrap(reflect.TypeOf(merchantBaseSummaryRequest{}), GetBaseSummary),
	})
	r.Handle("/chart/payment", http_util.MethodMap{
		http_util.MethodGet: http_util.SimpleUrlQueryWrap(reflect.TypeOf(merchantPaymentStatChartRequest{}), GetPaymentStatChart),
	})

	r.Handle("/application/{app_id}/detail", http_util.MethodMap{
		http_util.MethodGet: http_util.SimpleWrap(GetApplicationDetail),
	})
	r.Handle("/application/base", http_util.MethodMap{
		http_util.MethodPost: http_util.SimpleJsonBodyWrap(reflect.TypeOf(modifyApplicationBaseRequest{}), ModifyApplicationBase),
	})
	r.Handle("/application/currency", http_util.MethodMap{
		http_util.MethodPost: http_util.SimpleJsonBodyWrap(reflect.TypeOf(modifyApplicationCurrencyRequest{}), ModifyApplicationCurrency),
	})
	r.Handle("/application/api_key", http_util.MethodMap{
		http_util.MethodPost: http_util.SimpleJsonBodyWrap(reflect.TypeOf(modifyApplicationApiKeyRequest{}), ModifyApplicationApiKey),
	})
	r.Handle("/application/ipn_key", http_util.MethodMap{
		http_util.MethodPost: http_util.SimpleJsonBodyWrap(reflect.TypeOf(modifyApplicationIpnKeyRequest{}), ModifyApplicationIpnKey),
	})
	r.Handle("/application/slippage", http_util.MethodMap{
		http_util.MethodPost: http_util.SimpleJsonBodyWrap(reflect.TypeOf(modifyApplicationSlippageRequest{}), ModifyApplicationSlippage),
	})

	r.Handle("/application/{app_id}/payment/summary", http_util.MethodMap{
		http_util.MethodGet: http_util.SimpleUrlQueryWrap(reflect.TypeOf(paymentSummaryRequest{}), GetPaymentSummary),
	})
	r.Handle("/application/{app_id}/payment/list", http_util.MethodMap{
		http_util.MethodGet: http_util.SimpleUrlQueryWrap(reflect.TypeOf(paymentListRequest{}), GetPaymentList),
	})
	r.Handle("/application/{app_id}/payment/create", http_util.MethodMap{
		http_util.MethodPost: http_util.AutoSimpleJsonBodyWrap(NewPayment),
	})
	r.Handle("/application/{app_id}/payment/{payment_num}", http_util.MethodMap{
		http_util.MethodGet: http_util.SimpleWrap(GetPaymentDetail),
	})

	r.Handle("/application/{app_id}/withdraw/summary", http_util.MethodMap{
		http_util.MethodGet: http_util.SimpleUrlQueryWrap(reflect.TypeOf(withdrawSummaryRequest{}), GetWithdrawSummary),
	})
	r.Handle("/application/{app_id}/withdraw/list", http_util.MethodMap{
		http_util.MethodGet: http_util.SimpleUrlQueryWrap(reflect.TypeOf(withdrawListRequest{}), GetWithdrawList),
	})
	r.Handle("/application/{app_id}/withdraw/{withdraw_id}", http_util.MethodMap{
		http_util.MethodGet: http_util.SimpleWrap(GetWithdrawDetail),
	})

	r.Handle("/user/list", http_util.MethodMap{
		http_util.MethodGet: http_util.SimpleUrlQueryWrap(reflect.TypeOf(merchantUserListRequest{}), GetMerchantUserList),
	})
	r.Handle("/user/role/grant", http_util.MethodMap{
		http_util.MethodPost: http_util.SimpleJsonBodyWrap(reflect.TypeOf(grantUserMerchantRoleRequest{}), GrantUserMerchantRole),
	})
	r.Handle("/user/role/revoke", http_util.MethodMap{
		http_util.MethodPost: http_util.SimpleJsonBodyWrap(reflect.TypeOf(revokeUserMerchantRoleRequest{}), RevokeUserMerchantRole),
	})
	r.Handle("/user/role/change", http_util.MethodMap{
		http_util.MethodPost: http_util.SimpleJsonBodyWrap(reflect.TypeOf(changeUserMerchantRoleRequest{}), ChangeUserMerchantRole),
	})

	r.Handle("/vc", http_util.MethodMap{
		http_util.MethodGet:  http_util.SimpleWrap(GetMerchantVCList),
		http_util.MethodPost: http_util.AutoSimpleJsonBodyWrap(MarkVCReceived),
	})

	r.Handle("/vc/invalid", http_util.MethodMap{
		http_util.MethodPost: http_util.AutoSimpleJsonBodyWrap(MarkAllVCInvalid),
	})

	r.Handle("/vc/status", http_util.MethodMap{
		http_util.MethodPost: http_util.SimpleJsonBodyWrap(reflect.TypeOf(batchGetPaymentVcStatusRequest{}), BatchGetVcStatus),
	})

	r.Handle("/withdraw", http_util.MethodMap{
		http_util.MethodPost: http_util.AutoSimpleJsonBodyWrap(DoWithdraw),
	})
}
