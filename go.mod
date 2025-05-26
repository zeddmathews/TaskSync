module github.com/zeddmathews/tasksync

go 1.24.3

require (
	github.com/lib/pq v1.10.9
	google.golang.org/grpc v1.72.1
	google.golang.org/protobuf v1.36.6
)

require (
	golang.org/x/net v0.40.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.25.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250512202823-5a2f75b736a9 // indirect
)

replace github.com/zeddmathews/tasksync/proto => ./proto
