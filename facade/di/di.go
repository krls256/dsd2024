package di

import (
	hazelcast2 "github.com/krls256/dsd2024/pkg/hazelcast"
	"github.com/sarulabs/di/v2"
)

const (
	HazelcastClientName = "HazelcastClient"
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
	}
}
