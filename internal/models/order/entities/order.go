package entities

type Order struct {
	Id           int64   `json:"id" db:"id"`
	TotalAmount  float64 `json:"total_amount" db:"totalAmount"`
	Topic        string  `json:"topic" db:"topic"`
	SumPrice     float64 `json:"sum_price" db:"sumPrice"`
	State        string  `json:"state" db:"state"`
	DeliveryType string  `json:"delivery_type" db:"deliveryType"`
	ParcelNumber string  `json:"parcel_number" db:"parcelNumber"`
	SentDate     string  `json:"sent_date" db:"sentDate"`
	CustomerId   int64   `json:"customer_id" db:"customerId"`
	StoreId      int64   `json:"store_id" db:"storeId"`
	BankId       int64   `json:"bank_id" db:"bankId"`
	CreatedAt    string  `json:"created_at" db:"createdAt"`
	UpdatedAt    string  `json:"updated_at" db:"updatedAt"`
}
