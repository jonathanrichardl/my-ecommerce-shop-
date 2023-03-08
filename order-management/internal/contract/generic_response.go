package contract

type GenericResponse struct {
	Data  interface{} `json:"data"`
	Error Error       `json:"error"`
}

func BuildResponse(data interface{}) *GenericResponse {
	return &GenericResponse{
		Data: data,
	}
}

func BuildErrorResponse(errorMessage string) *GenericResponse {
	return &GenericResponse{
		Error: Error{ErrorMessage: errorMessage},
	}
}
