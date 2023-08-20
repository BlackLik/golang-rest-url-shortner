package schema

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// GetSuccess200Response returns an ErrorResponse with a code of 200 and a message of "OK".
//
// No parameters.
// Returns an ErrorResponse.
func GetSuccess200Response() Response {
	return Response{
		Code:    200,
		Message: "OK",
	}
}

// GetError400Response returns a Response object with a code of 400 and a "Bad Request" message.
//
// No parameters.
// Returns a Response object.
func GetError400Response() Response {
	return Response{
		Code:    400,
		Message: "Bad Request",
	}
}

// GetError401Response generates a 401 Unauthorized response.
//
// Returns:
//
//	Response: The generated response object.
func GetError401Response() Response {
	return Response{
		Code:    401,
		Message: "Unauthorized",
	}
}

// GetError404Response returns a Response object with a 404 status code and a "Not Found" message.
//
// No parameters.
// Returns a Response object.
func GetError404Response() Response {
	return Response{
		Code:    404,
		Message: "Not Found",
	}
}

// GetError500Response returns a Response with a code of 500 and a message of "Internal Server Error".
//
// No parameters.
// Returns a Response.
func GetError500Response() Response {
	return Response{
		Code:    500,
		Message: "Internal Server Error",
	}
}
