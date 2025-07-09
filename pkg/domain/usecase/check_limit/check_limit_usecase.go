package check_limit

import "rate-limiter/pkg/infrastructure/cache"

func NewCheckLimitUseCase(
	cache cache.CacheInterface,
) *CheckLimitUseCase {
	return &CheckLimitUseCase{
		cache: cache,
	}
}

type CheckLimitUseCase struct {
	cache cache.CacheInterface
}

func (c *CheckLimitUseCase) Execute(ip, token string) (bool, error) {
	return true, nil
}
