package v1

import (
	"github.com/gin-gonic/gin"
)

func InitGroup(group *gin.RouterGroup) {
	group.POST("/user", GenerateUserAccessToken)
}
