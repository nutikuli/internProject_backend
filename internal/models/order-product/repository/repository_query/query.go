package repository_query

var SQL_create_order_product = `INSERT INTO OrderProduct(orderId,productId,quantity) VALUES(?,?,?)`
var SQL_get_order_products_by_order_id = `SELECT * FROM OrderProduct WHERE orderId = ?`
var SQL_get_order_product_by_product_id = `SELECT * FROM OrderProduct WHERE productId = ?;`
