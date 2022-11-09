package repository

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/pastorilps/propostas_populares/app/docs"
	middleware "github.com/pastorilps/propostas_populares/middlewares"
	"github.com/pastorilps/propostas_populares/users/domain"
	"github.com/pastorilps/propostas_populares/users/entity"
	"github.com/sirupsen/logrus"
)

type userRepository struct {
	Conn *sql.DB
}

func NewUserRepo(Conn *sql.DB) domain.UserRepository {
	return &userRepository{Conn}
}

func (u *userRepository) DeleteUser(id int16) (err error) {
	query := `delete from public.user where id = $1`

	stmt, err := u.Conn.Prepare(query)
	if err != nil {
		logrus.Error("error pushing data for database", err)
		return
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return
	}

	rowsAfected, err := res.RowsAffected()
	if err != nil {
		return
	}

	if rowsAfected != 1 {
		logrus.Error("Weird Behavior, Total Affeced: %d", rowsAfected)
		return
	}

	return
}

func (u *userRepository) UpdateUser(ctx context.Context, er *entity.Receive_User) (err error) {
	query := `update public.user set name = $2, email = $3, password = $4, picture = $5, newsletter = $6 where id = $1`

	stmt, err := u.Conn.PrepareContext(ctx, query)
	if err != nil {
		log.Fatalln("Error in receive parameters context", err)
		return
	}

	res, err := stmt.ExecContext(ctx, er.ID, er.Name, er.Email, middleware.SHA256Encoder(er.Password), er.Picture, er.Newsletter)
	if err != nil {
		log.Fatalln("Error when sending data context for query consult", err)
		return
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return
	}

	if affect != 1 {
		log.Fatalln("Weird Behavior. Total Affected", affect)
		return
	}

	return
}

func (u *userRepository) DataFetchUpdateUser(ctx context.Context, id int16) (es *entity.Send_User, err error) {
	query := `select * from public.user where id = $1`

	list, err := u.fetchUpdateUser(ctx, query, id)
	if err != nil {
		log.Fatalln("Error when consulting data in query database", err)
		return
	}

	pkg := &entity.Send_User{}
	if len(list) > 0 {
		pkg = &list[0]
	} else {
		return nil, middleware.ErrorNotFound
	}

	return pkg, nil

}

func (u *userRepository) fetchUpdateUser(ctx context.Context, query string, args ...interface{}) (es []entity.Send_User, err error) {
	rows, err := u.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		log.Fatalln("Error in consult data query", err)
		return
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			log.Fatalln("Error in line closed", err)
		}
	}()

	res := make([]entity.Send_User, 0)
	for rows.Next() {
		str := entity.Send_User{}
		err = rows.Scan(
			&str.ID,
			&str.Name,
			&str.Email,
			&str.Password,
			&str.Picture,
			&str.Newsletter,
		)

		if err != nil {
			log.Fatalln("No Scanner Error", err)
			return
		}

		res = append(res, str)
	}

	return res, nil
}

func (u *userRepository) CreateUser(es *entity.Users) (ds *entity.Users, err error) {
	query := `insert into public.user (name, email, password, picture, newsletter) values ($1, $2, $3, $4, $5)`
	es.Password = middleware.SHA256Encoder(es.Password)

	stmt, err := u.Conn.Prepare(query)
	if err != nil {
		logrus.Error("Error in pushing data in database", err)
		return
	}

	res, err := stmt.Exec(
		es.Name,
		es.Email,
		es.Password,
		es.Picture,
		es.Newsletter,
	)
	if err != nil {
		logrus.Error("Error to attach data struct database", err)
		return
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return
	}

	if affect != 1 {
		logrus.Error("Error in affect in database", err)
		return
	}

	return

}

func (u *userRepository) FetchUserBydId(id int16) (*entity.Users, error) {
	query := `select * from public.user  where id = $1`

	list, err := u.fetchUserById(query, id)
	if err != nil {
		return nil, err
	}

	pkg := &entity.Users{}
	if len(list) > 0 {
		pkg = list[0]
	} else {
		return nil, middleware.ErrorNotFound
	}

	return pkg, nil
}

func (u *userRepository) fetchUserById(query string, args ...interface{}) ([]*entity.Users, error) {
	rows, err := u.Conn.Query(query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer rows.Close()
	res := make([]*entity.Users, 0)
	for rows.Next() {
		str := new(entity.Users)
		err := rows.Scan(
			&str.ID,
			&str.Name,
			&str.Email,
			&str.Password,
			&str.Picture,
			&str.Newsletter,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		res = append(res, str)
	}

	return res, nil
}

func (u *userRepository) FetchAllUsers() ([]*entity.Users, error) {
	query := `select * from public.user order by id asc`

	return u.fetchAllUsers(query)
}

func (u *userRepository) fetchAllUsers(query string, args ...interface{}) ([]*entity.Users, error) {
	rows, err := u.Conn.Query(query, args...)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer rows.Close()

	result := make([]*entity.Users, 0)
	for rows.Next() {
		us := new(entity.Users)
		err = rows.Scan(
			&us.ID,
			&us.Name,
			&us.Email,
			&us.Password,
			&us.Picture,
			&us.Newsletter,
		)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, us)
	}

	return result, nil
}
