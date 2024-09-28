package dto

type CreateItemDTO struct {
	Name    string  `json:"name" validate:"required"` // required field
	Comment *string `json:"comment" default:""`       // optional field with default value
}
