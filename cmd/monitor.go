package cmd

import (
	"fmt"
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/client"
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/configs"
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/router"
	"time"

	"github.com/jiangjinyuan/explorerBlockHeightMonitor/monitor"
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(generalCoinCmd)
}

var generalCoinCmd = &cobra.Command{
	Use:   "GeneralCoin",
	Short: "BTC/BCH/BSV/LTC/ETH Block Height Monitor",
	Long:  "This module can be used to monitor blockHeight of different coins by choosing different coins",
	Run:   GeneralCoinHeightMonitor,
}

func GeneralCoinHeightMonitor(cmd *cobra.Command, args []string) {
	// init log
	utils.ConfigLocalFilesystemLogger("./logs/", "monitor.log", 7*time.Hour*24, time.Second*20)
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

	// 开启异步定时器
	go func() {
		if err := runner.Start(); err != nil {
			log.Error(err)
		}
	}()
	defer runner.Close()

	r := router.InitRouter()
	if err := r.Run(fmt.Sprintf(":%d", configs.Config.Health.Port)); err != nil {
		log.Error(err)
	}
}
