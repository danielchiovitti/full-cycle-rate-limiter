package model

type ConstraintType string

const (
	I ConstraintType = "IP"
	T ConstraintType = "TOKEN"
	P ConstraintType = "PRE_LOADED"
)

type Constraint struct {
	Key       string         `json:"key"`
	KeyType   ConstraintType `json:"keyType"`
	Requests  int            `json:"requests"`
	BlockTime int            `json:"blockTime"`
}
