package administration

import (
	"time"
	"veric-backend/logic/db"
	"veric-backend/logic/db/types"
)

type loginUsePwdRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type loginUsePwdResponse struct {
	Id         uint     `json:"id"`
	Username   string   `json:"username"`
	Nickname   string   `json:"nickname"`
	Privileges []string `json:"privileges"`
	Token      string   `json:"token"`
}

type createNewCaptchaResponse struct {
	CaptchaId      string `json:"captcha_id"`
	CaptchaContent string `json:"captcha_content"`
}

type changePwdRequest struct {
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required"`
}

type globalOverviewRequest struct {
	Tz int `schema:"tz"`
}

type globalOverviewResponse struct {
	TotalMerchantCount     uint    `json:"total_merchant_count"`
	TodayNewMerchantCount  uint    `json:"today_new_merchant_count"`
	TotalRevenueAmount     float64 `json:"total_revenue_amount"`
	TodayRevenueAmount     float64 `json:"today_revenue_amount"`
	AvailableBalanceAmount float64 `json:"available_balance_amount"`
}

type balanceAmountItem struct {
	Currency                *db.Currency
	TotalFeeRemainingAmount *types.BigInt
}

// expense related
type expenseSummaryRequest struct {
	Tz int `schema:"tz"`
}

type expenseSummaryResponse struct {
	Latest90DaysTotalExpense float64 `json:"latest_90_days_total_expense"`
	TodayTotalExpense        float64 `json:"today_total_expense"`
}

type expenseCurrencyItem struct {
	Currency  *db.Currency
	TotalUsed uint64
}

type expenseListRequest struct {
	Page   uint   `schema:"page" validate:"required"`
	Size   uint   `schema:"size" validate:"required"`
	Status string `schema:"status"`
	Type   string `schema:"type"`
	Date   string `schema:"date"`
}

type expenseListResponse struct {
	Total    uint               `json:"total"`
	Expenses []*expenseListItem `json:"expenses"`
}

type expenseListItem struct {
	ChainName string    `json:"chain_name"`
	Type      string    `json:"type"`
	Amount    float64   `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	TxHash    string    `json:"tx_hash"`
}

// merchant related
type merchantListRequest struct {
	Status string `schema:"status" validate:"required"`
	Page   uint   `schema:"page" validate:"required"`
	Size   uint   `schema:"size" validate:"required"`
	Did    string `schema:"did"`
}

type changeMerchantStatusRequest struct {
	MerchantId uint   `json:"merchant_id" validate:"required"`
	NewStatus  string `json:"new_status" validate:"required"`
}

type suspendAllMerchantRequest struct {
	NewStatus string `json:"new_status" validate:"required"`
}

type merchantListResponse struct {
	Total      uint           `json:"total"`
	Merchants  []*db.Merchant `json:"merchants"`
	SuspendAll bool           `json:"suspend_all"`
}

// merchant application related
type applicationListRequest struct {
	Status     string `schema:"status" validate:"required"`
	Page       uint   `schema:"page" validate:"required"`
	Size       uint   `schema:"size" validate:"required"`
	MerchantId uint   `schema:"merchant_id" validate:"required"`
}

type changeApplicationStatusRequest struct {
	ApplicationId uint   `json:"app_id" validate:"required"`
	NewStatus     string `json:"new_status" validate:"required"`
}

type applicationListResponse struct {
	Total        uint             `json:"total"`
	Applications []db.Application `json:"applications"`
}

// pool related
type poolListRequest struct {
	Status string `schema:"status" validate:"required"`
	Page   uint   `schema:"page" validate:"required"`
	Size   uint   `schema:"size" validate:"required"`
}

type changePoolStatusRequest struct {
	PoolId    uint   `json:"pool_id" validate:"required"`
	NewStatus string `json:"new_status" validate:"required"`
}

type poolListResponse struct {
	Total uint      `json:"total"`
	Pools []db.Pool `json:"pools"`
}

// blockchain related
type blockchainListRequest struct {
	Status string `schema:"status" validate:"required"`
	Page   uint   `schema:"page" validate:"required"`
	Size   uint   `schema:"size" validate:"required"`
}

type changeBlockchainStatusRequest struct {
	ChainId   uint   `json:"chain_id" validate:"required"`
	NewStatus string `json:"new_status" validate:"required"`
}

type blockchainListResponse struct {
	Total       uint            `json:"total"`
	Blockchains []db.Blockchain `json:"blockchains"`
}

// currency related
type currencyListRequest struct {
	Status string `schema:"status" validate:"required"`
	Page   uint   `schema:"page" validate:"required"`
	Size   uint   `schema:"size" validate:"required"`
}

type changeCurrencyStatusRequest struct {
	CurrencyId uint   `json:"currency_id" validate:"required"`
	NewStatus  string `json:"new_status" validate:"required"`
}

type currencyListResponse struct {
	Total      uint          `json:"total"`
	Currencies []db.Currency `json:"currencies"`
}

// manager account related
type managerListRequest struct {
	Status string `schema:"status" validate:"required"`
	Page   uint   `schema:"page" validate:"required"`
	Size   uint   `schema:"size" validate:"required"`
}

type managerListResponse struct {
	Total    uint         `json:"total"`
	Managers []db.Manager `json:"managers"`
}

type createManagerRequest struct {
	Name       string `json:"name" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required"`
	Privileges []uint `json:"privileges" validate:"required"`
}

type changeManagerRequest struct {
	ManagerId  uint   `json:"manager_id" validate:"required"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Privileges []uint `json:"privileges"`
}

type deleteManagerRequest struct {
	ManagerId uint `json:"manager_id" validate:"required"`
}

type permissionListResponse struct {
	Total       uint            `json:"total"`
	Permissions []db.Permission `json:"permissions"`
}

// fee withdraw
type doCollectionRequest struct {
	FeeType    string `json:"fee_type"`
	CurrencyId uint   `json:"currency_id"`
	Remark     string `json:"remark"`
}

type allRemainingFeeListResponse struct {
	Total uint                `json:"total"`
	Fees  []*remainingFeeItem `json:"fees"`
}

type remainingFeeItem struct {
	ChainName                string
	ChainImageUrl            string
	CurrencyId               uint
	CurrencyName             string
	CurrencySymbol           string
	CurrencyImageUrl         string
	AvailablePlatformBalance float64
	AvailablePoolBalance     float64
}

type remainingFeeAmountItem struct {
	TotalFeeRemainingAmount    *types.BigInt
	PlatformFeeRemainingAmount *types.BigInt
	PoolFeeRemainingAmount     *types.BigInt
}

type collectionHistoryListRequest struct {
	Page       uint   `schema:"page" validate:"required"`
	Size       uint   `schema:"size" validate:"required"`
	CurrencyId uint   `schema:"currency_id"`
	Date       string `schema:"date"`
}

type collectionHistoryListResponse struct {
	Total       uint                        `json:"total"`
	Collections []collectionHistoryListItem `json:"collections"`
}

type collectionHistoryListItem struct {
	ChainName      string    `json:"chain_name"`
	CurrencySymbol string    `json:"currency_symbol"`
	Amount         float64   `json:"amount"`
	CreatedAt      time.Time `json:"created_at"`
	TxHash         string    `json:"tx_hash"`
}

// metric related
type currencyMaxBalanceListRequest struct {
	Secret string `schema:"secret"`
}

type currencyMaxBalanceListResponse struct {
	Total      uint                          `json:"total"`
	Currencies []*currencyMaxBalanceListItem `json:"currencies"`
}

type currencyMaxBalanceListItem struct {
	Currency   *db.Currency
	MaxBalance string
}
