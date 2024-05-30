package repository_query

var (
	InsertProductCategory         = `INSERT INTO ProductCategory (name, code, detail, status) VALUES (?, ?, ?, ?)`
	GetProductCategoryById        = `SELECT * FROM ProductCategory WHERE id = ?`
	GetProductCategoriesByStoreId = `SELECT * FROM ProductCategory WHERE storeId = ?`
	DeleteProductCategoryById     = `DELETE FROM ProductCategory WHERE id = ?`
)
