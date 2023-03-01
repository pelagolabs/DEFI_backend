package administration

import (
	"github.com/gorilla/mux"
	"reflect"
	"veric-backend/logic/config"
	"veric-backend/logic/http/http_util"
)

func RegisterRouter(r *mux.Router) {
	r.Handle("/login", http_util.GetHCaptchaClient().Handler(http_util.MethodMap{
		http_util.MethodPost: http_util.SimpleJsonBodyWrap(reflect.TypeOf(loginUsePwdRequest{}), LoginUsePwd),
	}))

	if config.Get().Debug.Enable {
		r.Handle("/login/fast_test", http_util.MethodMap{
			http_util.MethodPost: http_util.SimpleJsonBodyWrap(reflect.TypeOf(loginUsePwdRequest{}), FastLoginForTest),
		})
	}

	r.Handle("/login/captcha/create", http_util.MethodMap{
		http_util.MethodGet: http_util.SimpleWrap(CreateNewCaptcha),
	})
	r.Handle("/dashboard/overview", http_util.MethodMap{
		http_util.MethodGet: http_util.SimpleUrlQueryWrap(reflect.TypeOf(globalOverviewRequest{}), GetGlobalOverview),
	})

	r.Handle("/expense/summary", http_util.MethodMap{
		http_util.MethodGet: http_util.SimpleUrlQueryWrap(reflect.TypeOf(expenseSummaryRequest{}), GetExpenseSummary),
	})
	r.Handle("/expense/list", http_util.MethodMap{
		http_util.MethodGet: http_util.SimpleUrlQueryWrap(reflect.TypeOf(expenseListRequest{}), GetExpenseList),
	})

	r.Handle("/merchant/list", http_util.MethodMap{
		http_util.MethodGet: http_util.SimpleUrlQueryWrap(reflect.TypeOf(merchantListRequest{}), GetMerchantList),
	})
	r.Handle("/merchant/status", http_util.MethodMap{
		http_util.MethodPost: http_util.SimpleJsonBodyWrap(reflect.TypeOf(changeMerchantStatusRequest{}), ChangeMerchantStatus),
	})
	r.Handle("/merchant/suspend_all", http_util.MethodMap{
		http_util.MethodPost: http_util.SimpleJsonBodyWrap(reflect.TypeOf(suspendAllMerchantRequest{}), SuspendAllMerchant),
	})

	r.Handle("/application/list", http_util.MethodMap{
		http_util.MethodGet: http_util.SimpleUrlQueryWrap(reflect.TypeOf(applicationListRequest{}), GetApplicationList),
	})
	r.Handle("/application/status", http_util.MethodMap{
		http_util.MethodPost: http_util.SimpleJsonBodyWrap(reflect.TypeOf(changeApplicationStatusRequest{}), ChangeApplicationStatus),
	})

	r.Handle("/pool/list", http_util.MethodMap{
		http_util.MethodGet: http_util.SimpleUrlQueryWrap(reflect.TypeOf(poolListRequest{}), GetPoolList),
	})
	r.Handle("/pool/status", http_util.MethodMap{
		http_util.MethodPost: http_util.SimpleJsonBodyWrap(reflect.TypeOf(changePoolStatusRequest{}), ChangePoolStatus),
	})

	r.Handle("/blockchain/list", http_util.MethodMap{
		http_util.MethodGet: http_util.SimpleUrlQueryWrap(reflect.TypeOf(blockchainListRequest{}), GetBlockchainList),
	})
	r.Handle("/blockchain/status", http_util.MethodMap{
		http_util.MethodPost: http_util.SimpleJsonBodyWrap(reflect.TypeOf(changeBlockchainStatusRequest{}), ChangeBlockchainStatus),
	})

	r.Handle("/currency/list", http_util.MethodMap{
		http_util.MethodGet: http_util.SimpleUrlQueryWrap(reflect.TypeOf(currencyListRequest{}), GetCurrencyList),
	})
	r.Handle("/currency/status", http_util.MethodMap{
		http_util.MethodPost: http_util.SimpleJsonBodyWrap(reflect.TypeOf(changeCurrencyStatusRequest{}), ChangeCurrencyStatus),
	})

	r.Handle("/manager/list", http_util.MethodMap{
		http_util.MethodGet: http_util.SimpleUrlQueryWrap(reflect.TypeOf(managerListRequest{}), GetManagerList),
	})
	r.Handle("/manager/create", http_util.MethodMap{
		http_util.MethodPost: http_util.SimpleJsonBodyWrap(reflect.TypeOf(createManagerRequest{}), CreateManager),
	})
	r.Handle("/manager/change", http_util.MethodMap{
		http_util.MethodPost: http_util.SimpleJsonBodyWrap(reflect.TypeOf(changeManagerRequest{}), ChangeManager),
	})
	r.Handle("/manager/delete", http_util.MethodMap{
		http_util.MethodPost: http_util.SimpleJsonBodyWrap(reflect.TypeOf(deleteManagerRequest{}), DeleteManager),
	})
	r.Handle("/permission/all", http_util.MethodMap{
		http_util.MethodGet: http_util.SimpleWrap(GetAllPermission),
	})
	r.Handle("/manager/change_pwd", http_util.MethodMap{
		http_util.MethodPost: http_util.SimpleJsonBodyWrap(reflect.TypeOf(changePwdRequest{}), ChangePwd),
	})

	r.Handle("/fee/list", http_util.MethodMap{
		http_util.MethodGet: http_util.SimpleWrap(GetAllRemainingFeeList),
	})
	r.Handle("/fee/claim", http_util.MethodMap{
		http_util.MethodPost: http_util.AutoSimpleJsonBodyWrap(DoCollection),
	})
	r.Handle("/fee/history/list", http_util.MethodMap{
		http_util.MethodGet: http_util.SimpleUrlQueryWrap(reflect.TypeOf(collectionHistoryListRequest{}), GetCollectionHistoryList),
	})

	r.Handle("/metric/max_balance", http_util.MethodMap{
		http_util.MethodGet: http_util.SimpleUrlQueryWrap(reflect.TypeOf(currencyMaxBalanceListRequest{}), GetCurrencyMaxBalanceList),
	})
}
