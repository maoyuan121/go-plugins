package nats

import (
	"context"

	"github.com/asim/go-micro/v3/broker"
)

// setBrokerOption 返回一个用来设置 context 键值的函数
func setBrokerOption(k, v interface{}) broker.Option {
	return func(o *broker.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, k, v)
	}
}
