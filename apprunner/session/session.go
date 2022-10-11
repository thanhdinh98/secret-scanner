package session

import (
	"context"

	"guardian/apprunner/options"
	"guardian/apprunner/serviceprovider"
	"guardian/common/log"
)

type Session struct {
	Options  options.Options
	Services []serviceprovider.ServiceProvider
	Out      *log.Logger
}

func Initialize(Services ...serviceprovider.ServiceProvider) *Session {
	session := &Session{}
	session.Services = Services
	session.Options = options.Parse()
	session.InitLogger()
	return session
}

func (s *Session) InitLogger() {
	s.Out = &log.Logger{}
}

func (s Session) End(ctx context.Context) (mapErr map[string]error) {
	mapErr = make(map[string]error)
	for _, service := range s.Services {
		if err := service.Stop(ctx); err != nil {
			mapErr[service.Name()] = err
		}
	}
	return
}
