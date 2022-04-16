package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ping(c *gin.Context) {
	c.String(http.StatusOK, fmt.Sprintf("Pong!"))
}
