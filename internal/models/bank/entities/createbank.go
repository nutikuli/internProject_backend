package entities

// TODO:นำตัวไปใช้ใน Log Model
type BankCreatedReq struct {
	Name      string `json:"name" from:"name" binding:"required"`
	AccNumber string `json:"acc_number" from:"acc_number" binding:"required"`
	AccName   string `json:"acc_name" from:"acc_name" binding:"required"`
	AvatarUrl string `json:"avatar_url" from:"avatar_url" binding:"required"`
	StoreId   int64  `json:"store_id" from:"store_id" binding:"required"`
	Status    bool   `json:"status" from:"status" binding:"required"`
}
