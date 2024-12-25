package dtos

type CreateAuthorDto struct {
	FirstName   string  `json:"first_name" validate:"required"`
	LastName    *string `json:"last_name" `
	PhoneNumber *string `json:"phone_number"`
	Email       string  `json:"email" validate:"required"`
	Biography   *string `json:"biography"`
	Education   *string `json:"education"`
}

type CreateAuthorResponseDto struct {
	FirstName   string  `json:"first_name"`
	LastName    *string `json:"last_name" `
	PhoneNumber *string `json:"phone_number"`
	Email       string  `json:"email"`
	Biography   *string `json:"biography"`
	Education   *string `json:"education"`
}

type UpdateAuthorDto struct {
	ID          string  `json:"id"`
	FirstName   string  `json:"first_name"`
	LastName    *string `json:"last_name" `
	PhoneNumber *string `json:"phone_number"`
	Email       string  `json:"email"`
	Biography   *string `json:"biography"`
	Education   *string `json:"education"`
}
