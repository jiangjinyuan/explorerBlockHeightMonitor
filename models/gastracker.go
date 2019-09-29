package models

import (
	"encoding/json"
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/configs"
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/util"
	log "github.com/sirupsen/logrus"
)

type Gastracker struct {
	Height int64  `json:"height"`
	Hash   string `json:"hash"`
}

func (b *Gastracker) GetBlockInfo() {
	body := util.GetHttpResponse("GET", configs.Config.EtcApi.Gastracker)
	b.Unmarshal(body)
}

func (b *Gastracker) Unmarshal(body []byte) {
	var data map[string]interface{}
	err1 := json.Unmarshal(body, &data)
	if err1 != nil {
		log.Error(err1)
	}
	temp := data["items"].([]interface{})
	out, err2 := json.Marshal(temp[0])
	if err2 != nil {
		log.Error(err2)
	}
	err3 := json.Unmarshal(out, &b)
	if err3 != nil {
		log.Error(err3)
	}
}
