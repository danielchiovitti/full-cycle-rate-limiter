package check_limit_usecase_factory

import "rate-limiter/pkg/domain/usecase/check_limit"

func NewCheckLimitUseCaseFactory() CheckLimitUseCaseFactoryInterface {
	return &CheckLimitUseCaseFactory{}
}

type CheckLimitUseCaseFactory struct {
}

func (c *CheckLimitUseCaseFactory) Build() check_limit.CheckLimitUseCaseInterface {
	return &check_limit.CheckLimitUseCase{}
}
