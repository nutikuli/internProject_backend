package repository_query 

// get admin account
var SQL_get_account_admin = `SELECT id,name,password,phone,location,email,status,permissionId FROM Account WHERE role = ?;` 

// get admin account by id
var SQL_get_account_admin_by_id = `SELECT id,name,password,phone,location,email,status,permissionId FROM Account WHERE role = ? and id = ?;`

//insert admin account

var SQL_insert_account_admin = `INSERT INTO Account (name,password,phone,location,email,status,permissionId,createAt) VALUE(?,?,?,?,?,?,?,?) ;`

//update admin account 

var SQL_update_account_admin = `UPDATE  Account SET name= ?, password = ? ,phone = ?,location = ?,email = ?,status = ?,permissionId = ?,createAt=? WHERE id = ? ;`

//update password account admin

var SQL_update_password_account_admin = `UPDATE Account SET password = ?  WHERE id = ? and role = ? ;`

//delete admin account 

var SQL_delete_account_admin = `DELETE FROM  Account  WHERE id = ? ;`

