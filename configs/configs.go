package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

var Config = struct {
	AppName           string
	Slack             Slack
	Email             Email
	AlarmThreshold    map[string]int64
	TraverseSleepTime map[string]int64
	ExplorerDatabase  map[string]MySQLDB
	Redis             struct {
		Prefix  string
		Type    string
		Redis   RedisDB
		Cluster RedisClusterDB
	}
	SupportCoins     string
	MonitorExplorers []string
	MonitorCoins     []string
	Interval         int
	Health           Health
}{}

type Slack struct {
	WebHookURL string
	IsEnable   bool
}

type Email struct {
	SenderName     string
	SenderPassword string
	Host           string
	Port           int
	IsEnable       bool
}

type Health struct {
	Port            int
	IntervalSeconds int
}

func InitConfig(files string) {
	viper.SetConfigName(files) // 设置配置文件名 (不带后缀)
	viper.AddConfigPath(".")   // 第一个搜索路径
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&Config); err != nil {
		panic(err)
	}

	fmt.Println(&Config)
}
