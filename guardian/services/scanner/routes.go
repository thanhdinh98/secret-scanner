package scanner

import (
	v1 "guardian/guardian/services/scanner/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	rootGroup := router.Group("/api")
	v1.InitGroup(rootGroup.Group("/v1"))
}
