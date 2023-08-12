package request

type CreatUserRequest struct {
	Nama     string `json:"nama"  validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
}

type UpdateUserRequest struct {
	Nama  string `json:"nama"`
	Email string `json:"email"`
}
