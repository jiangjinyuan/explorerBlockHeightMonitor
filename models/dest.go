package models

import (
	"errors"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/jiangjinyuan/explorerBlockHeightMonitor/utils"

	"github.com/jiangjinyuan/explorerBlockHeightMonitor/dbs"

	log "github.com/sirupsen/logrus"
)

type Block struct {
	Coin         string `json:"coin" db:"coin"`
	ExplorerName string `json:"explorer_name" db:"explorer_name"`
	Height       int64  `json:"height" db:"height"`
	Hash         string `json:"hash" db:"hash"`
	Link         string `json:"link" db:"link"`
	CreatedAt    string `json:"created_at" db:"created_at"`
	UpdatedAt    string `json:"updated_at" db:"updated_at"`
}

func GetExplorerBlockInfo(coinList []string) (result []*Block, err error) {
	conn := "explorer:general:read"
	if exists := dbs.CheckDBConnExists(conn); !exists {
		log.WithField("conn", conn).Error("The DB connection not exists!")
		return result, errors.New("the DB connection not exists")
	}

	sqlQuery := "SELECT `coin`, `explorer_name`, `height`, `hash`, `link` FROM explorer_block_info WHERE `coin` IN (?)"
	query, args, _ := sqlx.In(sqlQuery, coinList)
	query = dbs.DBMaps[conn].Rebind(query)
	err = dbs.DBMaps[conn].Select(&result, query, args...)
	if err != nil {
		log.WithFields(log.Fields{"coins": coinList, "func": "GetExplorerBlockInfo"}).Error(err.Error())
		return result, err
	}

	return result, nil
}

func GetExplorerBlockInfoByTime(time string) (result []*Block, err error) {
	conn := "explorer:general:read"
	if exists := dbs.CheckDBConnExists(conn); !exists {
		log.WithField("conn", conn).Error("The DB connection not exists!")
		return result, errors.New("the DB connection not exists")
	}
	sqlQuery := "SELECT `coin`, `explorer_name`, `height`, `hash`, `link`, `created_at`, `updated_at` FROM explorer_block_info WHERE `updated_at` >= ?"
	err = dbs.DBMaps[conn].Select(&result, sqlQuery, time)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetExplorerBlockInfoByTime", "time": time}).Error(err.Error())
		return result, err
	}

	return result, nil
}

func WriteToExplorerBlockHeight(items []*Block) error {
	conn := "explorer:general:write"
	if exists := dbs.CheckDBConnExists(conn); !exists {
		log.WithField("conn", conn).Error("The DB connection not exists!")
		return errors.New("the DB connection not exists")
	}

	currentTime := time.Now().UTC().Format(utils.UTCDatetime)
	DB := dbs.DBMaps[conn]
	insert, err := DB.Prepare("INSERT INTO explorer_block_info VALUES( ?, ?, ?, ?, ?, ?, ?) ON DUPLICATE KEY UPDATE " +
		"height = ?, hash = ?, link = ?, updated_at = ?")
	if err != nil {
		log.Error(err)
		return err
	}
	tx, _ := DB.Begin()
	for _, item := range items {
		_, err := tx.Stmt(insert).Exec(
			item.Coin, item.ExplorerName, item.Height, item.Hash, item.Link, currentTime, currentTime,
			item.Height, item.Hash, item.Link, currentTime)
		if err != nil {
			log.Error(err)
			return err
		}
	}
	if err := tx.Commit(); err != nil {
		log.Error(err)
		return err
	}

	return nil
}
