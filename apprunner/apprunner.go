package apprunner

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"guardian/apprunner/session"
)

func RunSession(session *session.Session) {
	var quitChan = make(chan os.Signal, 1)
	session.Out.Info("services are starting...\n")
	for i := range session.Services {
		service := session.Services[i]
		go func() {
			session.Out.Info("service `%v` has started\n", service.Name())
			err := service.Start()
			if err != nil {
				session.Out.Fatal("service `%v` run failed\n", service.Name())
			}
		}()
	}
	signal.Notify(quitChan, os.Interrupt, syscall.SIGTERM)
	<-quitChan
	for i := range session.Services {
		session.Services[i].BeforeStop()
	}
	session.Out.Info("servies are shutting down...\n")
	ctx, cancel := context.WithTimeout(context.Background(), session.Options.GracefulTimeout)
	defer cancel()
	var (
		errMap      = session.End(ctx)
		stopOkCount = 0
	)
	for _, service := range session.Services {
		svcName := service.Name()
		if err, hasError := errMap[svcName]; !hasError {
			session.Out.Info("service `%v` has stopped\n", svcName)
			stopOkCount++
			continue
		} else {
			session.Out.Error("service `%v` has failed to stop | err=%s\n", svcName, err.Error())
		}
	}
	if stopOkCount < len(session.Services) {
		session.Out.Info("services haven't shut down properly (%v/%v)", stopOkCount, len(session.Services))
	} else {
		session.Out.Info("services have all shut down (%v/%v).", stopOkCount, len(session.Services))
	}
}
