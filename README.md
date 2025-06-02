# TaskSync
TaskSync is a client-server application where clients can create, update, delete, and fetch to-do tasks via gRPC. The project teaches how to define gRPC services in .proto files, implement server logic in Go, connect with clients, and handle typical features such as status codes, metadata, and streaming.



# Personal notes
go mod init tasksync
go mod tidy
protoc-gen-go -> Converts .proto into basic Go types
protoc-gen-go-grpc -> Converts .proto into gRPC service interfaces for Go
protoc --go_out=. --go-grpc_out=. proto/task.proto -> Generates code from .proto file and put it into the proto directory **rerun if changes to rpc calls and services
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    helloworld/helloworld.proto ** something like this is for using the relative path?
**use go build over go run, the executable works better in this case than the temporary in storage build of go run (NTS --> figure out why)

$1, $2 -> placeholders to prevent sql injection, follows order defined in query execution args

** DB testing
don't forget to launch docker first
docker-compose down
docker-compose down -v ->full reset
docker-compose up --build