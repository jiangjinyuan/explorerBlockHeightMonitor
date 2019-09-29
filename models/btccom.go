package models

import (
	"encoding/json"
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/configs"
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/util"
	log "github.com/sirupsen/logrus"
)

type BTCcom struct {
	Height int64  `json:"height"`
	Hash   string `json:"hash"`
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
	temp := data["block"].(map[string]interface{})
	out, err2 := json.Marshal(temp)
	if err2 != nil {
		log.Error(err2)
	}
	err3 := json.Unmarshal(out, &b)
	if err3 != nil {
		log.Error(err3)
	}
}
