package entities 


type Adminpermission struct{
	// Id int `db:"id"`
	menuPermission []string `db:"menuPermission"`
}