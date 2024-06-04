package entities

type CustomerUpdateReq struct {
	Name     string `json:"name" from:"name" binding:"required"`
	Phone    string `json:"phone" from:"phone" binding:"required"`
	Location string `json:"location" from:"location" binding:"required"`
	Email    string `json:"email" from:"email" binding:"required"`
	Status   bool   `json:"status" from:"status" binding:"required"`
}
