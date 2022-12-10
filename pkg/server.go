package pkg

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"user-management/proto"
)

type Server struct {
	Opt      grpc.ServerOption
	GrpcPort string
}

//func (s *Server) NewServer() error {
//	// Listen to service port
//	listener, err := net.Listen("tcp", ":"+s.GrpcPort)
//	if err != nil {
//		errMsg := fmt.Errorf("failed to listen on port %s: %w", s.GrpcPort, err)
//		return errMsg
//	}
//	defer listener.Close() //nolint:errcheck // no need to err handle here as runtime is already in shut down
//	grpc.NewServer(s.Opt)
//	proto.RegisterUserServiceServer(grpc.NewServer(s.Opt), s)
//	return nil
//}

func (s *Server) GetUser(ctx context.Context, req *user.GetUserRequest) (*user.GetUserResponse, error) {
	log.Print("GetUser called")
	mocked := &user.GetUserResponse{
		Id:   "123",
		Name: "test",
	}
	return mocked, nil
}
