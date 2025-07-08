package model

type ConstraintType string

const (
	I ConstraintType = "IP"
	T ConstraintType = "TOKEN"
)

type Constraint struct {
	Key       string
	Type      ConstraintType
	Requests  int
	BlockTime int
}
