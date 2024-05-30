package logdata

import ("context" 
"github.com/nutikuli/internProject_backend/internal/models/logdata/entities"

)

type LogRepository interface {
	CreateLogData(ctx context.Context, logdata *entities.LogCreateReq) (*int64, error)
	GetLogDatas(ctx context.Context) ([]entities.LogGetReq, error)
}
