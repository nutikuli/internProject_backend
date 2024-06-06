package entities 


type Adminpermission struct{
	Id int `db:"id"`
	MenuPermission string `db:"menuPermission"`
	Rolename string `db:"roleName"`
}