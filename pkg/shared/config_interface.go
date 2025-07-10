package shared

type ConfigInterface interface {
	GetPort() uint16
	GetMaxRequests() int
	GetBlockTime() int
	GetConstraintList() string
	GetCacheEngine() string
	GetRedisHost() string
	GetRedisPort() int
	GetRedisPassword() string
	GetTimeRange() int
}
