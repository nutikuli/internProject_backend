package repository_query

var SQL_get_order_by_customerId = `SELECT * FROM Order WHERE customerId = ?;`
var SQL_get_order_by_Id = `SELECT * FROM Order WHERE Id = ?;`
var SQL_get_order_by_storeId = `SELECT * FROM Order WHERE storeId = ?;`
var SQL_create_order = `INSERT INTO Order(totalAmount,topic,sumPrice,state,deliveryType,parcelNumber,sentDate,customerId,storeId,bankId) VALUE(?,?,?,?,?,?,?,?,?,?,?)`
