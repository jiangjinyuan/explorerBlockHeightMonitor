package models

import (
	"encoding/json"
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/configs"
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/util"
	log "github.com/sirupsen/logrus"
)

type Alethio struct {
	Height int64 `json:"number"`
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
	//fmt.Println(data)
	temp := data["data"].(map[string]interface{})
	dat := temp["attributes"]
	out, err2 := json.Marshal(dat)
	if err2 != nil {
		log.Error(err2)
	}
	err3 := json.Unmarshal(out, &b)
	if err3 != nil {
		log.Error(err3)
	}
}
