package v1

import (
	"net/http"

	"guardian/guardian/lib/management/user"
	"guardian/guardian/lib/management/user/session"

	"github.com/gin-gonic/gin"
)

func GenerateUserAccessToken(c *gin.Context) {
	var reqModel GenerateUserAccessTokenRequest
	if err := c.BindJSON(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	session := session.Initialize(reqModel.Email, reqModel.Password)
	accesstoken, err := user.GenerateAccessToken(c, session)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, GenerateUserAccessTokenResponse{
		AccessToken: accesstoken,
	})
	return
}

func CheckUserSessionExists(c *gin.Context) {

}
