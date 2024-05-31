package repository_query

var SQL_get_order_by_customerId = `SELECT * FROM Order WHERE customerId = ?;`
var SQL_get_order_by_Id = `SELECT * FROM Order WHERE id = ?;`
var SQL_get_order_by_storeId = `SELECT * FROM Order WHERE storeId = ?;`
var SQL_create_order = `INSERT INTO Order(totalAmount,topic,sumPrice,state,deliveryType,parcelNumber,sentDate,customerId,storeId,bankId) VALUE(?,?,?,?,?,?,?,?,?,?,?)`
var SQL_update_order_transport_detail = `UPDATE Order SET deliveryType = ?, parcelNumber = ?, sentDate = ? WHERE id = ?;`
var SQL_update_order_state = `UPDATE Order SET state = ? WHERE id = ?;`
