package monitor

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"reflect"
	"strconv"

	"github.com/jiangjinyuan/explorerBlockHeightMonitor/models"

	"github.com/jiangjinyuan/explorerBlockHeightMonitor/client"

	"github.com/PaesslerAG/jsonpath"
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/monitor/explorer"
	log "github.com/sirupsen/logrus"
)

type blockHeightMonitorRunner struct {
	explorerJSONConfig string
	confList           []*explorer.Explorer
}

type ExplorerConfig struct {
	Explorers []*explorer.Explorer `json:"explorers"`
}

// NewBlockHeightMonitorRunner 初始化
func NewBlockHeightMonitorRunner(explorerJsonConfig string) (*blockHeightMonitorRunner, error) {
	runner := &blockHeightMonitorRunner{
		explorerJSONConfig: explorerJsonConfig,
	}

	if err := runner.loadExplorerJsonByFile(); err != nil {
		log.WithField("fileName", explorerJsonConfig).Error("loadExplorerJsonByFile failed")
		return nil, err
	}

	for _, value := range runner.confList {
		log.WithField("step", "init-list-explorers").WithField("name", value.Name).
			WithField("coin", value.Coin).WithField("JsonPattern",
			value.JsonPattern).WithField("url", value.Url).WithField("enabled",
			value.Enabled)
	}

	return runner, nil
}

// loadExplorerJsonByFile 加载配置
func (r *blockHeightMonitorRunner) loadExplorerJsonByFile() error {
	content, err := ioutil.ReadFile(r.explorerJSONConfig)
	if err != nil {
		log.WithField("fileName", r.explorerJSONConfig).WithField("error", err).Error("loadExplorerJsonByFile failed")
		return err
	}

	return r.loadExplorerJson(content)
}

// loadExplorerJson
func (r *blockHeightMonitorRunner) loadExplorerJson(content []byte) error {
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

	r.RunGetBlockHeight(r.confList)

	return nil
}

// RunGetBlockHeight 开始运行，获取浏览器块高
func (r *blockHeightMonitorRunner) RunGetBlockHeight(pConfList []*explorer.Explorer) {
	defer func() {
		if err := recover(); err != nil {
			log.Error("panic", err)
		}
	}()

	var result []*models.BlockHeight
	for _, conf := range pConfList {
		log.Info(conf.Name, conf.Coin, *conf)
		if conf.Enabled {
			height, err := r.GetExplorerBlockHeight(conf)
			if err != nil {
				continue
			}
			block := &models.BlockHeight{
				Coin:         conf.Coin,
				ExplorerName: conf.Name,
				Height:       height,
			}
			result = append(result, block)
		}
	}

	// 写入数据库
	if err := models.WriteToExplorerBlockHeight(result); err != nil {
		log.Error("insert data to explorer_block_height table failed")
	}
}

// GetExplorerBlockHeight 获取浏览器块高
func (r *blockHeightMonitorRunner) GetExplorerBlockHeight(pConf *explorer.Explorer) (int64, error) {
	// 默认 json 格式
	return r.GetExplorerBlockHeightByJsonFormat(pConf)
}

// GetPoolHashrateByJsonFormat
func (r *blockHeightMonitorRunner) GetExplorerBlockHeightByJsonFormat(pConf *explorer.Explorer) (blockHeight int64, err error) {
	body, err := client.Client.Get(pConf.Url, pConf.CustomHeaders)
	if err != nil {
		log.Error(err)
		return blockHeight, err
	}
	log.WithField("body", body).Debug("get body")

	var object interface{}
	if err := json.Unmarshal(body, &object); err != nil {
		log.Error(err)
		return blockHeight, err
	}

	h, err := jsonpath.Get(pConf.JsonPattern, object)
	if err != nil {
		log.Error(err)
		return blockHeight, err
	}

	blockHeight = r.ParseBlockHeight(h)
	return blockHeight, nil
}

// ParseBlockHeight json 解析块高
func (r *blockHeightMonitorRunner) ParseBlockHeight(object interface{}) (blockHeight int64) {
	blockHeightType := reflect.TypeOf(object)
	switch blockHeightType.String() {
	case "int64":
		blockHeight = object.(int64)
		break
	case "string":
		blockHeightStr := object.(string)
		blockHeight, _ = strconv.ParseInt(blockHeightStr, 10, 64)
		break
	}

	return
}
