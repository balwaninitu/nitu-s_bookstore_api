package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//c is pointer for gin.context
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
