package repository_query

var SQL_get_order = `SELECT id,orderId,totalAmount,topic,sumPrice,state,deliveryType,parcelNumber,sentDate,customerId,storeId,bankId,createdAt,UpdatedAt FROM Order = ?;`
var SQL_get_order_by_storeId = `SELECT id,orderId,totalAmount,topic,sumPrice,state,deliveryType,parcelNumber,sentDate,customerId,storeId,bankId,createdAt,UpdatedAt FROM Order WHERE storeId = ?;`
