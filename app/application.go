package app

import (
	"github.com/balwaninitu/nitu-s_bookstore_api/logger"
	"github.com/gin-gonic/gin"
)

var (
	//default return engine with logger and recovery middleware
	//every request get handle by gingonic router
	//router will create different goroutine for different request/handle
	router = gin.Default()
)

//gingonic is web framework in go with martini-like API but much better
//when below func starts it map the url
func StartApplication() {
	mapUrls()

	logger.Info("about to start the application")
	router.Run(":8080")

}
