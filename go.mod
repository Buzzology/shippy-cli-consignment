module github.com/buzzology/go-microservices-tutorial/shippy-cli-consignment/v0.2

go 1.16

replace google.golang.org/grpc v1.38.0 => google.golang.org/grpc v1.26.0

require (
	github.com/buzzology/go-microservices-tutorial/shippy-service-consignment v0.2
	github.com/micro/go-micro/v2 v2.9.1
)
