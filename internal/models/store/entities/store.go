package entities

import (
	"github.com/nutikuli/internProject_backend/internal/models/account/entities"
)

type Store struct {
	StoreName     string `json:"store_name" db:"storeName"`
	StoreLocation string `json:"store_location" db:"storeLocation"`
	entities.Account
}
