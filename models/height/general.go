package height

import (
	"database/sql"
	"fmt"
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/configs"
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/models"
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/senders"
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/util"
	"strconv"
	"sync"
)

type BlockHeightMonitor struct {
	wg     sync.WaitGroup
	Height map[string]int64
	Hash   map[string]string
}

func (b *BlockHeightMonitor) Run(coin string, db *sql.DB){
	b.Height = make(map[string]int64)
	b.Hash = make(map[string]string)
	switch coin {
	case "btc":
		b.wg.Add(1)
		go func() {
			var temp1 models.BlockChair
			temp1.GetBlockInfo(coin)
			b.Height["BlockChair"] = temp1.Height
			b.Hash["BlockChair"] = temp1.Hash
			b.wg.Done()
		}()
		b.wg.Add(1)
		go func() {
			var temp2 models.BlockChainBtc
			temp2.GetBlockInfo()
			b.Height["BlockChain"] = temp2.Height
			b.Hash["BlockChain"] = temp2.Hash
			b.wg.Done()
		}()
		b.wg.Add(1)
		go func() {
			var temp3 models.ViaBtc
			temp3.GetBlockInfo(coin)
			height, _ := strconv.Atoi(temp3.Height)
			b.Height["ViaBtc"] = int64(height)
			b.Hash["ViaBtc"] = temp3.Hash
			b.wg.Done()
		}()

	case "bch":
		b.wg.Add(1)
		go func() {
			var temp1 models.BlockChair
			temp1.GetBlockInfo(coin)
			b.Height["BlockChair"] = temp1.Height
			b.Hash["BlockChair"] = temp1.Hash
			b.wg.Done()
		}()
		b.wg.Add(1)
		go func() {
			var temp2 models.Bitcoin
			temp2.GetBlockInfo()
			b.Height["Bitcoin"] = temp2.Height
			b.Hash["Bitcoin"] = temp2.Hash
			b.wg.Done()
		}()
		b.wg.Add(1)
		go func() {
			var temp3 models.ViaBtc
			temp3.GetBlockInfo(coin)
			height, _ := strconv.Atoi(temp3.Height)
			b.Height["ViaBtc"] = int64(height)
			b.Hash["ViaBtc"] = temp3.Hash
			b.wg.Done()
		}()
	case "ltc":
		b.wg.Add(1)
		go func() {
			var temp1 models.BlockCypher
			temp1.GetBlockInfo()
			b.Height["BlockCypher"] = temp1.Height
			b.Hash["BlockCypher"] = temp1.Hash
			b.wg.Done()
		}()
		b.wg.Add(1)
		go func() {
			var temp2 models.BlockChair
			temp2.GetBlockInfo(coin)
			b.Height["BlockChair"] = temp2.Height
			b.Hash["BlockChair"] = temp2.Hash
			b.wg.Done()
		}()
		b.wg.Add(1)
		go func() {
			var temp3 models.ViaBtc
			temp3.GetBlockInfo(coin)
			height, _ := strconv.Atoi(temp3.Height)
			b.Height["ViaBtc"] = int64(height)
			b.Hash["ViaBtc"] = temp3.Hash
			b.wg.Done()
		}()

	case "eth":
		b.wg.Add(1)
		go func() {
			var temp1 models.Etherscan
			temp1.GetBlockInfo()
			b.Height["Etherscan"], _ = strconv.ParseInt(temp1.Height, 0, 64)
			b.Hash["Etherscan"] = temp1.Hash
			b.wg.Done()
		}()
		b.wg.Add(1)
		go func() {
			var temp2 models.BlockChainEth
			temp2.GetBlockInfo()
			b.Height["BlockChain"], _ = strconv.ParseInt(temp2.Height, 0, 64)
			b.Hash["BlockChain"] = temp2.Hash
			b.wg.Done()
		}()
		b.wg.Add(1)
		go func() {
			var temp3 models.BlockChair
			temp3.GetBlockInfo(coin)
			b.Height["BlockChair"] = temp3.Height
			b.Hash["BlockChair"] = temp3.Hash
			b.wg.Done()
		}()
	case "etc":
		b.wg.Add(1)
		go func() {
			var temp1 models.Gastracker
			temp1.GetBlockInfo()
			b.Height["Gastracker"] = temp1.Height
			b.Hash["Gastracker"] = temp1.Hash
			b.wg.Done()
		}()
		b.wg.Add(1)
		go func() {
			var temp2 models.EtcBlockExplorer
			temp2.GetBlockInfo()
			b.Height["EtcBlockExplorer"] = temp2.Height
			b.Hash["EtcBlockExplorer"] = temp2.Hash
			b.wg.Done()
		}()
	}
	b.wg.Add(1)
	var temp4 models.BTCcom
	temp4.GetBlockInfo(coin)
	b.Height["BTCcom"] = temp4.Height
	b.Hash["BTCcom"] = temp4.Hash
	b.wg.Done()

	b.wg.Add(1)
	go func() {
		var temp5 models.Node
		temp5.GetBlockInfo(coin)
		b.Height["Node"] = temp5.Height
		b.Hash["Node"] = temp5.Hash
		b.wg.Done()
	}()
	b.wg.Wait()
	util.Insert(db, coin, b.Height, b.Hash)
	b.Compare(coin)
	//fmt.Println(b.Height)
	//fmt.Println(b.Hash)
}

