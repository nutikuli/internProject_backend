package entities

type FileEntityReq struct {
	EntityType string `json:"entity_type" form:"entity_type" validate:"required" binding:"required"`
	EntityId   int64  `json:"entity_id" form:"entity_id" validate:"required" binding:"required"`
}
