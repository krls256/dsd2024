package di

import (
	"github.com/hazelcast/hazelcast-go-client"
	"github.com/krls256/dsd2024/logging/repositories"
	"github.com/krls256/dsd2024/logging/services"
	hazelcast2 "github.com/krls256/dsd2024/pkg/hazelcast"
	"github.com/sarulabs/di/v2"
)

const (
	HazelcastClientName   = "HazelcastClient"
	LoggingRepositoryName = "LoggingRepository"
	LoggingServiceName    = "LoggingService"
)

func Defs() []di.Def {
	return []di.Def{
		{
			Name: HazelcastClientName,
			Build: func(ctn di.Container) (interface{}, error) {
				return hazelcast2.NewHazelcast(hazelcast2.Config{
					Paths:       []string{"0.0.0.0:5701", "0.0.0.0:5702", "0.0.0.0:5703"},
					ClusterName: "hzc",
				})
			},
		},
		{
			Name: LoggingRepositoryName,
			Build: func(ctn di.Container) (interface{}, error) {
				hc := ctn.Get(HazelcastClientName).(*hazelcast.Client)

				return repositories.NewLoggingRepository(hc), nil
			},
		},
		{
			Name: LoggingServiceName,
			Build: func(ctn di.Container) (interface{}, error) {
				repo := ctn.Get(LoggingRepositoryName).(*repositories.LoggingRepository)

				return services.NewLoggingService(repo), nil
			},
		},
	}
}
