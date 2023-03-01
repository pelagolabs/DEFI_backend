package config

type Config struct {
	Debug         *DebugConfig
	DB            *DBConfig
	HTTP          *HTTPConfig
	Rpc           *RpcConfig
	AccountPriKey *AccountPriKeyConfig
	Contract      *ContractConf
	Logic         *LogicConf
	HCaptcha      *HCaptchaConfig
	TianApi       *TianApiConfig
	PagerDuty     *PagerDutyConfig
}

type DebugConfig struct {
	Enable  bool
	Verbose bool
}

type DBConfig struct {
	Host     string
	User     string
	Password string
	DbName   string
	Port     uint16
	TimeZone string
}

type HTTPConfig struct {
	Listen           string
	Port             uint16
	JwtEncryptSecret string
	MetricSecret     string
}

type RpcConfig struct {
	RPCHost       string
	ChainId       int64
	AccountPriKey string
}

type AccountPriKeyConfig struct {
	DIDIssuer string
	Contract  string
}

type ContractConf struct {
	HarmonyEndpoint  string
	EthereumEndpoint string
	AddressPoolAddr  string
	MultiCallAddr    string
	DidRegistryAddr  string
}

type LogicConf struct {
	AdministrationCollectionAddr string
	SuspendAllMerchant           bool
	WalletPoolLackThreshold      float64
}

type HCaptchaConfig struct {
	SiteKey   string
	SecretKey string
}

type TianApiConfig struct {
	Key string
}

type PagerDutyConfig struct {
	Key string
}
