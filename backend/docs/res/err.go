package res

type APIError struct {
	Status string `json:"status" example:"error"`
	Payload errorMessage `json:"payload"`
}

type errorMessage struct {
	Message string `json:"message" example:"internal server error"`
}

type AuthError struct {
	Status string `json:"status" example:"error"`
	Payload authErrorMessage `json:"payload"`
}

type authErrorMessage struct {
	Message string `json:"message" example:"Unauthorized"`
}
