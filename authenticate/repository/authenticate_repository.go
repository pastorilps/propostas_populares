package repository

import (
	"context"
	"database/sql"

	domain "github.com/pastorilps/propostas_populares/authenticate/domain"
	"github.com/pastorilps/propostas_populares/authenticate/entity"
	"github.com/pastorilps/propostas_populares/middlewares"
	"github.com/sirupsen/logrus"
)

type authenticateRepository struct {
	DbConn *sql.DB
}

func NewAuthenticateRepo(DbConn *sql.DB) domain.AuthenticateRepository {
	return &authenticateRepository{DbConn}
}

func (ar *authenticateRepository) GetByUserLogin(ctx context.Context, req *entity.Receive_Login_Data) (res *entity.Send_User_Data, err error) {
	query := `select id, name, picture, newsletter from public.user where email = $1 and password = $2`

	list, err := ar.fetch(ctx, query, req.Username, middlewares.SHA256Encoder(req.Password))
	if err != nil {
		logrus.Error(err)
		return
	}

	pkg := &entity.Send_User_Data{}
	if len(list) > 0 {
		pkg = list[0]
	} else {
		return nil, middlewares.ErrorNotFound
	}

	return pkg, nil
}

func (ar *authenticateRepository) fetch(ctx context.Context, query string, args ...interface{}) (res []*entity.Send_User_Data, err error) {
	row, err := ar.DbConn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return
	}

	defer func() {
		errRow := row.Close()
		if errRow != nil {
			logrus.Error(errRow)
		}
	}()

	result := make([]*entity.Send_User_Data, 0)
	for row.Next() {
		as := new(entity.Send_User_Data)

		err = row.Scan(
			&as.UserID,
			&as.UserName,
			&as.UserPicture,
			&as.UserNews,
		)

		if err != nil {
			logrus.Error(err)
			return
		}
		result = append(result, as)
	}

	return result, nil
}
