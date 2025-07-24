package adapters

import (
	"awesomeProject2/internal/core/usecases/create"
	"awesomeProject2/internal/core/usecases/get"
	"awesomeProject2/proto"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GRPCServer struct {
	*grpc.Server
	proto.UnimplementedProfileServiceServer
	createUC *create.UseCase
	getUC    *get.UseCase
}

func NewMyGrpcServer(useCase *create.UseCase, getUseCase *get.UseCase) *GRPCServer {
	server := grpc.NewServer()

	grpcServer := &GRPCServer{
		createUC: useCase,
		getUC:    getUseCase,
		Server:   server,
	}

	proto.RegisterProfileServiceServer(server, grpcServer)
	return grpcServer
}

func (s *GRPCServer) Run() {
	lis, err := net.Listen("tcp", ":50001")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Println("gRPC server listening on :50051")
	if err := s.Server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *GRPCServer) CreateProfile(ctx context.Context, req *proto.CreateRequest) (*proto.CreateResponse, error) {
	log.Printf("Received request: id=%s, email=%s", req.Id, req.Email)
	command := create.Command{ID: req.Id, Email: req.Email}
	err := s.createUC.Handle(ctx, command)
	if err != nil {
		return nil, err
	}
	response := &proto.CreateResponse{Id: req.Id}
	return response, nil
}

func (s *GRPCServer) GetProfile(ctx context.Context, req *proto.GetRequest) (*proto.GetResponse, error) {
	log.Printf("Received request: id=%s", req.Id)
	command := get.Command{ID: req.Id}

	profile, err := s.getUC.Handle(ctx, command)
	if err != nil {
		return nil, err
	}
	response := &proto.GetResponse{Id: profile.ID, Email: profile.Email}
	return response, nil
}
