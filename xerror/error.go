package xerror

/*
	Error reporting library.

	- Separates public and private information on errors. Public may be reported back to client (API caller),
		private is intended for internal use only (logs)

	- Abstracts from transport layer. Instead of providing transport-specific codes, user provides basic classification.
		Later transport code may be derived from classification.

	New...-methods create instance from text
	Wrap...-methods create instance from `error` object
*/

const (
	// Avoid use net/http; no need for additional dependency for the sake of several well-known constants
	httpBadRequest          = 400
	httpUnauthorized        = 401
	httpForbidden           = 403
	httpInternalServerError = 500

	errCodeJsonRpcInvalidParams      = -32602
	errCodeJsonRpcInternalError      = -32603
	errCodeJsonRpcCustomUnauthorized = -32001
	errCodeJsonRpcCustomForbidden    = -32003
)

type IError interface {
	// Message to be provided to end-user
	PublicMessage() string
	// Optional text for internal logs with explanation
	PrivateDetails() string
	// Status to be reported when called as http endpoint
	HttpStatus() int
	// JSON RPC error code
	JsonRpcErrorCode() int
}

type _Error struct {
	publicMessage    string
	privateDetails   string
	httpStatus       int
	jsonRpcErrorCode int
}

func (e *_Error) PublicMessage() string  { return e.publicMessage }
func (e *_Error) PrivateDetails() string { return e.privateDetails }
func (e *_Error) HttpStatus() int        { return e.httpStatus }
func (e *_Error) JsonRpcErrorCode() int  { return e.jsonRpcErrorCode }

// Provides public message to avoind unintentional disclosure of private information
func (e *_Error) Error() string { return e.publicMessage }

func _New(httpStatus, jsonRpcErrorCode int, publicMessage, privateDetails string) IError {
	return &_Error{
		publicMessage:    publicMessage,
		privateDetails:   privateDetails,
		httpStatus:       httpStatus,
		jsonRpcErrorCode: jsonRpcErrorCode,
	}
}

func NewBadRequest(publicMessage string) IError {
	return _New(httpBadRequest, errCodeJsonRpcInvalidParams, publicMessage, "")
}

func NewBadRequestDetailed(publicMessage, privateDetails string) IError {
	return _New(httpBadRequest, errCodeJsonRpcInvalidParams, publicMessage, privateDetails)
}

func WrapBadRequest(err error) IError {
	return _New(httpBadRequest, errCodeJsonRpcInvalidParams, err.Error(), "")
}

func NewUnauthorized(publicMessage string) IError {
	return _New(httpUnauthorized, errCodeJsonRpcCustomUnauthorized, publicMessage, "")
}

func NewForbidden(publicMessage string) IError {
	return _New(httpForbidden, errCodeJsonRpcCustomForbidden, publicMessage, "")
}

func NewForbiddenDetailed(publicMessage, privateDetails string) IError {
	return _New(httpForbidden, errCodeJsonRpcCustomForbidden, publicMessage, privateDetails)
}

func NewCustom(httpStatus, jsonRpcErrorCode int, publicMessage string) IError {
	return _New(httpStatus, jsonRpcErrorCode, publicMessage, "")
}

func WrapFailure(err error) IError {
	return _New(httpInternalServerError, errCodeJsonRpcInternalError, err.Error(), "")
}

func NewFailure(publicMessage string) IError {
	return _New(httpInternalServerError, errCodeJsonRpcInternalError, publicMessage, "")
}

func NewFailureDetailed(publicMessage, privateDetails string) IError {
	return _New(httpInternalServerError, errCodeJsonRpcInternalError, publicMessage, privateDetails)
}

func WrapFailureDetailed(publicMessage string, err error) IError {
	return _New(httpInternalServerError, errCodeJsonRpcInternalError, publicMessage, err.Error())
}
