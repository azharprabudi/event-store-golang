package model

// ConfPropertyCollection ...
type ConfPropertyCollection struct {
	DefaultMaxLimitAction int `bson:"defaultMaxLimitAction"`
	CurrentMaxLimitAction int `bson:"currentMaxLimitAction"`
}
