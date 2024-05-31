package repository_query

var SQL_get_account_customer = `SELECT id,name,password,phone,location,email,status FROM Account WHERE role = ?;`
var SQL_get_account_customer_by_id = `SELECT id,name,password,phone,location,email,status FROM Account WHERE role = ? and id = ?;`
var SQL_create_account_customer = `INSERT INTO Account(id,name,password,phone,location,email,role,status) VALUE(?,?,?,?,?,?,?,?)`
var SQL_update_password_account_customer = `UPDATE Account SET password = ?  WHERE id = ? ;`
var SQL_update_account_customer = `UPDATE Account SET name= ?, password = ? ,phone = ?,location = ?,email = ?,status = ?, createAt=? WHERE id = ? ;`
var SQL_delete_account_customer = `DELETE FROM  Account  WHERE id = ? ;`
