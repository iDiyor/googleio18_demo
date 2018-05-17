package main

import (
	"log"
	"time"

	"golang.org/x/net/context"

	pb "googleio18_demo/proto"

	"google.golang.org/grpc"
)

const (
	mapsServiceAddress = "localhost:2020"
)

func main() {
	conn, err := grpc.Dial(mapsServiceAddress, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(5*time.Second))
	if err != nil {
		panic(err)
	}

	mapsClient := pb.NewMapsClient(conn)

	getDistance(mapsClient)
}

func getDistance(client pb.MapsClient) error {
	origin := "38.8633,65.7978"
	dest := "44.8633,75.7978"

	log.Println("------------------------")
	log.Println("Requesting distance ....")
	log.Printf(" > origin: %s", origin)
	log.Printf(" > dest: %s", dest)

	result, err := client.GetDistance(context.Background(), &pb.GetDistanceRequest{
		Origin: origin,
		Dest:   dest,
	})
	if err != nil {
		return err
	}

	log.Println("------------------------")
	log.Printf("Got distance: %d M", result.GetDistance())
	return nil
}
