package repository_query

var (
	InsertProduct        = `INSERT INTO Product (name, detail, price, status, productAvatar, stock, categoryId, storeId ) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	GetProductById       = `SELECT * FROM Product WHERE id = ?`
	GetProductsByStoreId = `SELECT * FROM Product WHERE storeId = ?`
)
