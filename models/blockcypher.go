package models

import (
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/configs"
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/util"
	"encoding/json"
	log "github.com/sirupsen/logrus"
)
type BlockCypher struct {
	Height int64 `json:"height"`
	Hash string `json:"hash"`
}

func (b *BlockCypher)GetBlockInfo(){
	body:=util.GetHttpResponse("GET",configs.Config.LtcApi.BlockCypher)
	b.Unmarshal(body)
}

func(b *BlockCypher)Unmarshal(body []byte){
	err:=json.Unmarshal(body,&b)
	if err!=nil{
		log.Error(err)
	}
}
