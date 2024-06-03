package repository_query

var (
	InsertProductCategory         = `INSERT INTO ProductCategory (name, code, detail, status) VALUES (?, ?, ?, ?)`
	GetProductCategoryById        = `SELECT * FROM ProductCategory WHERE id = ?`
	GetProductCategoriesByStoreId = `SELECT * FROM ProductCategory WHERE code = ?`
	DeleteProductCategoryById     = `DELETE FROM ProductCategory WHERE id = ?`
	UpdateProductCategoryById     = `UPDATE ProductCategory SET name = ?, code = ?, detail = ?, status = ? WHERE id = ?`
)
