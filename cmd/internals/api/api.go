package api

import (
	"log"
	"strconv"

	"github.com/ThrynSec/GoDumpPayload/cmd/internals/config"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func MapUrls() {
	router.GET(config.EndpointName, func(ctx *gin.Context) {
		ctx.Header("Content-Description", "File Transfer")
		ctx.Header("Content-Transfer-Encoding", "binary")
		ctx.Header("Content-Disposition", "attachment; filename="+config.EndpointName)
		ctx.Header("Content-Type", "application/octet-stream")
		ctx.File(config.FilePath)
	})
	log.Println("Endpoint created")
	router.Run(":" + strconv.Itoa(config.Port))
}
