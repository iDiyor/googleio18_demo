package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"net/http"
	"time"

	pb "googleio18_demo/proto"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port     = ":2020"
	httpPort = ":4040"
)

func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}

	defer lis.Close()

	go runGRPC(lis)
	runHTTP()
}

func runGRPC(lis net.Listener) {
	server := NewServer()

	grpcServer := grpc.NewServer()
	pb.RegisterMapsServer(grpcServer, server)

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}

func runHTTP() {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	addr := fmt.Sprintf("localhost%s", port)

	pb.RegisterMapsHandlerFromEndpoint(context.Background(), mux, addr, opts)

	log.Printf("gRPC server listening on port: %s", port)

	log.Printf("HTTP Listening on %s\n", httpPort)

	http.ListenAndServe(httpPort, mux)
}

//#########################
// Server implementation
//#########################

// Server ...
type Server struct {
	datastore map[string]interface{}
}

// NewServer ...
func NewServer() *Server {
	srv := &Server{
		datastore: make(map[string]interface{}),
	}

	return srv
}

// GetDistance ...
func (s *Server) GetDistance(ctx context.Context, req *pb.GetDistanceRequest) (*pb.GetDistanceResponse, error) {
	log.Println("------------------------")
	log.Println("Getting distance between:")
	log.Printf(" > origin: %s", req.GetOrigin())
	log.Printf(" > dest: %s", req.GetDest())

	time.Sleep(1 * time.Second)

	response := &pb.GetDistanceResponse{
		Distance: randNumber(100, 500),
	}

	return response, nil
}

// StreamLocation ...
func (s *Server) StreamLocation(stream pb.Maps_StreamLocationServer) error {
	for {
		location, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.EmptyResponse{})
		}
		if err != nil {
			return err
		}

		log.Println("------------------------")
		log.Println("Received a location")
		log.Printf(" > latitude: %.1f", location.GetLat())
		log.Printf(" > longitude: %.1f", location.GetLon())
	}
}

// GetVenues ...
func (s *Server) GetVenues(ctx context.Context, req *pb.GetVenuesRequest) (*pb.GetVenuesResponse, error) {
	log.Println("------------------------")
	log.Println("Finding venues....")
	log.Printf(" > location: %s", req.GetLocation())
	log.Printf(" > radius: %d M", req.GetRadius())

	venues := getDummyVenues()

	response := &pb.GetVenuesResponse{
		Venues: venues,
	}

	return response, nil
}

//#########################
// Helper methods
//#########################

func getDummyVenues() []*pb.Venue {
	venueA := &pb.Venue{
		Id:      "1",
		Name:    "Venue-A",
		Address: "Tashkent",
		Rating:  4,
	}

	venueB := &pb.Venue{
		Id:      "2",
		Name:    "Venue-B",
		Address: "London",
		Rating:  3,
	}

	venueC := &pb.Venue{
		Id:      "3",
		Name:    "Venue-C",
		Address: "New-York",
		Rating:  2,
	}

	return []*pb.Venue{
		venueA,
		venueB,
		venueC,
	}
}

func randNumber(min int, max int) int64 {
	rand.Seed(time.Now().Unix())
	return int64(rand.Intn(max-min) + min)
}
