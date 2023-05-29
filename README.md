# 微服务demo

# 安装以下工具

- protoc-gen-grpc-gateway
- protoc-gen-openapiv2
- protoc-gen-go
- protoc-gen-go-grpc
- buf
- wire

```go
// +build tools

package tools

import (
    _ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
    _ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
    _ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
    _ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
```
> protoc-gen-grpc-gateway & protoc-gen-openapiv2 & protoc-gen-go & protoc-gen-go-grpc
```shell
$ go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
$ go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

> Windows下
```shell
$ go install ^
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway ^
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 ^
    google.golang.org/protobuf/cmd/protoc-gen-go ^
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

> wire
```shell
$ go install github.com/google/wire/cmd/wire@latest
```

> buf Windows
```shell
https://github.com/bufbuild/buf/releases/latest
```