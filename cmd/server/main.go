package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/gsk967/userActivity/pb"
	"github.com/gsk967/userActivity/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
)

func runGrpcServer(userService pb.UsersServer, listener net.Listener) error {
	grpcServer := grpc.NewServer()

	pb.RegisterUsersServer(grpcServer, userService)
	reflection.Register(grpcServer)

	log.Printf("Server is running in grpc mode on address : %s", listener.Addr().String())

	err := grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("Error while serving grpc with http listen ... %w", err)
	}
	return nil
}

func runRestServer(userService pb.UsersServer, listener net.Listener) error {
	mux := runtime.NewServeMux()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := pb.RegisterUsersHandlerServer(ctx, mux, userService)
	if err != nil {
		return err
	}
	log.Printf("Server is running in rest mode on address : %s", listener.Addr().String())

	err = http.Serve(listener, mux)
	if err != nil {
		log.Fatal("Error while serving grpc with http listen ... %w", err)
	}
	return nil
}

func main() {
	port := flag.Int("port", 8080, "the server port")
	serverType := flag.String("type", "grpc", "type of server (grpc/rest)")
	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", *port))
	if err != nil {
		log.Fatal("Error while server creating ... %w", err)
	}

	userService := service.NewUserService(service.NewInMemoryUserStore(), service.InMemoryUserActivityStore())

	if *serverType == "grpc" {
		err = runGrpcServer(userService, listener)
		if err != nil {
			log.Fatal("Error while running grpc server... %w", err)
		}
	} else {
		err = runRestServer(userService, listener)
		if err != nil {
			log.Fatal("Error while running grpc server... %w", err)
		}
	}
}
