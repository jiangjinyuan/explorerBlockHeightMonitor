package cmd

import (
	"fmt"
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/configs"
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/dbs"
	"github.com/spf13/cobra"
	"os"
)

var cfgFile string

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
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	if cfgFile != "" {
		configs.InitConfig(cfgFile)
	}

	// init mysql
	config := make(map[string]configs.MySQLDSN)
	for i := range configs.Config.ExplorerDatabase {
		databaseConfig := configs.Config.ExplorerDatabase[i]
		configs.AddDatabaseConfig(&databaseConfig, config)
	}
	dbs.InitMySQLDB(config)

	// init redis
	// _ = dbs.InitRedisDB(configs.Config.Redis.Redis.Address, configs.Config.Redis.Redis.Password)
}
