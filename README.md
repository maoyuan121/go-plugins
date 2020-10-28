# Plugins [![License](https://img.shields.io/:license-apache-blue.svg)](https://opensource.org/licenses/Apache-2.0) [![GoDoc](https://godoc.org/github.com/micro/go-plugins?status.svg)](https://godoc.org/github.com/micro/go-plugins)

Go plugins 是 Go Micro 接口的第三方实现的地方。

## Overview
Go Micro 通过使用 Go 接口构架为可插拔的框架。插件使您能够换出底层基础设施，而不需要重写你所有的代码。这允许在多个环境中运行相同的软件，而无需进行大量工作。进一步阅读了解更多信息。
                                                         
## Getting Started

* [Contents](#contents)
* [Usage](#usage)
* [Build](#build)

## Contents

Contents of this repository:

| Directory | Description                                                     |
| --------- | ----------------------------------------------------------------|
| Broker    | PubSub messaging; NATS, NSQ, RabbitMQ, Kafka                    |
| Client    | RPC Clients; gRPC, HTTP                                         |
| Codec     | Message Encoding; BSON, Mercury                                 |
| Micro     | Micro Toolkit Plugins                                           |
| Registry  | 服务发现； Etcd, Gossip, NATS                           |
| Selector  | 负载均衡； Label, Cache, Static                            |
| Server    | RPC Servers; gRPC, HTTP                                         |
| Transport | Bidirectional Streaming; NATS, RabbitMQ                         | 
| Wrapper   | 中间件; Circuit Breakers, Rate Limiting, Tracing, Monitoring|

## Usage

插件可以通过以下方式添加到 go-micro。这样就可以通过命令行参数或环境变量来设置它们。

在 `plugins.go` 中导入插件

```go
package main

import (
	_ "github.com/micro/go-plugins/broker/rabbitmq/v2"
	_ "github.com/micro/go-plugins/registry/kubernetes/v2"
	_ "github.com/micro/go-plugins/transport/nats/v2"
)
```

创建你的服务，确保调用了 `service.Init`

```go
package main

import (
	"github.com/micro/go-micro/v2"
)

func main() {
	service := micro.NewService(
		// 设置服务名
		micro.Name("my.service"),
	)

	// 解析 CLI flags
	service.Init()
}
```

编译你的服务

```
go build -o service ./main.go ./plugins.go
```

### Environment Variables

使用环境变量来设置

```
MICRO_BROKER=rabbitmq \
MICRO_REGISTRY=kubernetes \ 
MICRO_TRANSPORT=nats \ 
./service
```

### Flags

或者使用命令行 flags 来设置

```shell
./service --broker=rabbitmq --registry=kubernetes --transport=nats
```

### Options

导入，在创建一个新的服务时通过 options 设置

```go
import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/registry/kubernetes/v2"
)

func main() {
	registry := kubernetes.NewRegistry() // a default to using env vars for master API

	service := micro.NewService(
		// 设置服务名
		micro.Name("my.service"),
		// 设置服务 registry
		micro.Registry(registry),
	)
}
```

## Build

反模式是修改 `main.go` 文件来包含插件。最佳实践建议是 include 插件在一个单独的文件和重建与它包括。这允许自动化构建插件和关注点的干净分离。


创建 plugins.go 文件

```go
package main

import (
	_ "github.com/micro/go-plugins/broker/rabbitmq/v2"
	_ "github.com/micro/go-plugins/registry/kubernetes/v2"
	_ "github.com/micro/go-plugins/transport/nats/v2"
)
```

结合 plugins.go 来编译

```shell
go build -o service main.go plugins.go
```

Run with plugins

```shell
MICRO_BROKER=rabbitmq \
MICRO_REGISTRY=kubernetes \
MICRO_TRANSPORT=nats \
service
```
