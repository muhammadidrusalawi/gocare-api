package response

import "time"

type AddressResponse struct {
	ID           string     `json:"id"`
	ReceiverName string     `json:"receiver_name"`
	Phone        string     `json:"phone"`
	Province     string     `json:"province"`
	City         string     `json:"city"`
	District     string     `json:"district"`
	PostalCode   string     `json:"postal_code"`
	FullAddress  string     `json:"full_address"`
	IsDefault    bool       `json:"is_default"`
	Label        string     `json:"label"`
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
}
