package main

import (
	"golang.org/x/net/context"
	pb "goTest/grpc/proto"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	//lis, err := net.Listen("tcp", port)
	//if err != nil {
	//	log.Fatalf("failed to listen: %v", err)
	//}
	//s := grpc.NewServer()
	//pb.RegisterGreeterServer(s, &server{})
	//// Register reflection service on gRPC server.
	//reflection.Register(s)
	//if err := s.Serve(lis); err != nil {
	//	log.Fatalf("failed to serve: %v", err)
	//}

	//str := strconv.FormatInt(10, 10)
	//var _startDate int64 = time.Now().Unix()
	//var startDate string = time.Unix(10, 0).Format("15:04:05")

	//print(startDate)

	var arr []int
	arr = append(arr, 11)
}