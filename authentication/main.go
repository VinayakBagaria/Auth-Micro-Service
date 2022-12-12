package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/VinayakBagaria/auth-micro-service/authentication/repository"
	"github.com/VinayakBagaria/auth-micro-service/authentication/service"
	"github.com/VinayakBagaria/auth-micro-service/db"
	"github.com/VinayakBagaria/auth-micro-service/pb"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

var (
	local bool
	port  int
)

func init() {
	flag.BoolVar(&local, "local", true, "run service local")
	flag.IntVar(&port, "port", 9001, "run authentication service port")
	flag.Parse()
}

func main() {
	if local {
		err := godotenv.Load()
		if err != nil {
			log.Panicln(err)
		}
	}

	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	if err != nil {
		log.Panicln(err)
	}
	defer conn.Close()

	usersRepository := repository.NewUsersRepository(conn)
	authService := service.NewAuthService(usersRepository)

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, authService)
	log.Printf("Authentication service running on [::]:%d\n", port)
	grpcServer.Serve(lis)
}
