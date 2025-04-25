package conf

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type (
	ServerName string
	Server     struct {
		GrpcAddress string `json:"grpc_address"`
		HttpAddress string `json:"http_address"`
	}
	LogConfig struct {
		LogLevel   string `required:"true"`
		LogPath    string `required:"true"`
		MaxAgeDays int    `required:"true"`
		MaxSize    int    `required:"true"`
		MaxBackups int    `required:"true"`
		Compress   bool
	}
	TencentSms struct {
		SecretID      string
		SecretKey     string
		AppID         string
		TemplateID    string
		Endpoint      string
		Region        string
		SignName      string
		ExpireMinutes int
	}

	Redis struct {
		Addr     string
		Db       int
		AuthInfo string
	}

	Metrics struct {
		Addr string
		Path string
	}

	VisionAiServerConfig struct {
		Port        int
		HTTPPort    int
		DialTimeout time.Duration `required:"true"`
		Mode        string
		// 服务部署地区
		Region string
	}
	Mysql struct {
		Addr          string `required:"true"`
		User          string
		Password      string
		Db            string
		Port          int
		MigrationsDir string
	}
	MongoDB struct {
		Addr       string `required:"true"`
		User       string
		Password   string
		Db         string
		Collection string
	}
	AmountConfig struct {
		SigninDayOfAmount []DayOfAmount
		SignMaxDays       int
		AdRewardMaxPerDay int64
		AdRewardPlay      int64
		AdRewardClick     int64
		AdRewardAction    int64
	}
	DayOfAmount struct {
		Day    int
		Amount int64
	}
	LlmConfig struct {
		GPT           GPTModel
		Doubao        DoubaoModel
		DoubaoVision  DoubaoVisionModel
		VolcEngineASR VolcEngineASR
		VolcEngineTTS VolcEngineTTS
		Kimi          KimiConfig
		DeepSeek      DeepSeek
		AzureOpenAI   AzureOpenAI
		Tencent       Tencent
		Gemini        Gemini
		TitlePrompt   string
	}
	GPTModel struct {
		APIKey   string
		Endpoint string
	}
	OssConfig struct {
		Endpoint        string
		AccessKeyID     string
		AccessKeySecret string
		Bucket          string
		UgcAddr         string
	}
	CosConfig struct {
		Endpoint        string
		AccessKeyID     string
		AccessKeySecret string
		Bucket          string
		UgcAddr         string
	}
	DoubaoModel struct {
		APIKey          string
		Endpoint        string
		Region          string
		Doubao          string
		DoubaoPro       string
		DeepSeekR1      string
		DoubaoProVision string
	}
	DoubaoVisionModel struct {
		APIKey       string
		Endpoint     string
		Region       string
		DoubaoVision string
	}
	VolcEngineASR struct {
		AppKey    string
		AccessKey string
		Endpoint  string
	}
	VolcEngineTTS struct {
		AppKey    string
		AccessKey string
		Endpoint  string
	}

	KimiConfig struct {
		Endpoint  string `yaml:"endpoint"`
		APIKey    string `yaml:"api_key"`
		ModelName string `yaml:"model_name"`
	}

	Tencent struct {
		Endpoint string `yaml:"endpoint"`
		APIKey   string `yaml:"api_key"`
	}

	DeepSeek struct {
		Endpoint   string `yaml:"endpoint"`
		APIKey     string `yaml:"api_key"`
		DeepSeekR1 string `yaml:"deepseek_r1"`
		DeepSeekV3 string `yaml:"deepseek_v3"`
	}
	AzureOpenAI struct {
		Endpoint string `yaml:"endpoint"`
		APIKey   string `yaml:"api_key"`
		GPT4O    string `yaml:"gpt_4o"`
	}

	Gemini struct {
		Endpoint string `yaml:"endpoint"`
		APIKey   string `yaml:"api_key"`
		AuthFile string `yaml:"auth_file"`
	}

	ComfyUI struct {
		Server   string
		Port     int
		Username string
		Password string
	}

	PayLinker struct {
		Addr string
	}

	AstroServer struct {
		Addr string
	}
)

type Config struct {
	ServerName           ServerName
	Server               Server
	LogConfig            LogConfig
	Metrics              Metrics
	VisionAiServerConfig VisionAiServerConfig
	AmountConfig         AmountConfig
	Mysql                Mysql
	Redis                Redis
	MongoDB              MongoDB
	TencentSms           TencentSms
	LlmConfig            LlmConfig
	OssConfig            OssConfig
	CosConfig            CosConfig
	ComfyUI              ComfyUI
	PayLinker            PayLinker
	AstroServer          AstroServer
}

var GlobalConfig Config

func ConfigInit(configPath string) (err error) {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(configPath)
	if err = viper.ReadInConfig(); err != nil {
		return err
	}
	err = viper.Unmarshal(&GlobalConfig)
	if err != nil {
		return err
	}
	return nil
}

func (c LogConfig) String() string {
	return fmt.Sprintf(`LogConfig: LogPath:%s, LogLevel:%s`, c.LogPath, c.LogLevel)
}

func IsDev() bool {
	return GlobalConfig.VisionAiServerConfig.Mode == "DEV"
}
func IsProd() bool {
	return !IsDev() && !IsLocal()
}
func IsLocal() bool {
	return GlobalConfig.VisionAiServerConfig.Mode == "LOCAL"
}
