package entities

type UpdatePass struct {
	Id       int64  `json:"id" db:"id"`
	Password string `json:"password" db:"password"`
	// TODO: เพิ่ม file entites ที่เป็น array
	Role string `json:"role" db:"role"`
}
