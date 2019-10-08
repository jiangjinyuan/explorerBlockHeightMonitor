package models

import (
	"encoding/json"
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/configs"
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/util"
	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
)

type Etherscan struct {
	Height string `mapstructure:"number"`
	Hash   string `mapstructure:"hash"`
}

func (b *Etherscan) GetBlockInfo() {
	body := util.GetHttpResponse("GET", configs.Config.EthApi.Etherscan)
	b.Unmarshal(body)
}

func (b *Etherscan) Unmarshal(body []byte) {
	var data map[string]interface{}
	err1 := json.Unmarshal(body, &data)
	if err1 != nil {
		log.Error(err1)
	}
	temp := data["result"].(map[string]interface{})
	err2 := mapstructure.Decode(temp, &b)
	if err2 != nil {
		panic(err2)
	}
}
