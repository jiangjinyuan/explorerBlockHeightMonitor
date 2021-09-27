package monitor

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"sync"
	"time"

	"github.com/jiangjinyuan/explorerBlockHeightMonitor/utils"

	"github.com/jiangjinyuan/explorerBlockHeightMonitor/senders"

	"github.com/jiangjinyuan/explorerBlockHeightMonitor/configs"
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/models"

	"github.com/jiangjinyuan/explorerBlockHeightMonitor/client"

	"github.com/PaesslerAG/jsonpath"
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/monitor/explorer"
	log "github.com/sirupsen/logrus"
)

type blockHeightMonitorRunner struct {
	once               sync.Once
	explorerJSONConfig string
	confList           []*explorer.Explorer
	exit               chan struct{}
}

type ExplorerConfig struct {
	Explorers []*explorer.Explorer `json:"explorers"`
}

// NewBlockHeightMonitorRunner 初始化
func NewBlockHeightMonitorRunner(explorerJSONConfig string) (*blockHeightMonitorRunner, error) {
	runner := &blockHeightMonitorRunner{
		explorerJSONConfig: explorerJSONConfig,
		exit:               make(chan struct{}),
	}

	if err := runner.loadExplorerJSONByFile(); err != nil {
		log.WithField("fileName", explorerJSONConfig).Error("loadExplorerJSONByFile failed")
		return nil, err
	}

	for _, value := range runner.confList {
		log.WithField("step", "init-list-explorers").WithField("name", value.Name).
			WithField("coin", value.Coin).WithField("HeightJSONPattern",
			value.HeightJSONPattern).WithField("HashJSONPattern", value.HashJSONPattern).
			WithField("url", value.URL).WithField("enabled", value.Enabled)
	}

	return runner, nil
}

// loadExplorerJSONByFile 加载配置
func (r *blockHeightMonitorRunner) loadExplorerJSONByFile() error {
	content, err := ioutil.ReadFile(r.explorerJSONConfig)
	if err != nil {
		log.WithField("fileName", r.explorerJSONConfig).WithField("error", err).Error("loadExplorerJSONByFile failed")
		return err
	}

	return r.loadExplorerJSON(content)
}

// loadExplorerJSON
func (r *blockHeightMonitorRunner) loadExplorerJSON(content []byte) error {
	var data ExplorerConfig
	err := json.Unmarshal(content, &data)
	if err != nil {
		log.WithField("content", string(content)).WithField("error", err).Error("Unmarshal json failed")
		return err
	}
	r.confList = data.Explorers
	return err
}

func (r *blockHeightMonitorRunner) Run() error {
	defer func() {
		if r := recover(); r != nil {
			log.Error("recover from panic")
		}
	}()

	log.Info("run get explorer block height begin...")
	if len(r.confList) == 0 {
		return errors.New("explorer config is empty")
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		r.RunGetBlockInfo(r.confList)
	}()

	go func() {
		defer wg.Done()
		r.RunCheckBlockHeight(r.confList)
	}()

	wg.Wait()

	return nil
}

// 定时器开始运行
func (r *blockHeightMonitorRunner) Start() error {
	timer := time.NewTicker(time.Duration(configs.Config.Interval) * time.Second)
	for {
		select {
		case <-timer.C:
			if err := r.Run(); err != nil {
				log.Errorf("run get block height failed, error: %v ", err)
			}

		case <-r.exit:
			log.Info("run get block height exit...")
			return nil
		}
	}
}

// RunGetBlockHeight 开始运行，获取浏览器块高
func (r *blockHeightMonitorRunner) RunGetBlockInfo(pConfList []*explorer.Explorer) {
	defer func() {
		if err := recover(); err != nil {
			log.Error("panic", err)
		}
	}()

	var result []*models.Block
	for _, conf := range pConfList {
		log.Infof("explorer name:%s, coin name:%s, conf:%v", conf.Name, conf.Coin, *conf)
		if conf.Enabled {
			block, err := r.GetExplorerBlockInfo(conf)
			if err != nil {
				continue
			}
			block.Coin = conf.Coin
			block.ExplorerName = conf.Name
			block.Link = conf.URL
			result = append(result, block)
		}
	}

	// 写入数据库
	if err := models.WriteToExplorerBlockHeight(result); err != nil {
		log.Error("insert data to explorer_block_height table failed")
	}
}

// GetExplorerBlockInfo 获取浏览器块高
func (r *blockHeightMonitorRunner) GetExplorerBlockInfo(pConf *explorer.Explorer) (result *models.Block, err error) {
	// 默认 json 格式
	return r.GetExplorerBlockInfoByJSONFormat(pConf)
}

