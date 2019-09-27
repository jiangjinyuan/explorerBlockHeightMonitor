package cmd

import (
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/configs"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var cfgFile string
var coin string

var rootCmd = &cobra.Command{
	Use:   "blockHeightMonitor",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "configs/config", "config file (default is $HOME/.blockHeightMonitor.yaml)")
	rootCmd.PersistentFlags().StringVar(&coin, "coin", "btc", "choose coin")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	if cfgFile != "" {
		configs.InitConfig(cfgFile)
	}
	if coin != "" {
		coin = strings.ToLower(coin)
	}

}
