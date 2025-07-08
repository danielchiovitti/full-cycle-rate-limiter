package cache

type CacheInterface interface {
	SetValue(key, value string, ttl int) error
	GetValue(key string) (string, error)
}
