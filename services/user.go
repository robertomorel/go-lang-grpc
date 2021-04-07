package services

import (
	"context"
	"fmt"

	"github.com/robertomorel/go-lang-grpc/pb"
)

// Precisamos implementar esse type, que está no arquivo "pb/user_grpc.pb.go"
//type UserServiceServer interface {
// Criando um serviço do RPC com uma função AddUser
//	AddUser(context.Context, *User) (*User, error)
//	mustEmbedUnimplementedUserServiceServer()
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
