package repositories

import (
	"context"
	"github.com/hazelcast/hazelcast-go-client"
	"github.com/krls256/dsd2024/logging/entities"
	"github.com/samber/lo"
)

const LogsCollectionName = "logs_collection"

func NewLoggingRepository(client *hazelcast.Client) *LoggingRepository {
	return &LoggingRepository{client: client}
}

type LoggingRepository struct {
	client *hazelcast.Client
}

func (r *LoggingRepository) start(ctx context.Context) (*hazelcast.Map, error) {
	return r.client.GetMap(ctx, LogsCollectionName)
}

func (r *LoggingRepository) Save(ctx context.Context, msg entities.Log) error {
	m, err := r.start(ctx)
	if err != nil {
		return err
	}

	_, err = m.Put(ctx, msg.ID, msg.Text)

	return err
}

func (r *LoggingRepository) All(ctx context.Context) ([]string, error) {
	m, err := r.start(ctx)
	if err != nil {
		return nil, err
	}

	vals, err := m.GetValues(ctx)
	if err != nil {
		return nil, err
	}

	return lo.Map(vals, func(item interface{}, index int) string {
		return item.(string)
	}), nil
}
