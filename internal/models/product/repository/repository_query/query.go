package repository_query

var (
	InsertProduct        = `INSERT INTO Product (name, detail, price, status, productAvatar, stock, categoryId, storeId ) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	GetProductById       = `SELECT * FROM Product WHERE id = ?`
	GetProductsByStoreId = `SELECT * FROM Product WHERE storeId = ?`
	DeleteProductById    = `DELETE FROM Product WHERE id = ?`
	GetProductsByOrderId = "SELECT p.* FROM Product p JOIN OrderProduct op ON p.id = op.productId JOIN `Order` o ON op.orderId = o.id	WHERE o.id = ?;"
	UpdateProductById    = `UPDATE Product SET name = ?, detail = ?, price = ?, status = ?, productAvatar = ?, stock = ?, categoryId = ?, storeId = ? WHERE id = ?`
	GetAllProducts       = `SELECT * FROM Product`
)
