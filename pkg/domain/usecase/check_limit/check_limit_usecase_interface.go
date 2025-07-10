package check_limit

import "context"

type CheckLimitUseCaseInterface interface {
	Execute(ctx context.Context, ip, token string) (bool, error)
}
