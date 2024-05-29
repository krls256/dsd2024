package di

import (
	"encoding/json"
	"github.com/krls256/dsd2024/pkg/consul"
	"github.com/krls256/dsd2024/pkg/hazelcast"
	"github.com/sarulabs/di/v2"
)

const (
	HazelcastClientName = "HazelcastClient"
	HazelcastConfigName = "HazelcastConfig"
)

func Defs() []di.Def {
	return []di.Def{
		{
			Name: HazelcastConfigName,
			Build: func(ctn di.Container) (interface{}, error) {
				bts, err := consul.GetValue("0.0.0.0:8500", "hazelcast")
				if err != nil {
					return nil, err
				}

				cfg := hazelcast.Config{}

				if err := json.Unmarshal(bts, &cfg); err != nil {
					return nil, err
				}

				return cfg, nil
			},
		},
		{
			Name: HazelcastClientName,
			Build: func(ctn di.Container) (interface{}, error) {
				cfg := ctn.Get(HazelcastConfigName).(hazelcast.Config)

				return hazelcast.NewHazelcast(cfg)
			},
		},
	}
}
