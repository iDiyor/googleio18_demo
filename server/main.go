package main

import (
	"io"
	"log"
	"math/rand"
	"net"
	"time"

	pb "googleio18_demo/proto"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":2020"
)

func main() {
	server := NewServer()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}

	defer lis.Close()

	grpcServer := grpc.NewServer()
	pb.RegisterMapsServer(grpcServer, server)

	log.Printf("gRPC server listening on port: %s", port)

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
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

func randNumber(min int, max int) int64 {
	rand.Seed(time.Now().Unix())
	return int64(rand.Intn(max-min) + min)
}
