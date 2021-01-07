package users

import (
	"github.com/gsk967/userActivity/pb"
	"google.golang.org/grpc"
	"log"
)

func GrpcClient(serverAddress string) (pb.UsersClient, error) {
	log.Printf("dial server %s", serverAddress)
	transportOption := grpc.WithInsecure()

	grpcClient, err := grpc.Dial(serverAddress, transportOption)
	if err != nil {
		log.Fatal("cannot dial server: ", err)
		return nil, err
	}

	userServiceClient := pb.NewUsersClient(grpcClient)
	return userServiceClient, nil
}
