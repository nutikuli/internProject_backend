package repository_query

var (
	InsertStoreAccount         = `INSERT INTO Account (email, password, name, phone, location, status, role, storeName, storeLocation) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	GetStoreAccountById        = `SELECT id, email, password, name, phone, location, status, role, storeName, storeLocation FROM Account WHERE id = ?`
	UpdateStoreAccountPassword = `UPDATE Account SET password = ? WHERE id = ? AND role = ?`
	UpdateStoreById            = `UPDATE Account SET email = ?, name = ?, phone = ?, location = ?,  status = ?, storeName = ?, storeLocation = ? WHERE id = ? AND role = ?`
	DeleteStoreById            = `DELETE FROM Account WHERE id = ?`
)
