package di

import (
	"github.com/hazelcast/hazelcast-go-client"
	"github.com/krls256/dsd2024/messages/handlers"
	"github.com/krls256/dsd2024/messages/services"
	"github.com/sarulabs/di/v2"
)

const (
	HazelcastClientName = "HazelcastClient"
	MessagesServiceName = "MessagesService"
	MessagesHandlerName = "MessagesHandler"
)

func Defs() []di.Def {
	return []di.Def{
		{
			Name: MessagesServiceName,
			Build: func(ctn di.Container) (interface{}, error) {
				hc := ctn.Get(HazelcastClientName).(*hazelcast.Client)

				return services.NewMessageService(hc)
			},
		},
		{
			Name: MessagesHandlerName,
			Build: func(ctn di.Container) (interface{}, error) {
				ms := ctn.Get(MessagesServiceName).(*services.MessageService)

				return handlers.NewMessagesHandler(ms), nil
			},
		},
	}
}
