package entities

type LogGetReq struct {
	Fullname      string `db:"fullname"`
	MenuRequest   string `db:"menuRequest"`
	ActionRequest string `db:"actionRequest"`
}
