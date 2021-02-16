package now

import "time"

// NowTukuru is time.Now() 生成機
type NowTukuru struct {
}

// New is time.Now() を実行して返す
func (nowTukuru *NowTukuru) New() time.Time {
	return time.Now()
}
