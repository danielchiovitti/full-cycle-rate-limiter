package get_value

type GetValueUseCaseInterface interface {
	Execute(keyType, key string) string
}
