package dtos

import "github.com/nutikuli/internProject_backend/internal/models/order/entities"

type OrderTransportStateReq struct {
	StateData     *entities.OrderStateReq           `json:"state_data" form:"state_data" binding:"required"`
	TransportData *entities.OrderTransportDetailReq `json:"transport_data" form:"transport_data"`
}
