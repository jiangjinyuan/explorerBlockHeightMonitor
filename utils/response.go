package utils

import "github.com/gin-gonic/gin"

const (
	Success         = 200
	Error           = 500
	ErrorAuth       = 403
	InvalidParams   = 10001
	InvalidAppID    = 10002
	InvalidCoinType = 10003
	ConfigError     = 10004
	Exist           = 409
)

var MessageFlags = map[int]string{
	Success:         "success",
	Error:           "fail",
	ErrorAuth:       "auth error",
	InvalidParams:   "invalid params",
	InvalidAppID:    "invalid app id",
	InvalidCoinType: "invalid coin type",
	ConfigError:     "config error",
	Exist:           "exists",
}

func GetStatus(status int) int {
	msg, ok := StatusFlags[status]
	if ok {
		return msg
	}
	return StatusFlags[Error]
}

var StatusFlags = map[int]int{
	Success: 0,
	Error:   1,
}

func GetMessage(status int) string {
	msg, ok := MessageFlags[status]
	if ok {
		return msg
	}
	return MessageFlags[Error]
}

func Response(c *gin.Context, httpCode, status int, data interface{}) {
	c.JSON(httpCode, gin.H{
		"status":  GetStatus(status),
		"message": GetMessage(status),
		"data":    data,
	})
}

func ResponseError(c *gin.Context, httpCode, status int, data interface{}) {
	c.JSON(httpCode, gin.H{
		"status":  GetStatus(status),
		"message": GetMessage(status),
		"data":    data,
	})
	c.AbortWithStatus(httpCode)
}

func ResponseList(c *gin.Context, httpCode, status int, list interface{}, count int) {
	limit := c.GetInt("limit")
	var pageTotal int
	if count%limit > 0 {
		pageTotal = (count / limit) + 1
	} else {
		pageTotal = count / limit
	}
	data := gin.H{
		"list":        list,
		"page":        c.GetInt("page"),
		"page_size":   limit,
		"page_total":  pageTotal,
		"total_count": count,
	}
	Response(c, httpCode, status, data)
}

func Redirect(c *gin.Context, httpCode int, url string) {
	c.Redirect(httpCode, url)
}
