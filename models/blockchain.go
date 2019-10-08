package models

import (
	"encoding/json"
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/configs"
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/util"
	log "github.com/sirupsen/logrus"
)

type BlockChainBtc struct {
	Height int64  `json:"height"`
	Hash   string `json:"hash"`
}
type BlockChainEth struct {
	Height string `json:"number"`
	Hash   string `json:"hash"`
}

func (b *BlockChainBtc) GetBlockInfo() {
	body := util.GetHttpResponse("GET", configs.Config.BtcApi.BlockChain)
	b.Unmarshal(body)
}

func (b *BlockChainBtc) Unmarshal(body []byte) {
	err := json.Unmarshal(body, &b)
	if err != nil {
		log.Error(err)
	}
}

func (a *BlockChainEth) GetBlockInfo() {
	body := util.GetHttpResponse("GET", configs.Config.EthApi.BlockChain)
	a.Unmarshal(body)
}

func (a *BlockChainEth) Unmarshal(body []byte) {
	var data map[string]interface{}
	err1 := json.Unmarshal(body, &data)
	if err1 != nil {
		log.Error(err1)
		//fmt.Println(err1)
	}
	temp := data["blockHeaders"].([]interface{})
	out, err2 := json.Marshal(temp[0])
	if err2 != nil {
		log.Error(err2)
	}
	err3 := json.Unmarshal(out, &a)
	if err3 != nil {
		log.Error(err3)
	}
}
