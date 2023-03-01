package config

import (
	"encoding/json"
	"github.com/spf13/viper"
	"log"
	"regexp"
	"strings"
)

var config Config = Config{
	Debug: &DebugConfig{
		Enable:  false,
		Verbose: false,
	},
	DB: &DBConfig{
		Host:     "127.0.0.1",
		Port:     33069,
		User:     "veric_backend",
		Password: "password",
		DbName:   "veric_backend",
		TimeZone: "Local",
	},
	HTTP: &HTTPConfig{
		Listen:           "0.0.0.0",
		Port:             8888,
		JwtEncryptSecret: "WQQ0@R625M2376x#8j$4l",
		MetricSecret:     "RmYieFS237623HXkJO7zFzGKdq2TVqyEU",
	},
	Rpc: &RpcConfig{
		RPCHost: "https://mainnet.infura.io/v3/",
		ChainId: 1666700000,
	},
	Logic: &LogicConf{
		SuspendAllMerchant:      false,
		WalletPoolLackThreshold: 0.8,
	},
	HCaptcha: &HCaptchaConfig{
		SiteKey:   "",
		SecretKey: "",
	},
	TianApi: &TianApiConfig{
		Key: "",
	},
	PagerDuty: &PagerDutyConfig{
		Key: "",
	},
}

func init() {
	log.Println("init config...")

	instance := viper.New()

	// only for dev
	instance.AddConfigPath("/etc/veric-backend/")
	instance.AddConfigPath(".")

	instance.SetConfigType("yaml")
	instance.SetConfigName("config.veric-backend.yaml")

	instance.SetEnvPrefix("vb")
	instance.AutomaticEnv()
	instance.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := instance.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			panic(err)
		}
	}

	err = instance.Unmarshal(&config)
	if err != nil {
		panic(err)
	}

	priKeyRegex := regexp.MustCompile("[0-9a-hA-H]{64}")
	configJson, err := json.Marshal(config)
	if err != nil {
		return
	}

	log.Printf("dump config: %s", priKeyRegex.ReplaceAll(configJson, []byte("[PRIVATE KEY]")))
}

func Get() *Config {
	return &config
}

func SetSuspendAllMerchant(val bool) {
	config.Logic.SuspendAllMerchant = val
}
