package level

import (
	"context"
	"database/sql"
	"github.com/golang/protobuf/ptypes/empty"
	authproto "github.com/vivaldy22/eatnfit-auth-service-grpc/proto"
	"github.com/vivaldy22/eatnfit-auth-service-grpc/tools/queries"
	"strconv"
)

type Service struct{
	db *sql.DB
}

func NewService(db *sql.DB) authproto.LevelCRUDServer {
	return &Service{db}
}

func (s *Service) GetAll(ctx context.Context, empty *empty.Empty) (*authproto.LevelList, error) {
	var levels = new(authproto.LevelList)
	rows, err := s.db.Query(queries.GET_ALL_LEVEL)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var each = new(authproto.Level)
		if err := rows.Scan(&each.LevelId, &each.LevelName, &each.LevelStatus); err != nil {
			return nil, err
		}
		levels.List = append(levels.List, each)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return levels, nil
}

func (s *Service) GetByID(ctx context.Context, id *authproto.ID) (*authproto.Level, error) {
	var level = new(authproto.Level)
	row := s.db.QueryRow(queries.GET_BY_ID_LEVEL, id.Id)

	err := row.Scan(&level.LevelId, &level.LevelName, &level.LevelStatus)
	if err != nil {
		return nil, err
	}
	return level, nil
}

func (s *Service) Create(ctx context.Context, level *authproto.Level) (*authproto.Level, error) {
	tx, err := s.db.Begin()

	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare(queries.CREATE_LEVEL)

	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(level.LevelName)

	if err != nil {
		return nil, tx.Rollback()
	}

	lastInsertID, err := res.LastInsertId()

	if err != nil {
		return nil, tx.Rollback()
	}

	level.LevelId = strconv.Itoa(int(lastInsertID))
	stmt.Close()
	return level, tx.Commit()
}

func (s *Service) Update(ctx context.Context, request *authproto.LevelUpdateRequest) (*authproto.Level, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare(queries.UPDATE_LEVEL)
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(request.Level.LevelName, request.Id.Id)
	if err != nil {
		return nil, tx.Rollback()
	}

	stmt.Close()
	request.Level.LevelId = request.Id.Id
	return request.Level, tx.Commit()
}

func (s *Service) Delete(ctx context.Context, id *authproto.ID) (*empty.Empty, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return new(empty.Empty), err
	}

	stmt, err := tx.Prepare(queries.DELETE_LEVEL)
	if err != nil {
		return new(empty.Empty), err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return new(empty.Empty), tx.Rollback()
	}

	stmt.Close()
	return new(empty.Empty), tx.Commit()
}

