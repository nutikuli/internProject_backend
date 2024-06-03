package repository_query

var SQL_get_account_customer = `SELECT id,name,password,phone,location,email,status FROM Account WHERE role = ?;`
var SQL_get_account_customer_by_id = `SELECT email, password, name, phone, location, status, role FROM Account WHERE id = ?`

var SQL_create_account_customer = `INSERT INTO Account(name,password,phone,location,email,role,status) VALUE(?,?,?,?,?,?,?)`

var SQL_update_password_account_customer = `UPDATE Account SET password = ?  WHERE id = ? AND role = ?;`
var SQL_update_account_customer = `UPDATE Account SET name= ?, password = ? ,phone = ?,location = ?,email = ?,status = ? WHERE id = ? ;`
var SQL_delete_account_customer = `DELETE FROM  Account  WHERE id = ? ;`
