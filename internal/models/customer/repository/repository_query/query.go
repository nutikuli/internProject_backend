package repository_query

var SQL_get_account_customer = `SELECT id,name,password,phone,location,email,status FROM Account WHERE role = ?;`
var SQL_get_account_customer_by_id = `SELECT id,name,password,phone,location,email,status FROM Account WHERE role = ? and id = ?;`
var SQL_create_account_customer = `INSERT INTO Account(id,name,password,phone,location,email,role,status,) VALUE(?,?,?,?,?,?,?,?)`

// var SQL_create_order = `INSERT INTO Order(id,orderId,totalAmount,topic,sumPrice,state,deliveryType,parcelNumber,sentDate,customerId,storeId,bankId,createdAt,updatedAt) VALUES(?, ?, ?, ?, ?, ?);`
