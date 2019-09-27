package models

import (
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/configs"
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/util"
	"encoding/json"
)

type CryptoID struct {
	Height int64
}

func (b *CryptoID)GetBlockInfo(){
	body := util.GetHttpResponse("GET",configs.Config.LtcApi.BlockChair)
	json.Unmarshal(body,&b.Height)
}
