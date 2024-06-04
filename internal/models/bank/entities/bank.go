package entities

type Bank struct {
	Id        int64  `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	AccNumber string `json:"acc_number" db:"accNumber"`
	AccName   string `json:"acc_name" db:"accName"`
	AvatarUrl string `json:"avatar_url" db:"avatarUrl"`
	Status    string `json:"status" db:"status"`
	StoreId   int64  `json:"store_id" db:"storeId"`
	CreatedAt string `json:"createdAt" db:"createdAt"`
	UpdatedAt string `json:"updatedAt" db:"updatedAt"`
}
