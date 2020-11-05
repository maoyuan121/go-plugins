package nats

import "crypto/tls"

// 我们用 Options 配置 nats stream
type Options struct {
	ClusterID string
	ClientID  string
	Address   string
	TLSConfig *tls.Config
}

// Option 是一个函数 用来配置 options
type Option func(o *Options)

// ClusterID sets the cluster id for the nats connection
func ClusterID(id string) Option {
	return func(o *Options) {
		o.ClusterID = id
	}
}

// ClientID sets the client id for the nats connection
func ClientID(id string) Option {
	return func(o *Options) {
		o.ClientID = id
	}
}

// Address of the nats cluster
func Address(addr string) Option {
	return func(o *Options) {
		o.Address = addr
	}
}

// TLSConfig to use when connecting to the cluster
func TLSConfig(t *tls.Config) Option {
	return func(o *Options) {
		o.TLSConfig = t
	}
}
