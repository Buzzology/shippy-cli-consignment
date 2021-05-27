package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/Buzzology/shippy-service-consignment/proto/consignment"
	pbVessel "github.com/Buzzology/shippy-service-vessel/proto/vessel"
	micro "github.com/micro/go-micro/v2"
)

const (
	address         = "localhost:50051"
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, err
}

func main() {

	service := micro.NewService(micro.Name("shippy.consignment.cli"))
	service.Init()

	client := pb.NewShippingService("shippy.service.consignment", service.Client())
	vesselClient := pbVessel.NewVesselService("shippy.service.vessel", service.Client())

	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)
	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	// Ensure that expected vessels are able to be retrieved from the cli
	vesselQueriesToTest := []*pbVessel.Specification{
		{Capacity: 1, MaxWeight: 1},
		{Capacity: int32(len(consignment.Containers)), MaxWeight: consignment.Weight},
	}

	for _, v := range vesselQueriesToTest {

		res, err := vesselClient.FindAvailable(context.Background(), v)

		if err != nil {
			log.Fatal(err)
		}

		if res == nil {
			log.Fatal("Expected to retrieve a vessel.")
		}
	}

	r, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}

	log.Printf("Created: %t", r.Created)

	// Retrieve all consignments
	getAll, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}

	// Display each of the retrieved consignments
	for _, v := range getAll.Consignments {
		log.Println(v)
	}
}
