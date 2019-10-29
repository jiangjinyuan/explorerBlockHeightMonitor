package cmd

import (
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/models/height"
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/util"
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/configs"
	"github.com/spf13/cobra"
	"time"
)

func init(){
	rootCmd.AddCommand(generalCoinCmd)
}

var generalCoinCmd=&cobra.Command{
	Use:   "GeneralCoin",
	Short: "BTC/BCH/LTC/ETH/ETC Block Height Monitor",
	Long: "This module can be used to monitor blockHeight of different coins by choosing different coins",
	Run:   GeneralCoinHeightMonitor,
}

func GeneralCoinHeightMonitor(cmd *cobra.Command, args []string) {
	var Time time.Duration
	switch coin {
	case "btc":
		util.ConfigLocalFilesystemLogger("./logs/", "btcMonitor.log", 7*time.Hour*24, time.Second*20)
		Time= time.Duration(configs.Config.TraverseSleepTime.Btc) * time.Second
	case "bch":
		util.ConfigLocalFilesystemLogger("./logs/", "bchMonitor.log", 7*time.Hour*24, time.Second*20)
		Time= time.Duration(configs.Config.TraverseSleepTime.Bch) * time.Second
	case "ltc":
		util.ConfigLocalFilesystemLogger("./logs/", "ltcMonitor.log", 7*time.Hour*24, time.Second*20)
		Time= time.Duration(configs.Config.TraverseSleepTime.Ltc) * time.Second
	case "eth":
		util.ConfigLocalFilesystemLogger("./logs/", "ethMonitor.log", 7*time.Hour*24, time.Second*20)
		Time= time.Duration(configs.Config.TraverseSleepTime.Eth) * time.Second
	case "etc":
		util.ConfigLocalFilesystemLogger("./logs/", "etcMonitor.log", 7*time.Hour*24, time.Second*20)
		Time= time.Duration(configs.Config.TraverseSleepTime.Etc) * time.Second
	}
	monitor := height.BlockHeightMonitor{}
	db := util.InitDB()
	defer db.Close()
	for{
		monitor.Run(coin,db)
		time.Sleep(Time)
	}

}
