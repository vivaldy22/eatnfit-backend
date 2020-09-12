package user

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	authproto "github.com/vivaldy22/eatnfit-auth-service-grpc/proto"
	"github.com/vivaldy22/eatnfit-auth-service-grpc/tools/queries"
	"strconv"
)

type Service struct{
	db *sql.DB
}

func NewService(db *sql.DB) authproto.UserCRUDServer {
	return &Service{db}
}

func (s *Service) GetAll(ctx context.Context, pagination *authproto.Pagination) (*authproto.UserList, error) {
	var users = new(authproto.UserList)
	page, _ := strconv.Atoi(pagination.Page)
	limit, _ := strconv.Atoi(pagination.Limit)
	offset := (page * limit) - limit
	query := fmt.Sprintf(queries.GET_ALL_USER, offset, limit)
	rows, err := s.db.Query(query, "%"+pagination.Keyword+"%", "%"+pagination.Keyword+"%")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var each = new(authproto.User)
		if err := rows.Scan(&each.UserId, &each.UserEmail, &each.UserPassword, &each.UserFName, &each.UserLName,
			&each.UserGender, &each.UserBalance, &each.UserLevel, &each.UserStatus); err != nil {
			return nil, err
		}
		users.List = append(users.List, each)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (s *Service) GetTotal(ctx context.Context, e *empty.Empty) (*authproto.Total, error) {
	var total int
	row := s.db.QueryRow(queries.GET_TOTAL_USER)
	err := row.Scan(&total)
	if err != nil {
		return nil, err
	}
	return &authproto.Total{TotalData: strconv.Itoa(total)}, nil
}

func (s *Service) GetByID(ctx context.Context, id *authproto.ID) (*authproto.User, error) {
	var user = new(authproto.User)
	row := s.db.QueryRow(queries.GET_BY_ID_USER, id.Id)

	err := row.Scan(&user.UserId, &user.UserEmail, &user.UserPassword, &user.UserFName, &user.UserLName,
		&user.UserGender, &user.UserBalance, &user.UserLevel, &user.UserStatus)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) GetByEmail(ctx context.Context, email *authproto.Email) (*authproto.User, error) {
	var user = new(authproto.User)
	row := s.db.QueryRow(queries.GET_BY_EMAIL_USER, email.Email)

	err := row.Scan(&user.UserId, &user.UserEmail, &user.UserPassword, &user.UserFName, &user.UserLName,
		&user.UserGender, &user.UserBalance, &user.UserLevel, &user.UserStatus)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) Create(ctx context.Context, user *authproto.User) (*authproto.User, error) {
	tx, err := s.db.Begin()

	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare(queries.CREATE_USER)

	if err != nil {
		return nil, err
	}

	id := uuid.New().String()
	_, err = stmt.Exec(id, user.UserEmail, user.UserPassword, user.UserFName, user.UserLName,
		user.UserGender, user.UserLevel)
	if err != nil {
		return nil, tx.Rollback()
	}

	user.UserId = id
	stmt.Close()
	return user, tx.Commit()
}

func (s *Service) CreateByAdmin(ctx context.Context, user *authproto.User) (*authproto.User, error) {
	tx, err := s.db.Begin()

	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare(queries.CREATE_USER_BY_ADMIN)

	if err != nil {
		return nil, err
	}

	id := uuid.New().String()
	_, err = stmt.Exec(id, user.UserEmail, user.UserPassword, user.UserFName, user.UserLName,
		user.UserGender, user.UserBalance, user.UserLevel)
	if err != nil {
		return nil, tx.Rollback()
	}

	user.UserId = id
	stmt.Close()
	return user, tx.Commit()
}

func (s *Service) Update(ctx context.Context, request *authproto.UserUpdateRequest) (*authproto.User, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare(queries.UPDATE_USER)
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(request.User.UserEmail, request.User.UserFName, request.User.UserLName,
		request.User.UserGender, request.User.UserBalance, request.User.UserLevel, request.Id.Id)
	if err != nil {
		return nil, tx.Rollback()
	}

	stmt.Close()
	request.User.UserId = request.Id.Id
	return request.User, tx.Commit()
}

func (s *Service) Delete(ctx context.Context, id *authproto.ID) (*empty.Empty, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return new(empty.Empty), err
	}

	stmt, err := tx.Prepare(queries.DELETE_USER)
	if err != nil {
		return new(empty.Empty), err
	}

	_, err = stmt.Exec(id.Id)
	if err != nil {
		return new(empty.Empty), tx.Rollback()
	}

	stmt.Close()
	return new(empty.Empty), tx.Commit()
}
