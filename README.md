# TaskSync
TaskSync is a client-server application where clients can create, update, delete, and fetch to-do tasks via gRPC. The project teaches how to define gRPC services in .proto files, implement server logic in Go, connect with clients, and handle typical features such as status codes, metadata, and streaming.



# Personal notes
protoc-gen-go -> Converts .proto into basic Go types
protoc-gen-go-grpc -> Converts .proto into gRPC service interfaces for Go
protoc --go_out=. --go-grpc_out=. proto/task.proto -> Generates code from .proto file and put it into the proto directory
