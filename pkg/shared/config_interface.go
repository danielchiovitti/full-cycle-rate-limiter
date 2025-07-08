package shared

type ConfigInterface interface {
	GetPort() uint16
	GetMaxRequests() int
	GetBlockTime() int
	GetConstraintList() string
}
