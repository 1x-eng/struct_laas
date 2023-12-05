# Structured Logging as a Service

a gRPC logging server in Go, designed to centralize logs (wrt their structure) from various services, potentially written in different languages. The server is based on the gRPC framework and uses Protocol Buffers for data serialization.

## Prerequisites

- Go (version 1.16 or higher)
- Protocol Buffer Compiler (protoc)
- gRPC Go plugins for Protocol Buffers:

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## Getting started
1. Clone
```sh
git clone https://github.com/1x-eng/struct_laas.git
cd struct_laas
```

2. (Optional) If you made any changes to proto file (the contract), regenerate/refresh protobuf code
```sh
protoc --go_out=. --go-grpc_out=. --proto_path=proto proto/logger.proto
```

3. Run the server
```go
go run server/main.go
```

## Testing
1. You could use a client of your choice - pick your poison, but, you'll need to write some code to connect to the gRPC server. (hint: `grpc.Dial(<server>)` if using golang)

2. If you want something quick and dirty, `gRPCurl` could be an option. 

```sh
grpcurl -plaintext -d '{
    "level": "INFO",
    "message": "Hello from grpcurl",
    "service": "your superb service/microservice/faas/whatever else makes you happy",
    "traceId": "1234",
    "timestamp": "'$(date --iso-8601=seconds)'",
    "additionalData": {
        "key1": "additional value",
        "key2": "more additional value, why not"
    }
}' localhost:50052 logger.LoggerService/Log
```

- If you are on OSX, timestamp might need to be - `"timestamp": "'$(date -u +"%Y-%m-%dT%H:%M:%SZ")'",`



