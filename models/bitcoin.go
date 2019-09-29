package models

import (
	"encoding/json"
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/configs"
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/util"
	log "github.com/sirupsen/logrus"
)

type Bitcoin struct {
	Height int64  `json:"blocks"`
	Hash   string `json:"bestblockhash"`
}

func (b *Bitcoin) GetBlockInfo() {
	body := util.GetHttpResponse("GET", configs.Config.BchApi.Bitcoin)
	b.Unmarshal(body)
}

func (b *Bitcoin) Unmarshal(body []byte) {
	err := json.Unmarshal(body, &b)
	if err != nil {
		log.Error(err)
	}
}
