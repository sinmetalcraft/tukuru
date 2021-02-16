package now

import (
	"context"
	"time"
)

// Factory is Now生成機として利用できるinterface
type Factory interface {
	New() time.Time
}

// Tukuru is Nowを生成するFactory
type Tukuru struct {
	factory Factory
}

// NewTukuru is Nowを作るFactoryを作る
// Defaultはtime.Now()を返すFactoryがセットされている
func NewTukuru(ctx context.Context, ops ...TukuruOptions) *Tukuru {
	opt := tukuruOptions{}
	for _, o := range ops {
		o(&opt)
	}
	nowTukuru := &Tukuru{
		factory: &NowTukuru{},
	}
	if opt.factory != nil {
		nowTukuru.factory = opt.factory
	}
	return nowTukuru
}

// New is Nowを生成して返す
func (tukuru *Tukuru) New() time.Time {
	return tukuru.factory.New()
}

type tukuruOptions struct {
	factory Factory
}

// TukuruOptions is NowTukuru生成時に利用するoptions
type TukuruOptions func(*tukuruOptions)

// WithFactory is Now生成機を設定する
func WithFactory(factory Factory) TukuruOptions {
	return func(ops *tukuruOptions) {
		ops.factory = factory
	}
}
