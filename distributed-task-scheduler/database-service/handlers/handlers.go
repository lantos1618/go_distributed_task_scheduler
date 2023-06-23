package handlers

import (
	"database/sql"

	"google.golang.org/grpc"
)

// Server represents the gRPC server
type Server struct {
	Database *sql.DB
}

// RegisterTaskServiceServer registers the task service server
func RegisterTaskServiceServer(s *grpc.Server, srv proto.TaskServiceServer) {
	proto.RegisterTaskServiceServer(s, srv)
}
