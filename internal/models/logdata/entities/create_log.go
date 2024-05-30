package entities

type LogCreateReq struct {
	Fullname      string `json:"fullname" from:"full_name" binding:"required"`
	MenuRequest   string `json:"menuRequest" from:"menu_request" binding:"required"`
	ActionRequest string `json:"actionRequest" from:"action_request" binding:"required"`
}
