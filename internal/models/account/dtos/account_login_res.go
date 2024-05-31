package dtos

type AccountLoginRes struct {
	UserToken   UserToken   `json:"user_token"`
	AccountData interface{} `json:"account_data"`
}
