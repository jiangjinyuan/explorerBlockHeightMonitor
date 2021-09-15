package cmd

import (
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/models/height"
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/util"
	"github.com/spf13/cobra"
	"time"
)

func init() {
	rootCmd.AddCommand(generalCoinCmd)
}

var generalCoinCmd = &cobra.Command{
	Use:   "generalCoin",
	Short: "BTC/BCH/LTC/ETH/ETC Block Height Monitor",
	Long:  "This module can be used to monitor blockHeight of different coins by choosing different coins",
	Run:   GeneralCoinHeightMonitor,
}

func GeneralCoinHeightMonitor(cmd *cobra.Command, args []string) {
	util.ConfigLocalFilesystemLogger("./logs/", "monitor.log", 7*time.Hour*24, time.Second*20)
	monitor := height.BlockHeightMonitor{}
	db := util.InitDB()
	defer db.Close()
	monitor.Run(coins, db)
}
