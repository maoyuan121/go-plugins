# Wrappers

Wrappers 是中间件的一种形式，它可以和 go-micro 服务一起使用。它可以包装 Client 和 Server handler。

## Client Interface

```go
// Wrapper 包装一个 Client 并返回一个 Client
type Wrapper func(Client) Client

// StreamWrapper 包装一个 Stream 并返回它 equivalent
type StreamWrapper func(Streamer) Streamer
```

## Handler Interface

```go
// HandlerFunc 表示 handler 的一个方法。
// 它主要用于包装。
// 交付给实际方法的是具体的请求和响应类型。
type HandlerFunc func(ctx context.Context, req Request, rsp interface{}) error

// SubscriberFunc 标识 subscriber 的一个方法。
// 它主要用于包装。
// What's handed to the actual method is the concrete publication message.
type SubscriberFunc func(ctx context.Context, msg Event) error

// HandlerWrapper 包装一个 HandlerFunc 并返回它
type HandlerWrapper func(HandlerFunc) HandlerFunc

// SubscriberWrapper 包装一个 SubscriberFunc 并返回它
type SubscriberWrapper func(SubscriberFunc) SubscriberFunc

// StreamerWrapper 包装一个 Streamer 接口 并返回它。
// Because streams exist for the lifetime of a method invocation this
// is a convenient way to wrap a Stream as its in use for trace, monitoring,
// metrics, etc.
type StreamerWrapper func(Streamer) Streamer
```

## Client Wrapper Usage

下面是一个基础的 client 的日志包装器

```go
type logWrapper struct {
	client.Client
}

func (l *logWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	md, _ := metadata.FromContext(ctx)
	fmt.Printf("[Log Wrapper] ctx: %v service: %s method: %s\n", md, req.Service(), req.Endpoint())
	return l.Client.Call(ctx, req, rsp)
}

func NewLogWrapper(c client.Client) client.Client {
	return &logWrapper{c}
}
```


## Handler Wrapper Usage

这个是一个 handler 的基本日志包装器

```go
func NewLogWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		log.Printf("[Log Wrapper] Before serving request method: %v", req.Endpoint())
		err := fn(ctx, req, rsp)
		log.Printf("[Log Wrapper] After serving request")
		return err
	}
}
```
