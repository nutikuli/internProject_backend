package dtos

type AccountLoginRes struct {
	UserToken   UserToken   `json:"token"`
	AccountData interface{} `json:"account_data"`
}
