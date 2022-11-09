package domain

import (
	"context"

	"github.com/pastorilps/propostas_populares/users/entity"
)

type UserUseCase interface {
	GetAllUsers() ([]*entity.Users, error)
	GetUserById(id int16) (*entity.Users, error)
	CreateUsers(es *entity.Users) (*entity.Users, error)
	UpdateUser(ctx context.Context, er *entity.Receive_User) (es *entity.Send_User, err error)
	DeleteUser(id int16) (err error)
}

type UserRepository interface {
	FetchAllUsers() ([]*entity.Users, error)
	FetchUserBydId(id int16) (*entity.Users, error)
	CreateUser(es *entity.Users) (ds *entity.Users, err error)
	UpdateUser(ctx context.Context, er *entity.Receive_User) (err error)
	DataFetchUpdateUser(ctx context.Context, id int16) (es *entity.Send_User, err error)
	DeleteUser(id int16) (err error)
}
