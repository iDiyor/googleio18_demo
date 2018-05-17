package main

import (
	"log"
	"math/rand"
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
	streamLocation(mapsClient)
	getVenues(mapsClient)
}

func getDistance(client pb.MapsClient) {
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
		panic(err)
	}

	log.Println("------------------------")
	log.Printf("Got distance: %d M", result.GetDistance())
}

func streamLocation(client pb.MapsClient) {
	stream, err := client.StreamLocation(context.Background())
	if err != nil {
		panic(err)
	}

	for _, loc := range getRandomLocations() {

		// time.Sleep(2 * time.Second)

		if err := stream.Send(loc); err != nil {
			log.Fatalf("%v.Send(%v) = %v", stream, loc, err)
		}
	}

	log.Println("------------------------")
	log.Println("Locations are sent")

	_, err = stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
	}
}

func getVenues(client pb.MapsClient) {
	result, err := client.GetVenues(context.Background(), &pb.GetVenuesRequest{
		Location: "38.8633,65.7978",
		Radius:   500,
	})
	if err != nil {
		panic(err)
	}

	log.Println("------------------------")
	log.Println("Venues")
	for _, venue := range result.GetVenues() {
		log.Println("------------------------")
		log.Printf("ID: %s", venue.GetId())
		log.Printf("Name: %s", venue.GetName())
		log.Printf("Address: %s", venue.GetAddress())
		log.Printf("Rating: %0.1f", venue.GetRating())
	}
}

//#########################
// Helper methods
//#########################
func getRandomLocations() []*pb.Location {
	result := make([]*pb.Location, 1000)

	for i := 0; i < 1000; i++ {
		lat := randNumber(10, 50)
		lon := randNumber(50, 100)
		location := &pb.Location{
			Lat: lat,
			Lon: lon,
		}

		result[i] = location
	}

	return result
}

func randNumber(min int, max int) float64 {
	rand.Seed(time.Now().UnixNano())
	return float64(rand.Intn(max-min) + min)
}
