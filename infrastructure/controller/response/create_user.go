package response

type CreateUserResponse struct {
	Message string `json:"message"`
}

func NewCreateUserResponse(message string) CreateUserResponse {
	return CreateUserResponse{
		Message: message,
	}
}
