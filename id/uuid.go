package id

import "github.com/google/uuid"

// UUID is UUID生成機
type UUIDTukuru struct {
}

// New is UUIDを生成して返す
func (uuidTukuru *UUIDTukuru) New() string {
	return uuid.NewString()
}
