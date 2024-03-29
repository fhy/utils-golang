package config

import "go.mongodb.org/mongo-driver/mongo/options"

type Server struct {
	IP           string `yaml:"ip"`
	Port         int    `yaml:"port"`
	ReadTimeout  int    `yaml:"readtimeout"`
	WriteTimeout int    `yaml:"writetimeout"`
}

type CookieConfig struct {
	Domain string `yaml:"domain"`
	AgeMax int    `yaml:"age"`
}

type LogConfig struct {
	Path          string `yaml:"path"`
	RotationCount int    `yaml:"rotation_count"`
	Level         string `yaml:"level"`
}

type Aliyun struct {
	AccessKeyId     string `yaml:"id"`
	AccessKeySecret string `yaml:"secret"`
	RegionId        string `yaml:"region"`
}

type Oss struct {
	AccessKeyId     string `yaml:"id"`
	AccessKeySecret string `yaml:"secret"`
	Endpoint        string `yaml:"endpoint"`
	Bucket          string `yaml:"bucket"`
	Prefix          string `yaml:"prefix"`
}

type DingTalkConfig struct {
	AppKey    string `yaml:"id"`
	AppSecret string `yaml:"secret"`
}

// database
// type DbConfig struct {
// 	Type   string      `yaml:"type"`
// 	Config interface{} `yaml:"config"`
// }

// jwt
type Jwt struct {
	Issuer      string `yaml:"issuer"`
	PubKeyPath  string `yaml:"pubkey"`
	PrivKeyPath string `yaml:"privkey"`
	Secret      string `yaml:"secret"`
}

type SqliteConfig struct {
	DbFile string `yaml:"dbfile"`
}

type MongoDBConfig struct {
	Host     string           `yaml:"host"`
	Port     int              `yaml:"port"`
	Username string           `yaml:"username"`
	Password string           `yaml:"password"`
	DB       string           `yaml:"db"`
	Timeout  int              `yaml:"timeout"`
	LogLevel options.LogLevel `yaml:"logLevel"`
}

type RedisConfig struct {
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	Password    string `yaml:"password"`
	MaxActive   int    `yaml:"maxActive"`
	MaxIdle     int    `yaml:"maxIdle"`
	IdleTimeout int    `yaml:"idleTimeout"`
	DB          int    `yaml:"db"`
}

type OfficialAccountConfig struct {
	AppID          string `yaml:"appID"`
	AppSecret      string `yaml:"appSecret"`
	Token          string `yaml:"token"`
	EncodingAESKey string `yaml:"encodingAESKey"`
}

type MiniProgramConfig struct {
	AppID     string `yaml:"appID"`
	AppSecret string `yaml:"appSecret"`
}

type WechatPayConfig struct {
	AppID     string `yaml:"app_id"`
	MchID     string `yaml:"mch_id"`
	Key       string `yaml:"key"`
	NotifyURL string `yaml:"notify_url"`
}

type RocketMQConfig struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	AccessKey string `yaml:"access"`
	SecretKey string `yaml:"secret"`
}

/*
func LoadConf(configFile string, config interface{}) error {
	fmt.Printf("loading configfile: %s\n", configFile)
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		fmt.Print("can't find config path\n")
		os.Exit(1)
	} else {
		if err != nil {
			fmt.Print("Decode Config Error", err)
			os.Exit(1)
		}
	}
	return LoadConfYaml(configFile, config)
}
*/