// GetExplorerBlockInfoByJSONFormat
func (r *blockHeightMonitorRunner) GetExplorerBlockInfoByJSONFormat(pConf *explorer.Explorer) (result *models.Block, err error) {
	result = &models.Block{
		Coin:         pConf.Coin,
		ExplorerName: pConf.Name,
	}
	body, err := client.Client.Get(pConf.URL, pConf.CustomHeaders)
	if err != nil {
		log.Error(err)
		return result, err
	}

	var object interface{}
	err = json.Unmarshal(body, &object)
	if err != nil {
		log.Error(err)
		return result, err
	}
	log.Debugf("body:%v", object)

	// parse block height
	height, err := jsonpath.Get(pConf.HeightJSONPattern, object)
	if err != nil {
		log.Error(err)
		return result, err
	}
	result.Height = r.ParseBlockHeight(height)

	// parse block hash
	hash, err := jsonpath.Get(pConf.HashJSONPattern, object)
	if err != nil {
		log.Error(err)
		return result, err
	}
	result.Hash = hash.(string)
	if pConf.Name == "blockcypher" {
		result.Hash = "0x" + result.Hash
	}

	return result, nil
}

// ParseBlockHeight json 解析块高
func (r *blockHeightMonitorRunner) ParseBlockHeight(object interface{}) (blockHeight int64) {
	blockHeightType := reflect.TypeOf(object)
	switch blockHeightType.String() {
	case "int64":
		blockHeight = object.(int64)
	case "string":
		blockHeightStr := object.(string)
		blockHeight, _ = strconv.ParseInt(blockHeightStr, 10, 64)
	case "float64":
		blockHeight = int64(object.(float64))
	}

	return
}

// RunCheckBlockHeight 开始运行，浏览器块高检查
func (r *blockHeightMonitorRunner) RunCheckBlockHeight(pConfList []*explorer.Explorer) {
	defer func() {
		if err := recover(); err != nil {
			log.Error("panic", err)
		}
	}()

	// get the block height list
	blockList, err := models.GetExplorerBlockInfo(configs.Config.MonitorCoins)
	if err != nil {
		log.Error("get explorer block info failed", err)
		return
	}

	// compare
	for _, coin := range configs.Config.MonitorCoins {
		coinBlockMap := make(map[string]*models.Block)
		for _, value := range blockList {
			if coin != value.Coin {
				continue
			}

			if _, exist := coinBlockMap[value.ExplorerName]; exist {
				continue
			}
			coinBlockMap[value.ExplorerName] = value
		}

		r.CompareBlockHeight(coinBlockMap)
	}
}

func (r *blockHeightMonitorRunner) CompareBlockHeight(coinBlockMap map[string]*models.Block) {
	for _, explorerName := range configs.Config.MonitorExplorers {
		mBlock, exist := coinBlockMap[explorerName]
		if !exist {
			log.Errorf("the monitor explorer %s not exist in config json file", explorerName)
			continue
		}

		heightErrorCnt := 0
		hashErrorCnt := 0
		heightErrorMsg := fmt.Sprintf("Monitor Warning: %s (UTC) \n"+
			"[coin: %s, explorer name: %s, height: %d, url: %s] block height behind others explorer. \n"+
			"explorer list [name:height:link]:\n",
			time.Now().UTC().Format(utils.UTCDatetime), mBlock.Coin, mBlock.ExplorerName, mBlock.Height, mBlock.Link)
		hashErrorMsg := fmt.Sprintf("Monitor Warning:  %s (UTC) \n"+
			"[coin: %s, explorer name: %s, height: %d, hash: %s, url: %s] block hash different from others explorer. \n"+
			"explorer list [name:height:hash:link]:\n",
			time.Now().UTC().Format(utils.UTCDatetime), mBlock.Coin, mBlock.ExplorerName, mBlock.Height, mBlock.Hash, mBlock.Link)
		for key, value := range coinBlockMap {
			if key != explorerName {
				// check block height
				if mBlock.Height < value.Height-configs.Config.AlarmThreshold[value.Coin] {
					msg := fmt.Sprintf("%s:%d:%s \n", value.ExplorerName, value.Height, value.Link)
					heightErrorMsg += msg
					heightErrorCnt++
				}

				// check block hash
				if mBlock.Height == value.Height && mBlock.Hash != value.Hash {
					msg := fmt.Sprintf("%s:%d:%s:%s \n", value.ExplorerName, value.Height, value.Hash, value.Link)
					hashErrorMsg += msg
					hashErrorCnt++
				}
			}
		}

		var sender senders.Senders
		if configs.Config.Slack.IsEnable {
			sender = senders.NewSlackSender()
		} else if configs.Config.Email.IsEnable {
			sender = senders.NewEmailSender()
		}

		if heightErrorCnt > 0 {
			// send message to channel
			r.SendMessage(heightErrorMsg, sender)
		}

		if hashErrorCnt > 0 {
			// send message to channel
			r.SendMessage(hashErrorMsg, sender)
		}
	}
}

func (r *blockHeightMonitorRunner) SendMessage(msg string, sender senders.Senders) {
	sender.Send(msg)
}

func (r *blockHeightMonitorRunner) Close() {
	r.once.Do(func() {
		close(r.exit)
	})
}
