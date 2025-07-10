package cache

import (
	"context"
	"rate-limiter/pkg/domain/model"
)

type CacheInterface interface {
	SetValue(ctx context.Context, keyType model.ConstraintType, key string, value interface{}, ttl int) error
	IncrValue(ctx context.Context, keyType model.ConstraintType, key string) (int64, error)
	GetValue(ctx context.Context, keyType model.ConstraintType, key string) (interface{}, error)
}
