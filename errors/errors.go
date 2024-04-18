package errors

type ErrorData struct {
	StatusCode  int
	Message     string
	Description string
}

var NotFound = ErrorData{
	StatusCode: 404,
	Message:    "Page Not Found",
}

var MethodNotAllowed = ErrorData{
	StatusCode: 405,
	Message:    "Method Not Allowed",
}

func BadRequest(message string) ErrorData {
	return ErrorData{
		StatusCode:  400,
		Message:     "Bad Request",
		Description: message,
	}
}

func InternalServer(message string) ErrorData {
    return ErrorData{
        StatusCode:  500,
        Message:     "Internal Server Error",
        Description: message,
    }
}
