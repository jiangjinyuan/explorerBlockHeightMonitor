package models

import (
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/configs"
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/models/RPCclient"
)

type Node struct {
	Height int64
	Hash string
}

func(b *Node) GetBlockInfo(coin string){
	endpoint:=""
	basicAuth := &RPCclient.BasicAuth{}
	switch coin{
	case "btc":
		endpoint=configs.Config.BtcNode.ServerEndPoint
		basicAuth.Username=configs.Config.BtcNode.UserName
		basicAuth.Password=configs.Config.BtcNode.Password
	case "bch":
		endpoint=configs.Config.BchNode.ServerEndPoint
		basicAuth.Username=configs.Config.BchNode.UserName
		basicAuth.Password=configs.Config.BchNode.Password
	case "ltc":
		endpoint=configs.Config.LtcNode.ServerEndPoint
		basicAuth.Username=configs.Config.LtcNode.UserName
		basicAuth.Password=configs.Config.LtcNode.Password
	case "eth":
		endpoint=configs.Config.EthNode.ServerEndPoint
		basicAuth.Username=configs.Config.EthNode.UserName
		basicAuth.Password=configs.Config.EthNode.Password
	case "etc":
		endpoint=configs.Config.EtcNode.ServerEndPoint
		basicAuth.Username=configs.Config.EtcNode.UserName
		basicAuth.Password=configs.Config.EtcNode.Password
	}
	c:=RPCclient.Client{}
	c.NewRPCClient(endpoint, basicAuth)
	b.Height= c.GetBlockCount(coin)
	b.Hash=c.GetBlockHash(coin)
}