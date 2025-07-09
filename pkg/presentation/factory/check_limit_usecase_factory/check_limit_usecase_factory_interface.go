package check_limit_usecase_factory

import "rate-limiter/pkg/domain/usecase/check_limit"

type CheckLimitUseCaseFactoryInterface interface {
	Build() check_limit.CheckLimitUseCaseInterface
}
