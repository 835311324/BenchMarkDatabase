package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HostAddOne(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"api": "HostAddOne"})
}

func HostDeleteOne(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"api": "HostDeleteOne"})
}

func HostListOne(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"api": "HostListOne"})
}

func HostList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"api": "HostList"})
}
