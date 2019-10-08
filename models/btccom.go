package models

import (
	"encoding/json"
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/configs"
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/util"
	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
)

type BTCcom struct {
	Height int64  `mapstructure:"height"`
	Hash   string `mapstructure:"hash"`
}

func (b *BTCcom) GetBlockInfo(coin string) {
	var body []byte
	switch coin {
	case "btc":
		body = util.GetHttpResponse("GET", configs.Config.BtcApi.BTCcom)
	case "bch":
		body = util.GetHttpResponse("GET", configs.Config.BchApi.BTCcom)
	case "ltc":
		body = util.GetHttpResponse("GET", configs.Config.LtcApi.BTCcom)
	case "eth":
		body = util.GetHttpResponse("GET", configs.Config.EthApi.BTCcom)
	case "etc":
		body = util.GetHttpResponse("GET", configs.Config.EtcApi.BTCcom)
	}
	b.Unmarshal(body)

}

func (b *BTCcom) Unmarshal(body []byte) {
	var data map[string]interface{}
	err1 := json.Unmarshal(body, &data)
	if err1 != nil {
		log.Error(err1)
	}
	temp := data["data"].(map[string]interface{})
	err2 := mapstructure.Decode(temp, &b)
	if err2 != nil {
		panic(err2)
	}
}
