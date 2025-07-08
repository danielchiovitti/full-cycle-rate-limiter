package get_value_usecase_factory

import (
	"rate-limiter/pkg/domain/usecase/get_value"
)

type GetValueUseCaseFactoryInterface interface {
	Build() *get_value.GetValueUseCase
}
