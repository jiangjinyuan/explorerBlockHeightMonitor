package configs

import (
	"github.com/spf13/viper"
)

var Config = struct {
	Slack             Slack
	Email             Email
	AlarmThreshold    AlarmThreshold
	TraverseSleepTime TraverseSleepTime
	ExplorerDatabase  map[string]MySQLDB
	Redis             struct {
		Prefix  string
		Type    string
		Redis   RedisDB
		Cluster RedisClusterDB
	}
	SupportCoins string
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

type DB struct {
	Driver   string
	UserName string
	Password string
	Host     string
	Port     string
	Schema   string
}

type AlarmThreshold struct {
	Btc int64
	Bch int64
	Ltc int64
	Eth int64
	Etc int64
}

type TraverseSleepTime struct {
	Btc int
	Bch int
	Ltc int
	Eth int
	Etc int
}

func InitConfig(files string) {
	viper.SetConfigName(files)  // 设置配置文件名 (不带后缀)
	viper.AddConfigPath(".")    // 第一个搜索路径
	err := viper.ReadInConfig() // 读取配置数据
	if err != nil {
		panic(err)
	}
	viper.Unmarshal(&Config) // 将配置信息绑定到结构体上
	//fmt.Println(Config)
}
