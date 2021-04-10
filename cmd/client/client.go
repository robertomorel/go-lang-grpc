package main

import (
	"context"
	"fmt"
	"log"

	"github.com/robertomorel/go-lang-grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	// Criando conexão com o servidor grpc
	// Certificado de segurança: grpc.WithInsecure(). Pode ser usado para ambiente de testes
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to gRPC Server: %v", err)
	}

	// Qndo o connection para de ser utilizado, o "defer" se encarrega de fechá-la
	defer connection.Close()

	client := pb.NewUserServiceClient(connection)

	AddUser(client)
}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Morel",
		Email: "r.morel@mail.com",
	}

	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make the gRPC request: %v", err)
	}

	fmt.Println(res)
}
