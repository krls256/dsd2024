package consul

import (
	"github.com/google/uuid"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/samber/lo"
	"log/slog"
	"math/rand/v2"
)

func Register(consulAddress string, serviceType, host string, port uint16) (cancel func(), err error) {
	config := consulapi.DefaultConfig()
	config.Address = consulAddress

	consul, err := consulapi.NewClient(config)
	if err != nil {
		return func() {}, err
	}

	serviceID := uuid.NewString()

	registration := &consulapi.AgentServiceRegistration{
		ID:      serviceID,
		Name:    serviceType,
		Port:    int(port),
		Address: host,
	}

	if err := consul.Agent().ServiceRegister(registration); err != nil {
		return func() {}, err
	}

	return func() {
		if err = consul.Agent().ServiceDeregister(serviceID); err != nil {
			slog.Error("cant deregister", "err", err)
		}
	}, nil
}

type ServiceAddress struct {
	Host string `json:"host"`
	Port uint16 `json:"port"`
}

func Discover(consulAddress, serviceType string) ([]ServiceAddress, error) {
	config := consulapi.DefaultConfig()
	config.Address = consulAddress

	consul, err := consulapi.NewClient(config)
	if err != nil {
		return nil, err
	}

	s, _, err := consul.Catalog().Service(serviceType, "", nil)
	if err != nil {
		return nil, err
	}

	return lo.Map(s, func(item *consulapi.CatalogService, index int) ServiceAddress {
		return ServiceAddress{
			Host: item.ServiceAddress,
			Port: uint16(item.ServicePort),
		}
	}), nil
}

func DiscoverRandom(consulAddress, serviceType string) (ServiceAddress, error) {
	addresses, err := Discover(consulAddress, serviceType)
	if err != nil || len(addresses) == 0 {
		return ServiceAddress{}, err
	}

	return addresses[rand.IntN(len(addresses))], nil
}

func GetValue(consulAddress string, key string) ([]byte, error) {
	config := consulapi.DefaultConfig()
	config.Address = consulAddress

	consul, err := consulapi.NewClient(config)
	if err != nil {
		return nil, err
	}

	kv, _, err := consul.KV().Get(key, nil)
	if err != nil {
		return nil, err
	}

	return kv.Value, nil
}
