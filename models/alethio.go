package models

import (
	"encoding/json"
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/configs"
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/util"
	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
)

type Alethio struct {
	Height int64 `mapstructure:"height"`
}

func (b *Alethio) GetBlockInfo() {
	body := util.Get(configs.Config.EthApi.BlockChair)
	b.Unmarshal(body)
}

func (b *Alethio) Unmarshal(body []byte) {
	var data map[string]interface{}
	err1 := json.Unmarshal(body, &data)
	if err1 != nil {
		log.Error(err1)
	}
	temp := data["data"].(map[string]interface{})
	dat := temp["attributes"]
	err2 := mapstructure.Decode(dat, &b)
	if err2 != nil {
		panic(err2)
	}
}
