module github.com/Buzzology/shippy-cli-consignment

go 1.16

replace google.golang.org/grpc v1.38.0 => google.golang.org/grpc v1.26.0

replace github.com/Buzzology/shippy-service-consignment => ../shippy-service-consignment

replace github.com/Buzzology/shippy-service-vessel => ../shippy-service-vessel

require (
	github.com/Buzzology/shippy-service-consignment v0.0.3
	github.com/Buzzology/shippy-service-vessel v0.0.4
	github.com/micro/go-micro/v2 v2.9.1
)
