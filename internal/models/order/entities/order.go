package entities

type Order struct {
	Id           int64   `db:"id"`
	TotalAmount  float64 `db:"totalAmount"`
	Topic        string  `db:"topic"`
	SumPrice     float64 `db:"sumPrice"`
	State        string  `db:"state"`
	DeliveryType string  `db:"deliveryType"`
	ParcelNumber string  `db:"parcelNumber"`
	SentDate     string  `db:"sentDate"`
	CustomerId   int64   `db:"customerId"`
	StoreId      int64   `db:"storeId"`
	BankId       int64   `db:"bankId"`
	CreatedAt    string  `db:"createdAt"`
	UpdatedAt    string  `db:"updatedAt"`
}
