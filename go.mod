module github.com/i-coder-robot/go-micro-action-user-api

go 1.14

require (
	github.com/coreos/etcd v3.3.22+incompatible // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/go-log/log v0.1.0
	github.com/go-redis/redis v6.15.7+incompatible
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.4.2
	github.com/i-coder-robot/go-micro-action-core v0.0.7
	github.com/i-coder-robot/go-micro-action-user v0.0.5
	github.com/micro/go-micro/v2 v2.9.1
)

replace google.golang.org/grpc v1.27.0 => google.golang.org/grpc v1.26.0
