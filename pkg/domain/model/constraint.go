package model

type ConstraintType string

const (
	CONSTRAINT_I             ConstraintType = "IP"
	CONSTRAINT_T             ConstraintType = "TOKEN"
	CONSTRAINT_LI            ConstraintType = "LOADED_IP"
	CONSTRAINT_LT            ConstraintType = "LOADED_TOKEN"
	CONSTRAINT_BLOCKED_IP    ConstraintType = "BLOCKED_IP"
	CONSTRAINT_BLOCKED_TOKEN ConstraintType = "BLOCKED_TOKEN"
)

type Constraint struct {
	Key       string         `json:"key"`
	KeyType   ConstraintType `json:"keyType"`
	Requests  int64          `json:"requests"`
	BlockTime int64          `json:"blockTime"`
}
