package merchant

import (
	"time"
	"veric-backend/logic/db"
	"veric-backend/logic/db/types"
)

type merchantSignContentRequest struct {
	Address string `schema:"address" validate:"required"`
}

type merchantSignContentResponse struct {
	Content string `json:"content"`
}

type userRelatedMerchantRequest struct {
	Address string `json:"address"  validate:"required"`
}

type userRelatedMerchantResponse struct {
	MerchantUser []merchantUserItem `json:"merchant_users"`
}

type userNicknameRequest struct {
	Address string `json:"address"  validate:"required"`
}

type userNicknameResponse struct {
	Nickname string `json:"nickname"`
}

type merchantUserItem struct {
	MerchantId   uint   `json:"merchant_id"`
	MerchantName string `json:"merchant_name"`
	Role         string `json:"role"`
}

type checkUserExistRequest struct {
	Address string `json:"address"  validate:"required"`
}

type loginUseEthSignatureRequest struct {
	EthAddress string `json:"eth_address" validate:"required"`
	SignData   string `json:"sign_data" validate:"required"`
	MerchantId uint   `json:"merchant_id" validate:"required"`
}

type loginUseEthSignatureResponse struct {
	Token     string `json:"token"`
	ExpiredAt string `json:"expired_at"`
}

type registerUseEthSignatureRequest struct {
	EthAddress  string `json:"eth_address" validate:"required"`
	SignData    string `json:"sign_data" validate:"required"`
	StoreName   string `json:"store_name" validate:"required"`
	StoreLink   string `json:"store_link"`
	CallbackUrl string `json:"callback_url"`
	DidDocument string `json:"did_document" validate:"required"`
}

type setNicknameUseEthSignatureRequest struct {
	EthAddress string `json:"eth_address" validate:"required"`
	SignData   string `json:"sign_data" validate:"required"`
	Nickname   string `json:"nickname" validate:"required"`
}

// global
type availableCurrencyResponse struct {
	Currencies []db.Currency `json:"currencies"`
}

type exchangeRateRequest struct {
	To string `schema:"to" validate:"required"`
}

type exchangeRateResponse struct {
	From string `json:"from"`
	To   string `json:"to"`
	Rate string `json:"rate"`
}

type exchangeRateItem struct {
	From      string `json:"from"`
	To        string `json:"to"`
	Rate      string `json:"rate"`
	ExpiredAt time.Time
}

type tianApiExchangeRateResponse struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Result struct {
		Money string `json:"money"`
	} `json:"result"`
}

// merchant
type merchantBaseSummaryRequest struct {
	Tz int `schema:"tz"`
}

type merchantBaseSummaryResponse struct {
	TotalBalance      float64 `json:"total_balance"`
	TodayTotalPayment float64 `json:"today_total_payment"`
}

type merchantBalanceCurrencyItem struct {
	Currency      *db.Currency
	TotalAmount   *types.BigInt
	ClaimedAmount *types.BigInt
}

type merchantPaymentStatChartRequest struct {
	Gap string `schema:"gap" validate:"required"`
	Tz  int    `schema:"tz"`
}

type merchantPaymentStatChartResponse struct {
	Gap               string                  `json:"gap"`
	PaymentAmountStat []paymentAmountStatItem `json:"payment_amount_stat"`
}

type paymentAmountStatItem struct {
	Title  string  `json:"title"`
	Amount float64 `json:"amount"`
}

// application
type modifyApplicationBaseRequest struct {
	AppId       uint   `json:"app_id"`
	Name        string `json:"name" validate:"required"`
	Describe    string `json:"describe"`
	Link        string `json:"link"`
	CallbackUrl string `json:"callback_url"`
	LegalTender string `json:"legal_tender" validate:"required,oneof=usd cad cny"`
}

type modifyApplicationCurrencyRequest struct {
	AppId      uint   `json:"app_id"`
	Currencies []uint `json:"currencies"`
}

type modifyApplicationApiKeyRequest struct {
	AppId uint `json:"app_id"`
}

type modifyApplicationApiKeyResponse struct {
	Key       string `json:"key"`
	CreatedAt string `json:"created_at"`
}

type modifyApplicationIpnKeyRequest struct {
	AppId uint `json:"app_id"`
}

type modifyApplicationIpnKeyResponse struct {
	Key       string `json:"key"`
	CreatedAt string `json:"created_at"`
}

type modifyApplicationSlippageRequest struct {
	AppId    uint    `json:"app_id"`
	Slippage float64 `json:"slippage"`
}

// merchant user
type merchantUserListRequest struct {
	Status string `schema:"status" validate:"required"`
	Page   uint   `schema:"page" validate:"required"`
	Size   uint   `schema:"size" validate:"required"`
}

type merchantUserListResponse struct {
	Total        uint              `json:"total"`
	MerchantUser []db.MerchantUser `json:"merchant_users"`
}

