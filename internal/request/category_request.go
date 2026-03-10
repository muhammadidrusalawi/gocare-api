package request

type CreateCategoryRequest struct {
	Name        string  `json:"name" validate:"required,max=255"`
	Description *string `json:"description" validate:"omitempty,max=1000"`
}

type UpdateCategoryRequest struct {
	Name        string  `json:"name" validate:"omitempty,max=255"`
	Description *string `json:"description" validate:"omitempty,max=1000"`
}
