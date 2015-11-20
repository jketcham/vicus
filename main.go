package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/jketcham/vicus/Godeps/_workspace/src/golang.org/x/net/context"
	"github.com/jketcham/vicus/Godeps/_workspace/src/google.golang.org/grpc"
	"github.com/jketcham/vicus/Godeps/_workspace/src/gopkg.in/mgo.v2/bson"

	"github.com/jketcham/vicus/model"
	pb "github.com/jketcham/vicus/proto/vicus"
	"github.com/jketcham/vicus/shared/database"
)

type vicusServer struct {
	serverName string
}

func (s *vicusServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	user := new(model.User)

	err := user.FindById(bson.ObjectId(req.Id))
	if err != nil {
		log.Fatalf("Couldn't get user: %s\n", err)
		return &pb.UserResponse{}, err
	}
	fmt.Printf("%v\n", user)

	userResponse := &pb.UserResponse{User: &pb.User{Id: user.Id.String(), FirstName: user.FirstName, LastName: user.LastName, Location: user.Location, Bio: user.Bio}}

	return userResponse, nil
}

func (s *vicusServer) GetUsers(ctx context.Context, req *pb.GetUsersRequest) (*pb.UsersResponse, error) {
	return &pb.UsersResponse{}, nil
}

func (s *vicusServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{}, nil
}

func (s *vicusServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserResponse, error) {
	user := new(model.User)

	err := user.FindById(bson.ObjectId(req.UserId))
	if err != nil {
		log.Fatalf("Couldn't get user: %s\n", err)
		return &pb.UserResponse{}, nil
	}

	err = user.Update(req.Email, req.Password)
	if err != nil {
		log.Fatalf("Couldn't update user: %s\n", err)
		return &pb.UserResponse{}, nil
	}

	return &pb.UserResponse{}, nil
}

func (s *vicusServer) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	user := new(model.User)

	err := user.FindById(bson.ObjectId(req.UserId))
	if err != nil {
		log.Fatalf("Couldn't get user: %s\n", err)
		return &pb.DeleteUserResponse{Status: "failure"}, nil
	}

	err = user.Delete()
	if err != nil {
		log.Fatalf("Couldn't delete user: %s\n", err)
		return &pb.DeleteUserResponse{Status: "failure"}, nil
	}

	return &pb.DeleteUserResponse{Status: "success"}, nil
}

func newServer() *vicusServer {
	s := &vicusServer{serverName: "service.vicus"}
	return s
}

func main() {
	fmt.Printf("starting vicus\n")
	var (
		port     = flag.Int("port", 8080, "The server port")
		mongoURL = flag.String("mongoURL", "localhost/vicus-dev", "URL for mongodb server")
	)
	flag.Parse()

	database.Connect(*mongoURL)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterVicusServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