type grantUserMerchantRoleRequest struct {
	TargetUser string `json:"target_user" validate:"required"`
	Role       string `json:"role" validate:"required,oneof=admin shop_manager employee"`
}

type revokeUserMerchantRoleRequest struct {
	RelationId uint `json:"relation_id" validate:"required"`
}

type changeUserMerchantRoleRequest struct {
	RelationIds []uint `json:"relation_ids" validate:"required"`
	Role        string `json:"role" validate:"required,oneof=admin shop_manager employee"`
}

type paymentSummaryRequest struct {
	Tz int `schema:"tz"`
}

type paymentSummaryResponse struct {
	Latest90DaysTotalPayment float64 `json:"latest_90_days_total_payment"`
	TodayTotalPayment        float64 `json:"today_total_payment"`
}

type paymentListRequest struct {
	Status     string `schema:"status"`
	Page       uint   `schema:"page" validate:"required"`
	Size       uint   `schema:"size" validate:"required"`
	PaymentNum string `schema:"payment_num"`
	CurrencyId uint   `schema:"currency_id"`
	Date       string `schema:"date"`
	DisplayAll uint   `schema:"display_all"`
}

type paymentListResponse struct {
	Total    uint              `json:"total"`
	Payments []paymentListItem `json:"payments"`
}

type paymentListItem struct {
	PaymentNum       string          `json:"payment_num"`
	CurrencySymbol   string          `json:"currency_symbol"`
	Amount           float64         `json:"amount"`
	Status           string          `json:"status"`
	VCs              []paymentVCItem `json:"vcs"`
	CreatedAt        time.Time       `json:"created_at"`
	FinishTime       time.Time       `json:"finish_time"`
	OrderAmount      string          `json:"order_amount"`
	CollectionAmount string          `json:"collection_amount"`
	Slippage         float64         `json:"slippage"`
}

type batchGetPaymentVcStatusRequest struct {
	VcIds []string `json:"vc_ids" validate:"required"`
}

type batchGetPaymentVcStatusResponse struct {
	Total uint            `json:"total"`
	Vcs   []paymentVCItem `json:"vcs"`
}

type paymentVCItem struct {
	VCID     string `json:"vcid"`
	VCStatus string `json:"vc_status"`
}

type paymentDetailResponse struct {
	PaymentNum        string   `json:"payment_num"`
	OriginAmount      float64  `json:"origin_amount"`
	PayAmount         float64  `json:"pay_amount"`
	PayCurrencySymbol string   `json:"pay_currency_symbol"`
	ExchangeRate      float64  `json:"exchange_rate"`
	OutcomeAmount     float64  `json:"outcome_amount"`
	PayInAddress      string   `json:"pay_in_address"`
	PayOutAddress     string   `json:"pay_out_address"`
	PaymentHash       []string `json:"payment_hash"`
	Status            string   `json:"status"`
	CreatedAt         string   `json:"created_at"`
	CompletedAt       string   `json:"completed_at"`
}

type withdrawSummaryRequest struct {
	Tz int `schema:"tz"`
}

type withdrawSummaryResponse struct {
	Latest90DaysTotalWithdraw float64 `json:"latest_90_days_total_withdraw"`
	TodayTotalWithdraw        float64 `json:"today_total_withdraw"`
}

type withdrawListRequest struct {
	Status      string `schema:"status" validate:"required"`
	Page        uint   `schema:"page" validate:"required"`
	Size        uint   `schema:"size" validate:"required"`
	WithdrawNum string `schema:"withdraw_num"`
	CurrencyId  uint   `schema:"currency_id"`
	Date        string `schema:"date"`
}

type withdrawListResponse struct {
	Total     uint               `json:"total"`
	Withdraws []withdrawListItem `json:"withdraws"`
}

type withdrawListItem struct {
	WithdrawNum    string    `json:"withdraw_num"`
	CurrencySymbol string    `json:"currency_symbol"`
	Amount         float64   `json:"amount"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
	TxHash         string    `json:"tx_hash"`
}

type newPaymentRequest struct {
	AmountInCent uint64 `json:"amount_in_cent" validate:"required"`
	ChainId      uint   `json:"chain_id" validate:"required"`
	Currency     string `json:"currency" validate:"required"`
}

type newPaymentResponse struct {
	CurrencyName         string    `json:"currency_name"`
	CurrencyDecimalCount uint      `json:"currency_decimal_count"`
	PaymentNum           string    `json:"payment_num"`
	Amount               string    `json:"amount"`
	FriendlyAmount       string    `json:"friendly_amount"`
	AmountInCent         uint64    `json:"amount_in_cent"`
	CollectionAddress    string    `json:"collection_address"`
	Status               string    `json:"status"`
	Title                string    `json:"title"`
	Desc                 string    `json:"desc"`
	FinishTime           time.Time `json:"finish_time"`
}
