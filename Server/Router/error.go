package Router

type HTTPError struct {
	Message   string
	ErrorCode int
}

const (
	ErrorTwitsNotFound          = "ErrorTwitsNotFound"
	ErrorDangerousUsersNotFound = "ErrorDangerousUsersNotFound"
	BadRequest                  = "BadRequest"
)

var customHttpErrors = map[string]HTTPError{
	ErrorTwitsNotFound:          {Message: "Twits Not found", ErrorCode: 404},
	ErrorDangerousUsersNotFound: {Message: "Dangerous Users Not Found", ErrorCode: 404},
	BadRequest:                  {Message: "Bad Request", ErrorCode: 400},
}

func GetHTTPError(errorName string) HTTPError {
	if customError, ok := customHttpErrors[errorName]; ok {
		return customError
	}

	return HTTPError{Message: "Internal Server Error", ErrorCode: 500}
}
