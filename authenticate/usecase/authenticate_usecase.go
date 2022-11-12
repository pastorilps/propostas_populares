package usecase

import (
	"context"
	"time"

	domain "github.com/pastorilps/propostas_populares/authenticate/domain"
	"github.com/pastorilps/propostas_populares/authenticate/entity"
	jwtMiddleware "github.com/pastorilps/propostas_populares/middlewares"
	"github.com/sirupsen/logrus"
)

type authenticateUsecase struct {
	authenticateRepo domain.AuthenticateRepository
}

const layoutISO = "2006-01-02 00:00:00"

var expirationTime = time.Now().Add(24 * time.Hour)

func NewAuthenticateUsecase(dr domain.AuthenticateRepository) domain.AuthenticateUsecase {
	return &authenticateUsecase{
		authenticateRepo: dr,
	}
}

func (da *authenticateUsecase) GetByUserLogin(ctx context.Context, req *entity.Receive_Login_Data) (tko *entity.Auth_Token, err error) {
	res, err := da.authenticateRepo.GetByUserLogin(ctx, req)
	if err != nil {
		logrus.Error(err)
		return
	}

	strToken, err := da.generateToken(res)
	if err != nil {
		logrus.Error(err)
		return
	}

	ut := entity.Auth_Token{
		Username: req.Username,
		Token:    strToken,
		Expires:  expirationTime.Format(layoutISO),
	}

	return &ut, nil
}

func (da *authenticateUsecase) generateToken(res *entity.Send_User_Data) (tko string, err error) {
	mp := make(map[string]interface{})
	mp["UserID"] = res.UserID
	mp["UserName"] = res.UserName
	mp["UserPicture"] = res.UserPicture
	mp["UserNews"] = res.UserNews

	token, err := jwtMiddleware.GenerateJwt(mp)
	if err != nil {
		logrus.Error(err)
		return " ", err
	}

	return token, nil
}
