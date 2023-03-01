package administration

import (
	"go.uber.org/zap"
	"net/http"
	"strings"
	"time"
	"veric-backend/internal/log"
	"veric-backend/internal/util"
	"veric-backend/logic/blockchain/exchange"
	"veric-backend/logic/db"
	"veric-backend/logic/db/types"
	"veric-backend/logic/http/http_util"
)

func GetExpenseSummary(typ interface{}, r *http_util.HTTPContext) (resp interface{}, respErr error) {
	tz := r.QueryWithDefaultInt("tz", 0)

	// get admin
	jwtClaims, jwtErr := r.GetAdminJwt("admin-auth-token")
	if jwtErr != nil {
		return nil, http_util.NewHttpError(http.StatusForbidden, "need permission for this operate")
	}

	if !r.CheckPermissionAny(jwtClaims.Privileges, []string{"expense"}) {
		return nil, http_util.NewHttpError(http.StatusForbidden, "no permission for this operation")
	}

	// get native currency
	currencies, getNativeErr := db.FindAllNativeCurrency(db.WithPreload("Blockchain"))
	if getNativeErr != nil {
		return nil, getNativeErr
	}
	chainNativeCurrencySet := make(map[string]*db.Currency, 0)
	for _, currency := range currencies {
		chainNativeCurrencySet[currency.Blockchain.Name] = currency
	}

	nowTime := time.Now().Add(time.Duration(tz) * time.Hour)
	dateBegin := nowTime.AddDate(0, 0, -90).Format(util.DateLayer)
	dateToday := nowTime.Format(util.DateLayer)

	expenseSummary := &expenseSummaryResponse{
		Latest90DaysTotalExpense: float64(0),
		TodayTotalExpense:        float64(0),
	}

	// get latest 90 days tx expense
	latest90Txs, sum90Err := db.FindTxByCreateTime(util.GetTimeStrFixedByTZ(dateBegin+" 00:00:00", tz), util.GetTimeStrFixedByTZ(dateToday+" 23:59:59", tz))
	if sum90Err != nil {
		return nil, sum90Err
	}

	latest90ChainExpenseAmountSet := make(map[string]*types.BigInt, 0)
	for _, txItem := range latest90Txs {
		if txItem.GasUsed.IsZero() {
			continue
		}

		chainName := strings.Replace(txItem.ChainName, "Dynamic", "", -1)

		if _, ok := latest90ChainExpenseAmountSet[chainName]; !ok {
			latest90ChainExpenseAmountSet[chainName] = types.NewBigIntZero()
		}

		latest90ChainExpenseAmountSet[chainName].Add(txItem.GasUsed)
	}

	// amount to usd
	for chain, amount := range latest90ChainExpenseAmountSet {
		if currency, ok := chainNativeCurrencySet[chain]; ok {
			expenseAmountRaw, exchangeErr := exchange.DefaultManage.ExchangeCoinToUSD(amount.RawBigInt(), currency)
			if exchangeErr != nil {
				log.GetLogger().Warn("exchange currency to USD failed. skipped", zap.String("currency", currency.Name))
				continue
			}

			expenseAmount, _ := expenseAmountRaw.Float64()
			expenseSummary.Latest90DaysTotalExpense += expenseAmount
		} else {
			log.GetLogger().Warn("can not get chain native currency. skipped", zap.String("chain", chain))
		}
	}

	// get today summary
	todayTxs, sumErr := db.FindTxByCreateTime(util.GetTimeStrFixedByTZ(dateToday+" 00:00:00", tz), util.GetTimeStrFixedByTZ(dateToday+" 23:59:59", tz))
	if sumErr != nil {
		return nil, sumErr
	}

	todayChainExpenseAmountSet := make(map[string]*types.BigInt, 0)
	for _, txItem := range todayTxs {
		if txItem.GasUsed.IsZero() {
			continue
		}

		chainName := strings.Replace(txItem.ChainName, "Dynamic", "", -1)

		if _, ok := todayChainExpenseAmountSet[chainName]; !ok {
			todayChainExpenseAmountSet[chainName] = types.NewBigIntZero()
		}

		todayChainExpenseAmountSet[chainName].Add(txItem.GasUsed)
	}

	// amount to usd
	for chain, amount := range todayChainExpenseAmountSet {
		if currency, ok := chainNativeCurrencySet[chain]; ok {
			expenseAmountRaw, exchangeErr := exchange.DefaultManage.ExchangeCoinToUSD(amount.RawBigInt(), currency)
			if exchangeErr != nil {
				log.GetLogger().Warn("exchange currency to USD failed. skipped", zap.String("currency", currency.Name))
				continue
			}

			expenseAmount, _ := expenseAmountRaw.Float64()
			expenseSummary.TodayTotalExpense += expenseAmount
		} else {
			log.GetLogger().Warn("can not get chain native currency. skipped", zap.String("chain", chain))
		}
	}

	return expenseSummary, nil
}

func GetExpenseList(typ interface{}, r *http_util.HTTPContext) (resp interface{}, respErr error) {
	page := r.QueryWithDefaultInt("page", 1)
	size := r.QueryWithDefaultInt("size", 10)
	status := r.QueryWithDefault("status", "all")
	method := r.QueryWithDefault("type", "all")
	date := r.QueryWithDefault("date", "")

	// get admin
	jwtClaims, jwtErr := r.GetAdminJwt("admin-auth-token")
	if jwtErr != nil {
		return nil, http_util.NewHttpError(http.StatusForbidden, "need permission for this operate")
	}

	if !r.CheckPermissionAny(jwtClaims.Privileges, []string{"expense"}) {
		return nil, http_util.NewHttpError(http.StatusForbidden, "no permission for this operation")
	}

	// get native currency
	currencies, getNativeErr := db.FindAllNativeCurrency(db.WithPreload("Blockchain"))
	if getNativeErr != nil {
		return nil, getNativeErr
	}
	chainNativeCurrencySet := make(map[string]*db.Currency, 0)
	for _, currency := range currencies {
		chainNativeCurrencySet[currency.Blockchain.Name] = currency
	}

	// calc pagination
	offset := (page - 1) * size

	// get transaction list
	var (
		expenses = make([]*expenseListItem, 0)
		total    uint
	)

	total = db.GetTxCountByStatusMethodDate(status, method, date)
	txSet, findErr := db.FindPaginationTxByStatusMethodDate(status, method, date, offset, size)
	if findErr != nil {
		return nil, findErr
	}

	for _, txItem := range txSet {
		chainName := strings.Replace(txItem.ChainName, "Dynamic", "", -1)
		if currency, ok := chainNativeCurrencySet[chainName]; ok {
			amount, _ := exchange.DefaultManage.ExchangeCoinToNaturalAmount(txItem.GasUsed.RawBigInt(), currency).Float64()

			expenses = append(expenses, &expenseListItem{
				ChainName: chainName,
				Type:      txItem.Method,
				Amount:    amount,
				Status:    txItem.Status,
				CreatedAt: txItem.CreatedAt,
				TxHash:    txItem.Hash,
			})
		} else {
			log.GetLogger().Warn("can not get chain native currency. skipped", zap.String("chain", txItem.ChainName))
		}

	}

	return &expenseListResponse{
		Total:    total,
		Expenses: expenses,
	}, nil
}
