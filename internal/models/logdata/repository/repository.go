package repository

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"
	"github.com/nutikuli/internProject_backend/internal/models/logdata"
	"github.com/nutikuli/internProject_backend/internal/models/logdata/entities"
	"github.com/nutikuli/internProject_backend/internal/models/logdata/repository/repository_query"
)

type logRepo struct {
	db *sqlx.DB
}

func NewLoggerRepository(db *sqlx.DB) logdata.LogRepository {
	return &logRepo{
		db: db,
	}
}

// ld = log data

func (ld *logRepo) CreateLogData(ctx context.Context, logdata *entities.LogCreateReq) (*int64, error) {

	res, err := ld.db.ExecContext(ctx, repository_query.SQL_insert_logdata, logdata.Fullname, logdata.MenuRequest, logdata.ActionRequest)
	if err != nil {
		log.Info(err)
		return nil, err
	}
	createdId, err := res.LastInsertId()
	if err != nil {
		log.Info(err)
		return nil, err
	}

	return &createdId, nil
}

func (ld *logRepo) GetLogDatas(ctx context.Context) ([]entities.LogGetReq, error) {

	var logdata = make([]entities.LogGetReq, 0)

	err := ld.db.SelectContext(ctx, &logdata, repository_query.SQL_get_logdata)
	if err != nil {
		log.Info(err)
		return nil, err
	}

	return logdata, nil
}
