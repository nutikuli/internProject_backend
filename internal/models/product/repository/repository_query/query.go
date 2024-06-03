package repository_query

var (
	InsertProduct        = `INSERT INTO Product (name, detail, price, status, productAvatar, stock, categoryId, storeId ) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	GetProductById       = `SELECT * FROM Product WHERE id = ?`
	GetProductsByStoreId = `SELECT * FROM Product WHERE storeId = ?`
	DeleteProductById    = `DELETE FROM Product WHERE id = ?`
	GetProductsByOrderId = `SELECT * FROM Product WHERE id IN (SELECT productId FROM OrderDetail WHERE orderId = ?)`
	UpdateProductById    = `UPDATE Product SET name = ?, detail = ?, price = ?, status = ?, productAvatar = ?, stock = ?, categoryId = ?, storeId = ? WHERE id = ?`
	GetAllProducts       = `SELECT * FROM Product`
)
