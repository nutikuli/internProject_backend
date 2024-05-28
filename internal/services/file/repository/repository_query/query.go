package repository_query

var (
	FileInsertByEntityAndId      = `INSERT INTO File (name, pathUrl, type,  entityType, entityId) VALUES (?, ?, ?, ?, ?)`
	QueryFileSelectAll           = `SELECT * FROM File`
	QueryFileSelectByIdAndEntity = `SELECT * FROM File WHERE entityType = ? AND entityId = ?`
	ExecFileDeleteByIdAndEntity  = `DELETE FROM File WHERE entityType = ? AND entityId = ?`
)
