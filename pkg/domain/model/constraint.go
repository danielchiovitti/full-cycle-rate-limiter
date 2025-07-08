package model

type ConstraintType string

const (
	I ConstraintType = "IP"
	T ConstraintType = "TOKEN"
)

type Constraint struct {
	Key       string
	KeyType   ConstraintType
	Requests  int
	BlockTime int
}
