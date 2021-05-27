module github.com/Buzzology/shippy-cli-consignment

go 1.16

replace google.golang.org/grpc v1.38.0 => google.golang.org/grpc v1.26.0

require (
	github.com/Buzzology/shippy-service-consignment v0.0.3
	github.com/micro/go-micro/v2 v2.9.1
)
