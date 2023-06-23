package main

import (
	"log"
	"net"

	"github.com/lib/pq"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

func main() {
	// Set up a connection to the PostgreSQL server.
	conn, err := pq.NewConnector("user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Register our service with gRPC.
	s := grpc.NewServer()
	handlers.RegisterTaskServiceServer(s, &handlers.Server{Database: conn})

	// Listen on the defined port.
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Start the server.
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
