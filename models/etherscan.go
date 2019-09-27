package models

import (
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/configs"
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/util"
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

type Etherscan struct {
	Height string `json:"number"`
	Hash string `json:"hash"`
}

func (b *Etherscan)GetBlockInfo(){
	body := util.GetHttpResponse("GET",configs.Config.EthApi.Etherscan)
	b.Unmarshal(body)
}

func(b *Etherscan)Unmarshal(body []byte){
	var data map[string]interface{}
	err1:=json.Unmarshal(body,&data)
	if err1!=nil{
		log.Error(err1)
	}
	temp:=data["result"].(map[string]interface{})
	out,err2:=json.Marshal(temp)
	if err2!=nil{
		log.Error(err2)
	}
	err3:=json.Unmarshal(out,&b)
	if err3!=nil{
		log.Error(err3)
	}
}