package repository_query

var SQL_create_order_product = `INSERT INTO OrderProduct(orderId,productId,quantity) VALUE=(?,?,?)`
