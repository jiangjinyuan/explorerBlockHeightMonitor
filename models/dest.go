package models

import (
	"errors"
	"time"

	"github.com/jiangjinyuan/explorerBlockHeightMonitor/util"

	"github.com/jiangjinyuan/explorerBlockHeightMonitor/dbs"

	log "github.com/sirupsen/logrus"
)

type BlockHeight struct {
	Coin         string    `json:"coin"`
	ExplorerName string    `json:"explorer_name"`
	Height       int64     `json:"height"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type HeartBeat struct {
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updated_at"`
}

func WriteToExplorerBlockHeight(items []*BlockHeight) error {
	conn := ""
	if exists := dbs.CheckDBConnExists(conn); !exists {
		log.WithField("conn", conn).Error("The DB connection not exists!")
		return errors.New("the DB connection not exists")
	}

	currentTime := time.Now().UTC().Format(util.UTCDatetime)
	DB := dbs.DBMaps[conn]
	insert, err := DB.Prepare("INSERT INTO explorer_block_height VALUES( ?, ?, ?, ?, ?) ON DUPLICATE KEY UPDATE " +
		"height = ?, updated_at = ?")
	if err != nil {
		log.Error(err)
		return err
	}
	tx, _ := DB.Begin()
	for _, item := range items {
		_, err := tx.Stmt(insert).Exec(
			item.Coin, item.ExplorerName, item.Height, currentTime, currentTime, item.Height, currentTime)
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
