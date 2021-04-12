// Run: go run cmd/client/client.go

package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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

	//AddUser(client)
	//AddUserVerbose(client)
	//AddUsers(client)
	AddUserStreamBoth(client)
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

func AddUserVerbose(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Morel",
		Email: "r.morel@mail.com",
	}

	responseStream, err := client.AddUserVerbose(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make the gRPC request: %v", err)
	}

	// Faz um looping enquanto tiver recebendo informações
	for {
		stream, err := responseStream.Recv()
		// Se não tem mais arquivos a receber...
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Could not receive the msg: %v", err)
		}
		fmt.Println("Status: ", stream.Status, " - ", stream.GetUser())
	}
}

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		&pb.User{
			Id:    "0",
			Name:  "Morel 0",
			Email: "r.morel0@mail.com",
		},
		&pb.User{
			Id:    "2",
			Name:  "Morel 2",
			Email: "r.morel2@mail.com",
		},
		&pb.User{
			Id:    "4",
			Name:  "Morel 4",
			Email: "r.morel4@mail.com",
		},
		&pb.User{
			Id:    "6",
			Name:  "Morel 6",
			Email: "r.morel6@mail.com",
		},
	}

	// context.Background() -> garante que se a mensagem não for chegar ele já para. Controla o fluxo de dados
	stream, err := client.AddUsers(context.Background())

	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	// Mandar as requisições por streaming dentro de um laço
	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}

	fmt.Println(res)
}

func AddUserStreamBoth(client pb.UserServiceClient) {
	// context.Background() -> garante que se a mensagem não for chegar ele já para. Controla o fluxo de dados
	stream, err := client.AddUserStreamBoth(context.Background())

	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	reqs := []*pb.User{
		&pb.User{
			Id:    "0",
			Name:  "Morel 0",
			Email: "r.morel0@mail.com",
		},
		&pb.User{
			Id:    "2",
			Name:  "Morel 2",
			Email: "r.morel2@mail.com",
		},
		&pb.User{
			Id:    "4",
			Name:  "Morel 4",
			Email: "r.morel4@mail.com",
		},
		&pb.User{
			Id:    "6",
			Name:  "Morel 6",
			Email: "r.morel6@mail.com",
		},
	}

	// Criando uma variável "wait" do tipo channel
	// Channel é um local onde você manda uma comunicação entre duas Go Routines
	wait := make(chan int)

	//GO Routing - "Thread" controlada pelo Golang, que cria e controla milhões de threads do Go
	// O sistema pode ficar esperando para sempre enviando e recebendo informação
	go func() {
		// Percorrendo e enviando requisições
		for _, req := range reqs {
			fmt.Println("Sending user: ", req.GetName())
			stream.Send(req)
			time.Sleep(time.Second * 2)
		}
		//Quanto terminar de enviar tudo, fecha o envio
		stream.CloseSend()
	}()

	go func() {
		for {
			// Recebendo o streaming de dados da thread anterior
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error receiving data: %v", err)
				break
			}
			fmt.Printf("Receiving user %v with status %v\n", res.GetUser().GetName(), res.GetStatus())
		}
		// Forçar o channel a morrer qndo o servidor parar de mandar informação
		close(wait)
	}()

	// Enquanto esse channel não morrer, a aplicação não morre, ou seja, ficaria em um looping infinito
	// Segura o processo rodando enquando as informações do streaming ainda estão chegando
	<-wait
}
