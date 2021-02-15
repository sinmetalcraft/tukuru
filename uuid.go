package tukuru

import "github.com/google/uuid"

// UUID is UUID生成機
type UUID struct {
}

// New is UUIDを生成して返す
func (uuidTukuru *UUID) New() string {
	return uuid.NewString()
}
