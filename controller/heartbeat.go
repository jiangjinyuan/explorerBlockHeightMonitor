package controller

import (
	"net/http"
	"time"

	"github.com/jiangjinyuan/explorerBlockHeightMonitor/utils"

	"github.com/gin-gonic/gin"
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/configs"
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/models"
)

func Ping(ctx *gin.Context) {
	nowTimeSecond := time.Now().Unix()
	nowTimeSecond -= int64(configs.Config.Health.IntervalSeconds)
	results, err := models.GetExplorerBlockInfoByTime(time.Unix(nowTimeSecond, 0).UTC().Format(utils.UTCDatetime))
	if err != nil || len(results) == 0 {
		utils.ResponseError(ctx, http.StatusInternalServerError, utils.Error, nil)
		return
	}

	utils.Response(ctx, http.StatusOK, utils.Success, results)
}
