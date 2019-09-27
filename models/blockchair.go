package models

import (
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/configs"
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/util"
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

type BlockChair struct {
	Height int64  `json:"id"`
	Hash   string `json:"hash"`
}

func (b *BlockChair) GetBlockInfo(coin string) {
	var body []byte
	switch coin {
	case "btc":
		body = util.GetHttpResponse("GET",configs.Config.BtcApi.BlockChair)
	case "bch":
		body = util.GetHttpResponse("GET",configs.Config.BchApi.BlockChair)
	case "ltc":
		body = util.GetHttpResponse("GET",configs.Config.LtcApi.BlockChair)
	case "eth":
		body = util.GetHttpResponse("GET",configs.Config.EthApi.BlockChair)
		//body=util.HttpHandle("GET",configs.Config.EthApi.BlockChair)
	}
	b.Unmarshal(body)
}

func (b *BlockChair) Unmarshal(body []byte) {
	var data map[string]interface{}
	err1 := json.Unmarshal(body, &data)
	if err1 != nil {
		log.Error(err1)
	}
	temp := data["data"].([]interface{})
	out, err2 := json.Marshal(temp[0])
	if err2 != nil {
		log.Error(err2)
	}
	err3 := json.Unmarshal(out, &b)
	if err3 != nil {
		log.Error(err3)
	}
}
