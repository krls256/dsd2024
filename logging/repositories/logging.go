package repositories

import (
	"context"
	"github.com/hazelcast/hazelcast-go-client"
	"github.com/krls256/dsd2024/logging/entities"
	pkgHazelcast "github.com/krls256/dsd2024/pkg/hazelcast"
	"github.com/samber/lo"
)

func NewLoggingRepository(client *hazelcast.Client, cfg pkgHazelcast.Config) *LoggingRepository {
	return &LoggingRepository{
		client: client,
		cfg:    cfg,
	}
}

type LoggingRepository struct {
	client *hazelcast.Client
	cfg    pkgHazelcast.Config
}

func (r *LoggingRepository) start(ctx context.Context) (*hazelcast.Map, error) {
	return r.client.GetMap(ctx, r.cfg.MapName)
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
