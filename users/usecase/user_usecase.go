package usecase

import (
	"context"
	"log"

	"github.com/pastorilps/propostas_populares/users/domain"
	"github.com/pastorilps/propostas_populares/users/entity"
	"github.com/sirupsen/logrus"
)

type userUseCase struct {
	userRepo domain.UserRepository
}

func NewUserUseCase(ur domain.UserRepository) domain.UserUseCase {
	return &userUseCase{
		userRepo: ur,
	}
}

func (u *userUseCase) DeleteUser(id int16) (err error) {
	err = u.userRepo.DeleteUser(id)
	if err != nil {
		return nil
	}

	return nil
}

func (u *userUseCase) UpdateUser(ctx context.Context, er *entity.Receive_User) (es *entity.Send_User, err error) {

	err = u.userRepo.UpdateUser(ctx, er)
	if err != nil {
		log.Fatalln("I am don't updated data register Contract Status in Usecase", err)
		return nil, err
	}

	res, err := u.userRepo.DataFetchUpdateUser(ctx, er.ID)
	if err != nil {
		log.Fatalln("Not return data fetch in database", err)
		return nil, err
	}

	return res, nil
}

func (u *userUseCase) CreateUsers(es *entity.Users) (*entity.Users, error) {
	res, err := u.userRepo.CreateUser(es)
	if err != nil {
		logrus.Error("Error return data for saving in database", err)
	}

	return res, err
}

func (u *userUseCase) GetUserById(id int16) (*entity.Users, error) {
	res, err := u.userRepo.FetchUserBydId(id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *userUseCase) GetAllUsers() ([]*entity.Users, error) {
	list, err := u.userRepo.FetchAllUsers()
	if err != nil {
		return nil, err
	}

	return list, nil
}
