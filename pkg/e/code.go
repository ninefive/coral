package e

var (
	OK                  = &Errno{Code: 200, Message: "OK"}
	InternalServerError = &Errno{Code: 500, Message: "Internal server error"}

	ErrBind       = &Errno{Code: 200001, Message: "Error occurred while binding the request body to the struct."}
	ErrValidation = &Errno{Code: 200002, Message: "Validation failed."}
	ErrDatabase   = &Errno{Code: 200003, Message: "Database error."}
	ErrToken      = &Errno{Code: 200004, Message: "Error occurred while signing the JSON web token."}

	ErrEncrypt           = &Errno{Code: 100001, Message: "Error occurred while encrypting the user password."}
	ErrTokenInvalid      = &Errno{Code: 100002, Message: "The token was invalid."}
	ErrPasswordIncorrect = &Errno{Code: 100003, Message: "The password was incorrect."}
	ErrUserNotFound      = &Errno{Code: 100004, Message: "The user was not found."}
	ErrMissingHeader     = &Errno{Code: 100005, Message: "The length of the `Authorization` header is zero."}
)
