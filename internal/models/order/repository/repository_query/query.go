package repository_query

var SQL_get_order_by_customer = `SELECT id,orderId,totalAmount,topic,sumPrice,state,deliveryType,parcelNumber,sentDate,customerId,storeId,bankId,createdAt,UpdatedAt FROM Order WHERE customerId = ?;`
var SQL_get_order_by_Id = `SELECT id,orderId,totalAmount,topic,sumPrice,state,deliveryType,parcelNumber,sentDate,customerId,storeId,bankId,createdAt,UpdatedAt FROM Order WHERE Id = ?;`
var SQL_create_order = `INSERT INTO Order(id,orderId,totalAmount,topic,sumPrice,state,deliveryType,parcelNumber,sentDate,customerId,storeId,bankId,createdAt,UpdatedAt) VALUE(?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
