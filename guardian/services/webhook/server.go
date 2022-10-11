package webhook

import (
	"fmt"
	"net/http"

	"guardian/apprunner"
	"guardian/apprunner/lib/gin"
	"guardian/apprunner/serviceprovider"
	"guardian/apprunner/session"
)

func StartServer(host string, port int) {
	router := gin.New()

	appsession := session.Initialize(&serviceprovider.HttpService{
		Http: &http.Server{
			Handler: router,
			Addr:    fmt.Sprintf("%s:%d", host, port),
		},
	})
	apprunner.RunSession(appsession)
}
