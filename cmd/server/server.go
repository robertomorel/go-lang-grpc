package main

import (
	"log"
	"net"

	"github.com/robertomorel/go-lang-grpc/pb"
	"github.com/robertomorel/go-lang-grpc/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Criando listener para ficar ouvindo em um end. e porta
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Could not conect: %v", err)
	}

	grpcServer := grpc.NewServer()

	// Registrando o serviço "user.go" neste servidor
	pb.RegisterUserServiceServer(grpcServer, services.NewUserService())
	// Rodando no modo reflection para que o client tenha acesso aos métodos do server
	reflection.Register(grpcServer)

	// Servir(subir) o listener
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Could not serve: %v", err)
	}
}
