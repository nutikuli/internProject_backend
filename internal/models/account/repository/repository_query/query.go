package repository_query
//get customer account
var SQL_get_account_customer = `SELECT id,name,password,phone,location,email,status FROM Account WHERE role = ?;`
// get store account
var SQL_get_account_storeaccount = `SELECT id,name,password,phone,location,email,status,storeName,storeLocation FROM Account WHERE role = ?;`
// get admin account
var SQL_get_account_admin = `SELECT id,name,password,phone,location,email,status,permissionId FROM Account WHERE role = ?;`
//get customer account by id
var SQL_get_account_customer_by_id = `SELECT id,name,password,phone,location,email,status FROM Account WHERE role = ? and id = ?;`
// get store account by id
var SQL_get_account_storeaccount_by_id = `SELECT id,name,password,phone,location,email,status,storeName,storeLocation FROM Account WHERE role = ? and id = ?;`
// get admin account by id
var SQL_get_account_admin_by_id = `SELECT id,name,password,phone,location,email,status,permissionId FROM Account WHERE role = ? and id = ?;`