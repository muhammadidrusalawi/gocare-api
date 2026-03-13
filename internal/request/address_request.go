package request

type CreateAddressRequest struct {
	ReceiverName string `json:"receiver_name" validate:"required,max=255"`
	Phone        string `json:"phone" validate:"required,max=20"`
	Province     string `json:"province" validate:"required,max=255"`
	City         string `json:"city" validate:"required,max=255"`
	District     string `json:"district" validate:"required,max=255"`
	PostalCode   string `json:"postal_code" validate:"required,max=10"`
	FullAddress  string `json:"full_address" validate:"required"`
	Label        string `json:"label" validate:"required,oneof=home office apartment other"`
}

type UpdateAddressRequest struct {
	ReceiverName *string `json:"receiver_name" validate:"omitempty,max=255"`
	Phone        *string `json:"phone" validate:"omitempty,max=20"`
	Province     *string `json:"province" validate:"omitempty,max=255"`
	City         *string `json:"city" validate:"omitempty,max=255"`
	District     *string `json:"district" validate:"omitempty,max=255"`
	PostalCode   *string `json:"postal_code" validate:"omitempty,max=10"`
	FullAddress  *string `json:"full_address" validate:"omitempty"`
	Label        *string `json:"label" validate:"omitempty,oneof=home office apartment other"`
}
