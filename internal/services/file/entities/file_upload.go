package entities

type FileUploaderReq struct {
	EntityType string `json:"entity_type" form:"entity_type" validate:"required" binding:"required"`
	EntityId   int64  `json:"entity_id" form:"entity_id" validate:"required" binding:"required"`
	FileName   string `json:"file_name" form:"file_name" validate:"required" binding:"required"`
	FileData   string `json:"file_data" form:"file_data" validate:"base64" binding:"required"`
	FileType   string `json:"file_type" form:"file_type" validate:"required" binding:"required"`
}
