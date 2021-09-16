package cmd

import (
	"time"

	"github.com/jiangjinyuan/explorerBlockHeightMonitor/client"
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/monitor"
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/util"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(generalCoinCmd)
}

var generalCoinCmd = &cobra.Command{
	Use:   "generalCoin",
	Short: "BTC/BCH/BSV/LTC/ETH/ETC Block Height Monitor",
	Long:  "This module can be used to monitor blockHeight of different coins by choosing different coins",
	Run:   GeneralCoinHeightMonitor,
}

func GeneralCoinHeightMonitor(cmd *cobra.Command, args []string) {
	// init log
	util.ConfigLocalFilesystemLogger("./logs/", "monitor.log", 7*time.Hour*24, time.Second*20)
	// init http client
	client.Client = client.NewHTTPClient()

	// 初始化 runner
	runner, err := monitor.NewBlockHeightMonitorRunner("conf/explorer.json")
	if err != nil {
		panic(err)
	}

	if err := runner.Run(); err != nil {
		log.Error(err)
	}
}
