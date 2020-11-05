# Endpoint Wrapper

Endpint 包装器是一种函数，它允许您在更细粒度的级别上执行包装器。
此时客户端或处理程序包装器在任何请求方法上执行。
端点包装器可以更容易地指定要执行的精确方法，否则将充当传递。



## Usage

创建服务，如下一样添加包装器。

```go
srv := micro.NewService(
	micro.Name("com.example.srv.foo"),
)

srv.Init(
	// cw is your client wrapper
	// hw is your handler wrapper
	// Foo.Bar and Foo.Baz are the methods to execute on
	micro.WrapClient(endpoint.NewClientWrapper(cw, "Foo.Bar", "Foo.Baz"))
	micro.WrapHandler(endpoint.NewHandlerWrapper(hw, "Foo.Bar", "Bar.Baz", "Debug.Health")),
)

```
