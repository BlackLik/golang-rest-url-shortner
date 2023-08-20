package schema

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func GetSuccess200Response() ErrorResponse {
	return ErrorResponse{
		Code:    200,
		Message: "OK",
	}
}

func GetError400Response() ErrorResponse {
	return ErrorResponse{
		Code:    400,
		Message: "Bad Request",
	}
}

func GetError404Response() ErrorResponse {
	return ErrorResponse{
		Code:    404,
		Message: "Not Found",
	}
}

func GetError500Response() ErrorResponse {
	return ErrorResponse{
		Code:    500,
		Message: "Internal Server Error",
	}
}
