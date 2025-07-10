package check_limit

import (
	"context"
	"encoding/json"
	"rate-limiter/pkg/domain/model"
	"rate-limiter/pkg/infrastructure/cache"
	"rate-limiter/pkg/shared"
)

func NewCheckLimitUseCase(
	cache cache.CacheInterface,
	config shared.ConfigInterface,
) *CheckLimitUseCase {
	return &CheckLimitUseCase{
		cache:  cache,
		config: config,
	}
}

type CheckLimitUseCase struct {
	cache  cache.CacheInterface
	config shared.ConfigInterface
}

func (c *CheckLimitUseCase) Execute(ctx context.Context, ip, token string) (bool, error) {
	var valIp int64
	var valToken int64
	var cIp *model.Constraint
	var cToken *model.Constraint

	if ip != "" {
		j, err := c.cache.GetValue(ctx, model.CONSTRAINT_BLOCKED_IP, ip)
		if err != nil {
			return false, err
		}

		// ip is blocked
		if j != nil {
			return false, nil
		}

		j, err = c.cache.GetValue(ctx, model.CONSTRAINT_LI, ip)
		if err != nil {
			return false, err
		}

		if j != nil {
			cIp = &model.Constraint{}
			err = json.Unmarshal([]byte(j.(string)), cIp)
			if err != nil {
				return false, err
			}
		}

		valIp, err = c.cache.IncrValue(ctx, model.CONSTRAINT_I, ip)
		if err != nil {
			return false, err
		}
	}

	if token != "" {
		j, err := c.cache.GetValue(ctx, model.CONSTRAINT_BLOCKED_TOKEN, token)
		if err != nil {
			return false, err
		}

		// token is blocked
		if j != nil {
			return false, nil
		}

		j, err = c.cache.GetValue(ctx, model.CONSTRAINT_LT, token)
		if err != nil {
			return false, err
		}

		if j != nil {
			cToken = &model.Constraint{}
			err = json.Unmarshal([]byte(j.(string)), cToken)
			if err != nil {
				return false, err
			}
		}

		valToken, err = c.cache.IncrValue(ctx, model.CONSTRAINT_T, token)
		if err != nil {
			return false, err
		}
	}

	maxRequests := int64(0)
	blockTime := int64(0)
	if cIp != nil && cToken != nil {
		maxRequests = cToken.Requests
		blockTime = cToken.BlockTime

		if valToken > maxRequests {
			err := c.cache.SetValue(ctx, model.CONSTRAINT_BLOCKED_TOKEN, token, token, int(blockTime))
			return false, err
		}
	} else if cIp != nil {
		maxRequests = cIp.Requests
		blockTime = cIp.BlockTime

		if valIp > cIp.Requests {
			err := c.cache.SetValue(ctx, model.CONSTRAINT_BLOCKED_IP, ip, ip, int(blockTime))
			return false, err
		}
	} else if cToken != nil {
		maxRequests = cToken.Requests
		blockTime = cToken.BlockTime

		if valToken > maxRequests {
			err := c.cache.SetValue(ctx, model.CONSTRAINT_BLOCKED_TOKEN, token, token, int(blockTime))
			return false, err
		}
	} else {
		maxRequests = int64(c.config.GetMaxRequests())
		blockTime = int64(c.config.GetBlockTime())

		if ip != "" {
			if valIp > maxRequests {
				err := c.cache.SetValue(ctx, model.CONSTRAINT_BLOCKED_IP, ip, ip, int(blockTime))
				return false, err
			}
		}

		if token != "" {
			if valToken > maxRequests {
				err := c.cache.SetValue(ctx, model.CONSTRAINT_BLOCKED_TOKEN, token, token, int(blockTime))
				return false, err
			}
		}
	}

	return true, nil
}