func (b BlockHeightMonitor) Compare(coin string) {
	text := make(map[string]string)  //write the information which to send
	count := make(map[string]int64)  //Record the blockHeight difference between other explorers and BTC.com
	var N int64   //the AlarmThreshold of every coin
	switch coin {
	case "btc":
		N = configs.Config.AlarmThreshold.Btc
		text["0"] = fmt.Sprintf("The BTC of BTC.com(v3)'s latest blockHeight: ")
		text["0"] += fmt.Sprintf("%d",b.Height["BTCcom"])
		count["BlockChain"] = b.Height["BlockChain"] - b.Height["BTCcom"]
		count["BlockChair"] = b.Height["BlockChair"] - b.Height["BTCcom"]
		count["ViaBtc"] = b.Height["ViaBtc"] - b.Height["BTCcom"]

	case "bch":
		N = configs.Config.AlarmThreshold.Bch
		text["0"] = fmt.Sprintf("The BCH of BTC.com(v3)'s latest blockHeight: ")
		text["0"] += fmt.Sprintf("%d",b.Height["BTCcom"])
		count["Bitcoin"] = b.Height["Bitcoin"] - b.Height["BTCcom"]
		count["BlockChair"] = b.Height["BlockChair"] - b.Height["BTCcom"]
		count["ViaBtc"] = b.Height["ViaBtc"] - b.Height["BTCcom"]

	case "ltc":
		N = configs.Config.AlarmThreshold.Ltc
		text["0"] = fmt.Sprintf("The LTC of BTC.com(v3)'s latest blockHeight: ")
		text["0"] += fmt.Sprintf("%d",b.Height["BTCcom"])
		count["ViaBtc"] = b.Height["ViaBtc"] - b.Height["BTCcom"]
		count["BlockChair"] = b.Height["BlockChair"] - b.Height["BTCcom"]
		count["BlockCypher"] = b.Height["BlockCypher"] - b.Height["BTCcom"]
	case "eth":
		N = configs.Config.AlarmThreshold.Eth
		text["0"] = fmt.Sprintf("The ETH of BTC.com(v3)'s latest blockHeight: ")
		text["0"] += fmt.Sprintf("%d",b.Height["BTCcom"])
		count["Etherscan"] = b.Height["Etherscan"] - b.Height["BTCcom"]
		count["BlockScout"] = b.Height["BlockScout"] - b.Height["BTCcom"]
		count["BlockChair"] = b.Height["BlockChair"] - b.Height["BTCcom"]
	case "etc":
		N = configs.Config.AlarmThreshold.Etc
		text["0"] = fmt.Sprintf("The ETC of BTC.com(v3)'s latest blockHeight: ")
		text["0"] += fmt.Sprintf("%d",b.Height["BTCcom"])
		count["Gastracker"] = b.Height["Gastracker"] - b.Height["BTCcom"]
		count["EtcBlockExplorer"] = b.Height["EtcBlockExplorer"] - b.Height["BTCcom"]
	}

	count["Node"] = b.Height["Node"] - b.Height["BTCcom"]

	for key, result := range count {
		if result >= N {
			text[key] = fmt.Sprintf("Behind the " + key)
			text[key] += fmt.Sprintf(" %d blocks,", result)
			text[key] += fmt.Sprintf(" %s's latest blockHeight is ", key)
			text[key] += fmt.Sprintf(" %d;", b.Height[key])
		} /*else if result < 0 {
			text[key] = fmt.Sprintf("Beyond the " + key)
			text[key] += fmt.Sprintf(" %d blocks,", -result)
			text[key] += fmt.Sprintf(" %s's latest blockHeight is ", key)
			text[key] += fmt.Sprintf(" %d;", b.Height[key])
		}*/
	}

	if len(text) == 1 {
		return
	}
	textHeight:=""
	for key, result := range b.Height {
		tempHeight:=fmt.Sprintf(key+":")
		tempHeight+=fmt.Sprintf("%d  ",result)
		textHeight+=tempHeight
	}
	fmt.Println(textHeight)

	senders.SlackPoster.SendText(text," All latest blockHeight——"+textHeight)  //send alarm info to slack channel
	senders.EmailPublisher.SendText(text," All latest blockHeight——"+textHeight) //send alarm info to email
}
