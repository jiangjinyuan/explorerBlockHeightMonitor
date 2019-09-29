package models

import (
	"encoding/json"
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/configs"
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/util"
	log "github.com/sirupsen/logrus"
)

type EtcBlockExplorer struct {
	Height int64  `json:"blockNumber"`
	Hash   string `json:"blockHash"`
}

func (b *EtcBlockExplorer) GetBlockInfo() {
	tmp := `{"action":"latest_txs"}`
	body := util.PostHttpResponse(configs.Config.EtcApi.EtcBlockExplorer, tmp)
	//fmt.Println(body)
	b.Unmarshal(body)
}

func (b *EtcBlockExplorer) Unmarshal(body []byte) {
	var data map[string]interface{}
	err1 := json.Unmarshal(body, &data)
	if err1 != nil {
		log.Error(err1)
	}
	temp := data["txs"].([]interface{})
	out, err2 := json.Marshal(temp[0])
	if err2 != nil {
		log.Error(err2)
	}
	err3 := json.Unmarshal(out, &b)
	if err3 != nil {
		log.Error(err3)
	}
}
