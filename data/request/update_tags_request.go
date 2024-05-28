package request

type UpdateTagsRequest struct {
	ID   int    `validate:"required"`
	Name string `json:"name" validate:"required,min=1,max=255"`
}
