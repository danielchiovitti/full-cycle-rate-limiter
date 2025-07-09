package check_limit

type CheckLimitUseCaseInterface interface {
	Execute(ip, token string) (bool, error)
}
