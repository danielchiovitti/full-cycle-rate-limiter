package get_value

import (
	"rate-limiter/pkg/infrastructure/cache"
	"sync"
)

var getValueUseCaseLock sync.Mutex
var getValueUseCaseInstance *GetValueUseCase

func NewGetValueUseCase(
	cache cache.CacheInterface,
) *GetValueUseCase {
	if getValueUseCaseInstance == nil {
		getValueUseCaseLock.Lock()
		defer getValueUseCaseLock.Unlock()
		if getValueUseCaseInstance == nil {
			getValueUseCaseInstance = &GetValueUseCase{
				cache: cache,
			}
		}
	}
	return getValueUseCaseInstance
}

type GetValueUseCase struct {
	cache cache.CacheInterface
}
