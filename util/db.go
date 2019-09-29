package util

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/configs"
	log "github.com/sirupsen/logrus"
	"time"
)

func InitDB() (db *sql.DB) {
	//打开一个数据库
	db, err := sql.Open(configs.Config.DB.Driver, configs.Config.DB.UserName+":"+configs.Config.DB.Password+"@tcp("+configs.Config.DB.Host+":"+configs.Config.DB.Port+")/"+configs.Config.DB.Schema+"?readTimeout=6m&writeTimeout=6m&timeout=10m&charset=utf8")
	if err != nil {
		log.Error("connect to db failed!")
		panic(err)
		return
	}

	db.SetMaxOpenConns(20)               //Set the maximum number of concurrently open connections
	db.SetMaxIdleConns(15)               //sets maximum of idle connections to be retained in the connection pool.
	db.SetConnMaxLifetime(time.Hour * 3) //sets the maximum length of time that a connection can be reused for.
	return db
}

func Insert(db *sql.DB, coin string, Height map[string]int64, Hash map[string]string) {
	switch coin {
	case "btc":
		s := fmt.Sprintf("insert into %s.%s(BlockChain_height,BlockChair_height,ViaBtc_height,BTCcom_height,Node_height,BlockChain_hash,BlockChair_hash,ViaBtc_hash,BTCcom_hash,Node_hash) values(?,?,?,?,?,?,?,?,?,?)", configs.Config.DB.Schema, coin)
		_, err := db.Exec(s, Height["BlockChain"], Height["BlockChair"], Height["ViaBtc"], Height["BTC.com"], Height["Node"], Hash["BlockChain"], Hash["BlockChair"], Hash["ViaBtc"], Hash["BTC.com"], Hash["Node"])
		if err != nil {
			log.Error(err)
			panic(err)
		}

	case "bch":
		s := fmt.Sprintf("insert into %s.%s(BlockChair_height,Bitcoin_height,ViaBtc_height,BTCcom_height,Node_height,BlockChair_hash,Bitcoin_hash,ViaBtc_hash,BTCcom_hash,Node_hash) values(?,?,?,?,?,?,?,?,?,?)", configs.Config.DB.Schema, coin)
		_, err := db.Exec(s, Height["BlockChair"], Height["Bitcoin"], Height["ViaBtc"], Height["BTC.com"], Height["Node"], Hash["BlockChair"], Hash["Bitcoin"], Hash["ViaBtc"], Hash["BTC.com"], Hash["Node"])
		if err != nil {
			log.Error(err)
			panic(err)
		}

	case "ltc":
		s := fmt.Sprintf("insert into %s.%s(BlockChair_height,ViaBtc_height,BlockCypher_height,BTCcom_height,Node_height,BlockChair_hash,ViaBtc_hash,BlockCypher_hash,BTCcom_hash,Node_hash) values(?,?,?,?,?,?,?,?,?,?)", configs.Config.DB.Schema, coin)
		_, err := db.Exec(s, Height["BlockChair"], Height["ViaBtc"], Height["BlockCypher"], Height["BTC.com"], Height["Node"], Hash["BlockChair"], Hash["ViaBtc"], Hash["BlockCypher"], Hash["BTC.com"], Hash["Node"])
		if err != nil {
			log.Error(err)
			panic(err)
		}

	case "eth":
		s := fmt.Sprintf("insert into %s.%s(Etherscan_height,BlockChain_height,BlockChair_height,BTCcom_height,Node_height,Etherscan_hash,BlockChain_hash,BlockChair_hash,BTCcom_hash,Node_hash) values(?,?,?,?,?,?,?,?,?,?)", configs.Config.DB.Schema, coin)
		_, err := db.Exec(s, Height["Etherscan"], Height["BlockChain"], Height["BlockChair"], Height["BTC.com"], Height["Node"], Hash["Etherscan"], Hash["BlockChain"], Hash["BlockChair"], Hash["BTC.com"], Hash["Node"])
		if err != nil {
			log.Error(err)
			panic(err)
		}
	case "etc":
		//db.Exec("insert into ?.?(BlockChain,BlockChair,TokenView,BTCcom) values(?,?,?,?)", configs.Config.DB.Schema, coin, Height["BlockChain"], Height["BlockChair"], Height["TokenView"], Height["BTC.com"])
	}

}
