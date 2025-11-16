package validation

type CreateNote struct {
	Title   string `json:"title" validate:"required,min=1,max=200"`
	Content string `json:"content" validate:"required"`
}

type UpdateNote struct {
	Title   *string `json:"title" validate:"omitempty,min=1,max=200"`
	Content *string `json:"content" validate:"omitempty"`
}
