package repository_query

// get admin account
var SQL_get_account_admin = `SELECT id,name,password,phone,location,email,status,permissionId FROM Account WHERE role = ?;`

var SQL_getall_account_admin = `SELECT id,name,password,phone,location,email,status,permissionId,role FROM Account WHERE role = 'ADMIN';`

// get admin account by id
var SQL_get_account_admin_by_id = `SELECT id,name,password,phone,location,email,status,permissionId FROM Account WHERE role = ? and id = ?;`

var SQL_get_account_adminid = `SELECT email, password, name, phone, location, status, role  FROM Account WHERE id = ?`

//insert admin account

var SQL_insert_account_admin = `INSERT INTO Account (name,password,phone,location,email,role,status,permissionid) VALUE(?,?,?,?,?,?,?,?) ;`

//update admin account

var SQL_update_account_admin = `UPDATE  Account SET name= ?, password = ? ,phone = ?,location = ? ,email = ?,status = ? WHERE id = ? ;`

//update password account admin

var SQL_update_password_account_admin = `UPDATE Account SET password = ?  WHERE id = ? and role = ? ;`

//delete admin account

var SQL_delete_account_admin = `DELETE FROM  Account  WHERE id = ? ;`
