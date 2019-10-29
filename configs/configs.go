package configs

import (
	"github.com/spf13/viper"
)

var Config = struct {
	BtcApi         BTC
	BchApi         BCH
	LtcApi         LTC
	EthApi         ETH
	EtcApi         ETC
	BtcNode        BtcNode
	BchNode        BchNode
	LtcNode        LtcNode
	EthNode        EthNode
	EtcNode        EtcNode
	Slack          Slack
	Email          Email
	DB             DB
	AlarmThreshold AlarmThreshold
	TraverseSleepTime TraverseSleepTime
}{}

type BTC struct {
	BlockChain string
	BlockChair string
	ViaBtc     string
	BTCcom     string
}

type BCH struct {
	BlockChair string
	Bitcoin    string
	ViaBtc     string
	BTCcom     string
}

type LTC struct {
	BlockChair  string
	ViaBtc      string
	BlockCypher string
	BTCcom      string
}

type ETH struct {
	Etherscan  string
	BlockChain string
	BlockChair string
	BTCcom     string
}

type ETC struct {
	Gastracker       string
	BlockScout       string
	EtcBlockExplorer string
	BTCcom           string
}

type BtcNode struct {
	ServerEndPoint string
	UserName       string
	Password       string
}

type BchNode struct {
	ServerEndPoint string
	UserName       string
	Password       string
}

type LtcNode struct {
	ServerEndPoint string
	UserName       string
	Password       string
}

type EthNode struct {
	ServerEndPoint string
	UserName       string
	Password       string
}

type EtcNode struct {
	ServerEndPoint string
	UserName       string
	Password       string
}
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
