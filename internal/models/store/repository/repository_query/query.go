package repository_query

var (
	InsertStoreAccount  = `INSERT INTO Account (email, password, name, phone, location, status, role, storeName, storeLocation) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	GetStoreAccountById = `SELECT email, password, name, phone, location, status, role, storeName, storeLocation FROM Account WHERE id = ?`
)
