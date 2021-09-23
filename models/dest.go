package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/jiangjinyuan/explorerBlockHeightMonitor/util"

	"github.com/jiangjinyuan/explorerBlockHeightMonitor/dbs"

	log "github.com/sirupsen/logrus"
)

type Block struct {
	Coin         string    `json:"coin"`
	ExplorerName string    `json:"explorer_name"`
	Height       int64     `json:"height"`
	Hash         string    `json:"hash"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type HeartBeat struct {
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetExplorerBlockInfo(coinList []string) (result map[string]*Block, err error) {
	var object []*Block
	result = make(map[string]*Block)
	conn := ""
	if exists := dbs.CheckDBConnExists(conn); !exists {
		log.WithField("conn", conn).Error("The DB connection not exists!")
		return result, errors.New("the DB connection not exists")
	}

	sqlQuery := "SELECT `coin`, `explorer_name`, `height`, `hash` FROM explorer_block_info WHERE `coin` IN (?)"
	query, args, _ := sqlx.In(sqlQuery, coinList)
	query = dbs.DBMaps[conn].Rebind(query)
	err = dbs.DBMaps[conn].Select(&object, query, args...)
	if err != nil {
		log.WithFields(log.Fields{"coins": coinList, "func": "GetExplorerBlockInfo"}).Error(err.Error())
		return result, err
	}

	for _, value := range object {
		key := fmt.Sprintf("%s-%s", value.Coin, value.ExplorerName)
		if _, exists := result[key]; !exists {
			result[key] = value
		}
	}

	return result, nil
}

func WriteToExplorerBlockHeight(items []*Block) error {
	conn := ""
	if exists := dbs.CheckDBConnExists(conn); !exists {
		log.WithField("conn", conn).Error("The DB connection not exists!")
		return errors.New("the DB connection not exists")
	}

	currentTime := time.Now().UTC().Format(util.UTCDatetime)
	DB := dbs.DBMaps[conn]
	insert, err := DB.Prepare("INSERT INTO explorer_block_info VALUES( ?, ?, ?, ?, ?, ?) ON DUPLICATE KEY UPDATE " +
		"height = ?, hash = ?, updated_at = ?")
	if err != nil {
		log.Error(err)
		return err
	}
	tx, _ := DB.Begin()
	for _, item := range items {
		_, err := tx.Stmt(insert).Exec(
			item.Coin, item.ExplorerName, item.Height, item.Hash, currentTime, currentTime,
			item.Height, item.Hash, currentTime)
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
