package tukuru

import "context"

// IDFactory is ID生成機として利用できる interface
type IDFactory interface {
	New() string
}

// ID is IDを生成するFactory
type ID struct {
	idFactory IDFactory
}

// NewIDTukuru is IDを作るFactoryを作る
// DefaultはUUIDを作るFactoryがセットされている
func NewIDTukuru(ctx context.Context, ops ...IDTukuruOptions) *ID {
	opt := idTukuruOptions{}
	for _, o := range ops {
		o(&opt)
	}

	idTukuru := &ID{
		idFactory: &UUID{},
	}
	if opt.idFactory != nil {
		idTukuru.idFactory = opt.idFactory
	}

	return idTukuru
}

// New is IDを生成して返す
func (idTukuru *ID) New() string {
	return idTukuru.idFactory.New()
}

type idTukuruOptions struct {
	idFactory IDFactory
}

// IDTukuruOptions is IDTukuru生成時に利用する options
type IDTukuruOptions func(*idTukuruOptions)

// WithIDFactory is
func WithIDFactory(factory IDFactory) IDTukuruOptions {
	return func(ops *idTukuruOptions) {
		ops.idFactory = factory
	}
}
