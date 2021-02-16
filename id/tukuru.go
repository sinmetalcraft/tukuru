package id

import "context"

// Factory is ID生成機として利用できる interface
type Factory interface {
	New() string
}

// Tukuru is IDを生成するFactory
type Tukuru struct {
	factory Factory
}

// NewTukuru is IDを作るFactoryを作る
// DefaultはUUIDを作るFactoryがセットされている
func NewTukuru(ctx context.Context, ops ...TukuruOptions) *Tukuru {
	opt := tukuruOptions{}
	for _, o := range ops {
		o(&opt)
	}

	idTukuru := &Tukuru{
		factory: &UUIDTukuru{},
	}
	if opt.factory != nil {
		idTukuru.factory = opt.factory
	}

	return idTukuru
}

// New is IDを生成して返す
func (tukuru *Tukuru) New() string {
	return tukuru.factory.New()
}

type tukuruOptions struct {
	factory Factory
}

// TukuruOptions is IDTukuru生成時に利用する options
type TukuruOptions func(*tukuruOptions)

// WithFactory is ID生成機を設定する
func WithFactory(factory Factory) TukuruOptions {
	return func(ops *tukuruOptions) {
		ops.factory = factory
	}
}
