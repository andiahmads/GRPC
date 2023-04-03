# Learning-GRPC

Simple Implementation GRPC for baginer

# Features

- Unary GRPC
- Server Streaming
- Client Streaming
- Bidirectional GRPC
- GRPC Deadline

# Setup Local Machine 
for generate proto file, please visit

```
manual instalation
https://grpc.io/docs/languages/go/quickstart/


for ubuntu user
https://snapcraft.io/protobuf


for mac user
brew install protobuf
```

So make sure install the dependency in your local:

```go
  go mod tidy
```

If you have protoc compailer and you want generate proto file, please running:

```go
  make hello-grpc
```
running server
```go
  make hello-server
```

running client
```go
  make hello-server
```

download evans-cli for GRPC testing
```
https://github.com/ktr0731/evans
```

running evan-cli
```go
  make grpc-test
```
