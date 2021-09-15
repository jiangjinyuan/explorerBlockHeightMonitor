package cmd

import (
	"fmt"
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/configs"
	"github.com/spf13/cobra"
	"os"
)

var cfgFile string

var coins string

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
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "configs/config", "config file (default is $HOME/config.yaml)")
	rootCmd.PersistentFlags().StringVar(&coins, "coins", "", "choose coins, for example: btc,bch,bsv")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	if cfgFile != "" {
		configs.InitConfig(cfgFile)
	}

	if coins != "" {
		coins = configs.Config.SupportCoins
	}
}
