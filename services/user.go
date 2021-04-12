package services

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/robertomorel/go-lang-grpc/pb"
)

// Precisamos implementar esse type, que está no arquivo "pb/user_grpc.pb.go"
//type UserServiceServer interface {
// Criando um serviço do RPC com uma função AddUser
//	AddUser(context.Context, *User) (*User, error)
//	mustEmbedUnimplementedUserServiceServer()
//  AddUserVerbose(...)
//  AddUsers(...)
//}

type UserService struct {
	// Cria uma composição para não precisar ficar implementando isto
	pb.UnimplementedUserServiceServer
}

// Criando um construtor
func NewUserService() *UserService {
	return &UserService{}
}

/*
	AddUser recebe:
		Contexto do tipo context.Context
		Request do tipo *pb.User
	AddUser retorna
		*pb.User
		error
*/
func (*UserService) AddUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	// Insert - Database
	fmt.Println(req.Name)

	return &pb.User{
		Id:    "123",
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}, nil // No lugar de nil, poderíamos retornar o erro, caso tivesse
}

/*
	AddUserVerbose request:
		Request do tipo *pb.User
	AddUserVerbose response:
		stream do tipo pb.UserService_AddUserVerboseServer
	AddUserVerbose retorno:
		error
*/
func (*UserService) AddUserVerbose(req *pb.User, stream pb.UserService_AddUserVerboseServer) error {
	// Insert - Database
	fmt.Println(req.Name)

	// Mandar pedaço por pedaço
	stream.Send(&pb.UserResultStream{
		Status: "Init",
		User:   &pb.User{},
	})

	//Aguarda 3s
	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "Inserting",
		User:   &pb.User{},
	})

	//Aguarda 3s
	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "User has been inserted",
		User: &pb.User{
			Id:    "123",
			Name:  req.GetName(),
			Email: req.GetEmail(),
		},
	})

	//Aguarda 3s
	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "Completed!",
		User: &pb.User{
			Id:    "123",
			Name:  req.GetName(),
			Email: req.GetEmail(),
		},
	})

	//Aguarda 3s
	time.Sleep(time.Second * 3)

	return nil
}

func (*UserService) AddUsers(stream pb.UserService_AddUsersServer) error {
	// Lista vazia de users
	users := []*pb.User{}

	for {
		// Recebendo a streaming de dados num looping infinito
		req, err := stream.Recv()
		// Se o client parou de mandar, enviamos a coleção de user e fechamos a conexão
		if err == io.EOF {
			return stream.SendAndClose(&pb.Users{
				User: users,
			})
		}
		if err != nil {
			log.Fatalf("Error receiving stream: %V", err)
		}

		//Adiciona cada novo usuário
		users = append(users, &pb.User{
			Id:    req.GetId(),
			Name:  req.GetName(),
			Email: req.GetEmail(),
		})

		fmt.Println("Adding ", req.GetName())
	}
}

func (*UserService) AddUserStreamBoth(stream pb.UserService_AddUserStreamBothServer) error {
	for {
		// Recebendo a streaming de dados num looping infinito
		req, err := stream.Recv()
		// Se o client parou de mandar, enviamos a coleção de user e fechamos a conexão
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error receiving stream from the client: %V", err)
		}

		err = stream.Send(&pb.UserResultStream{
			Status: "Added",
			User:   req,
		})
		if err != nil {
			log.Fatalf("Error sending stream to the client: %V", err)
		}
	}
}
