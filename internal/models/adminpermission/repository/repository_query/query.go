package repositoryquery


//INSERT ADMIN PERMISSION

var SQL_insert_permission_admin = `INSERT INTO AdminPermission (menuPermission) VALUE(?) ;`

//GET ADMIN PERMISSION

var SQL_get_permisson_admin = `SELECT menuPermission FROM AdminPermission WHERE id = ? ;`

//GET ADMIN PERMISSION BY ID
var SQL_get_adminpermission_by_id = `SELECT menuPermission FROM AdminPermission WHERE id = ? ;`