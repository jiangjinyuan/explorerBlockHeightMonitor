package models

import (
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/configs"
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/util"
	"encoding/json"
	log "github.com/sirupsen/logrus"
)
type ViaBtc struct {
	Height string `json:"height"`
	Hash string `json:"hash"`
}

func (b *ViaBtc)GetBlockInfo(coin string){
	var body []byte
	switch coin{
	case "btc":
		body=util.GetHttpResponse("GET",configs.Config.BtcApi.ViaBtc)
	case "bch":
		body=util.GetHttpResponse("GET",configs.Config.BchApi.ViaBtc)
	case "ltc":
		body=util.GetHttpResponse("GET",configs.Config.LtcApi.ViaBtc)
	}
	b.Unmarshal(body)
}

func(b *ViaBtc)Unmarshal(body []byte){
	var data map[string]interface{}
	err1:=json.Unmarshal(body,&data)
	if err1!=nil{
		log.Error(err1)
	}
	temp1:=data["data"].(map[string]interface{})
	temp:=temp1["data"].([]interface{})
	out,err2:=json.Marshal(temp[0])
	if err2!=nil{
		log.Error(err2)
	}
	err3:=json.Unmarshal(out,&b)
	if err3!=nil{
		log.Error(err3)
	}
}
