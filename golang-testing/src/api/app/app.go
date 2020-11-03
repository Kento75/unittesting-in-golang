package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApp() {
	mapUrls()
	// MEMO: ポート指定だけだと警告が出るのでIPも指定
	if err := router.Run("127.0.0.1:8080"); err != nil {
		panic(err)
	}
}
