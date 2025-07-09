package check_limit

func NewCheckLimitUseCase() *CheckLimitUseCase {
	return &CheckLimitUseCase{}
}

type CheckLimitUseCase struct {
}

func (c *CheckLimitUseCase) Execute(ip, token string) (bool, error) {
	return true, nil
}
