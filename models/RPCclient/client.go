package RPCclient

import (
	rpc "github.com/KeisukeYamashita/go-jsonrpc"
	log "github.com/sirupsen/logrus"
	"strconv"
)

type BasicAuth struct {
	Username string
	Password string
}

type Client struct {
	*rpc.RPCClient
	height    int64
	ethHeight string
}

type EthGetBlockHash struct {
	Hash   string `json:"hash"`
	Height string `json:"number"`
}

func (c *Client) NewRPCClient(endpoint string, basicAuth *BasicAuth) {
	c.RPCClient = rpc.NewRPCClient(endpoint)
	c.RPCClient.SetBasicAuth(basicAuth.Username, basicAuth.Password)
}

//GetBlockCount gets the latest block height.
func (c *Client) GetBlockCount(coin string) int64 {
	method := ""
	out := int64(0)
	var resp *rpc.RPCResponse
	var err error
	switch coin {
	case "btc", "bch", "ltc":
		method = "getblockcount"
		resp, err = c.RPCClient.Call(method)
		if err != nil {
			log.Error(err)
			return -1
		}
		if resp.Error != nil {
			log.Error(resp.Error)
			return -1
		}
		err = resp.GetObject(&c.height)
		if err != nil {
			log.Error("err")
			return -1
		}
		out = c.height
	case "eth":
		method = "eth_blockNumber"
		resp, err = c.RPCClient.Call(method)
		if err != nil {
			log.Error(err)
			return -1
		}
		if resp.Error != nil {
			log.Error(resp.Error)
			return -1
		}
		resp.GetObject(&c.ethHeight)
		out, _ = strconv.ParseInt(c.ethHeight, 0, 64)
	}

	return out
}

//GetBlockHash gets the latest block height's hash.
func (c *Client) GetBlockHash(coin string) string {
	method := ""
	var resp *rpc.RPCResponse
	var err error
	var hash string
	switch coin {
	case "btc", "bch", "ltc":
		method = "getblockhash"
		resp, err = c.RPCClient.Call(method, c.height)
		if err != nil {
			log.Error(err)
			return "-1"
		}
		if resp.Error != nil {
			log.Error(resp.Error)
			return "-1"
		}
		err = resp.GetObject(&hash)
		if err != nil {
			log.Error(err)
			return "-1"
		}

	case "eth":
		method = "eth_getBlockByNumber"
		resp, err = c.RPCClient.Call(method, c.ethHeight, true)
		if err != nil {
			log.Error(err)
			return "-1"
		}
		if resp.Error != nil {
			log.Error(resp.Error)
			return "-1"
		}
		var count map[string]string
		resp.GetObject(&count)
		hash = count["hash"]
	}
	return hash
}
