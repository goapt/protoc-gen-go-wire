# protoc-gen-go-wire

从 protobuf 文件中生成wire依赖注入的变量别名
## 安装

请确保安装了以下依赖:

- [go 1.16](https://golang.org/dl/)
- [protoc](https://github.com/protocolbuffers/protobuf)
- [protoc-gen-go](https://github.com/protocolbuffers/protobuf-go)

```bash
go install github.com/goapt/protoc-gen-go-wire@latest
```

### 文件生成

```bash
  protoc -I ./example/ \
  --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative,require_unimplemented_servers=false \
  --go-wire_out=. --go-wire_opt=paths=source_relative \
  example/test.proto
```
