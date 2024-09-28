package dto

type UpdateItemDTO struct {
	ID      int     `json:"id" validate:"required"`   // ID field for the item
	Name    string  `json:"name" validate:"required"` // required field
	Comment *string `json:"comment" default:""`       // optional field with default value
}
