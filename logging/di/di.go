package di

import (
	"github.com/hazelcast/hazelcast-go-client"
	"github.com/krls256/dsd2024/logging/handlers"
	"github.com/krls256/dsd2024/logging/repositories"
	"github.com/krls256/dsd2024/logging/services"
	pkgHazelcast "github.com/krls256/dsd2024/pkg/hazelcast"
	"github.com/sarulabs/di/v2"
)

const (
	HazelcastClientName = "HazelcastClient"
	HazelcastConfigName = "HazelcastConfig"

	LoggingRepositoryName = "LoggingRepository"
	LoggingServiceName    = "LoggingService"
	LoggingHandlerName    = "LoggingHandler"
)

func Defs() []di.Def {
	return []di.Def{
		{
			Name: LoggingRepositoryName,
			Build: func(ctn di.Container) (interface{}, error) {
				hc := ctn.Get(HazelcastClientName).(*hazelcast.Client)
				cfg := ctn.Get(HazelcastConfigName).(pkgHazelcast.Config)

				return repositories.NewLoggingRepository(hc, cfg), nil
			},
		},
		{
			Name: LoggingServiceName,
			Build: func(ctn di.Container) (interface{}, error) {
				repo := ctn.Get(LoggingRepositoryName).(*repositories.LoggingRepository)

				return services.NewLoggingService(repo), nil
			},
		},
		{
			Name: LoggingHandlerName,
			Build: func(ctn di.Container) (interface{}, error) {
				ls := ctn.Get(LoggingServiceName).(*services.LoggingService)

				return handlers.NewLoggingHandler(ls), nil
			},
		},
	}
}
