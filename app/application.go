package app

import "github.com/gin-gonic/gin"

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
	router.Run(":8080")

}
