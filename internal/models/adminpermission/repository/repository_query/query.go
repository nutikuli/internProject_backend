package repositoryquery


//INSERT ADMIN PERMISSION

var SQL_insert_permission_admin = `INSERT INTO AdminPermission (menuPermission,roleName) VALUE(?,?) ;`

//GET ADMIN PERMISSION

var SQL_get_permisson_admin = `SELECT menuPermission , roleName FROM AdminPermission WHERE id = ? ;`

//GET ADMIN PERMISSION BY ID
var SQL_get_adminpermission_by_id = `SELECT menuPermission , roleName FROM AdminPermission WHERE id = ? ;` 

//UPDATE ADMIN PERMISSION BY ID
var SQL_update_adminperrmision_by_id = `UPDATE AdminPermission SET menuPermission = ? , roleName = ? WHERE id = ? ;` 

//DELETE ADMIN PERMISSION BY ID
var SQL_delete_adminpermission_by_id = `DELETE FROM  AdminPermission  WHERE id = ? ;` 


//GET ALL
var SQL_getall_adminpermission = `SELECT * FROM  AdminPermission  ;`