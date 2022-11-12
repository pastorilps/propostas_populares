package authenticate

import (
	"context"

	entity "github.com/pastorilps/propostas_populares/authenticate/entity"
)

type AuthenticateRepository interface {
	GetByUserLogin(ctx context.Context, req *entity.Receive_Login_Data) (res *entity.Send_User_Data, err error)
}

type AuthenticateUsecase interface {
	GetByUserLogin(ctx context.Context, req *entity.Receive_Login_Data) (tko *entity.Auth_Token, err error)
}
