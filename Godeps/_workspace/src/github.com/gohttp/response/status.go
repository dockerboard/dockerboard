package response

import "net/http"
import "fmt"

// respond with the given message.
func text(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(code)
	fmt.Fprintln(w, msg)
}

// write `msg` to the response writer, currently only a single argument is supported.
func write(w http.ResponseWriter, code int, msg []interface{}) {
	if len(msg) == 0 {
		text(w, code, http.StatusText(code))
		return
	}

	switch msg[0].(type) {
	case string:
		text(w, code, msg[0].(string))
	default:
		JSON(w, msg[0], code)
	}
}

// Continue response.
func Continue(w http.ResponseWriter, msg ...interface{}) {
	write(w, http.StatusContinue, msg)
}

// SwitchingProtocols response.
func SwitchingProtocols(w http.ResponseWriter, msg ...interface{}) {
	write(w, http.StatusSwitchingProtocols, msg)
}

// OK response.
func OK(w http.ResponseWriter, msg ...interface{}) {
	write(w, http.StatusOK, msg)
}

// Created response.
func Created(w http.ResponseWriter, msg ...interface{}) {
	write(w, http.StatusCreated, msg)
}

// Accepted response.
func Accepted(w http.ResponseWriter, msg ...interface{}) {
	write(w, http.StatusAccepted, msg)
}

// NonAuthoritativeInfo response.
func NonAuthoritativeInfo(w http.ResponseWriter, msg ...interface{}) {
	write(w, http.StatusNonAuthoritativeInfo, msg)
}

// NoContent response.
func NoContent(w http.ResponseWriter, msg ...interface{}) {
	write(w, http.StatusNoContent, msg)
}

// ResetContent response.
func ResetContent(w http.ResponseWriter, msg ...interface{}) {
	write(w, http.StatusResetContent, msg)
}

// PartialContent response.
func PartialContent(w http.ResponseWriter, msg ...interface{}) {
	write(w, http.StatusPartialContent, msg)
}

// MultipleChoices response.
func MultipleChoices(w http.ResponseWriter, msg ...interface{}) {
	write(w, http.StatusMultipleChoices, msg)
}

// MovedPermanently response.
func MovedPermanently(w http.ResponseWriter, msg ...interface{}) {
	write(w, http.StatusMovedPermanently, msg)
}

// Found response.
func Found(w http.ResponseWriter, msg ...interface{}) {
	write(w, http.StatusFound, msg)
}

// SeeOther response.
func SeeOther(w http.ResponseWriter, msg ...interface{}) {
	write(w, http.StatusSeeOther, msg)
}

// NotModified response.
func NotModified(w http.ResponseWriter, msg ...interface{}) {
	write(w, http.StatusNotModified, msg)
}

// UseProxy response.
func UseProxy(w http.ResponseWriter, msg ...interface{}) {
	write(w, http.StatusUseProxy, msg)
}

// TemporaryRedirect response.
func TemporaryRedirect(w http.ResponseWriter, msg ...interface{}) {
	write(w, http.StatusTemporaryRedirect, msg)
}

// BadRequest response.
func BadRequest(w http.ResponseWriter, msg ...interface{}) {
	write(w, http.StatusBadRequest, msg)
}

// Unauthorized response.
func Unauthorized(w http.ResponseWriter, msg ...interface{}) {
	write(w, http.StatusUnauthorized, msg)
}

// PaymentRequired response.
func PaymentRequired(w http.ResponseWriter, msg ...interface{}) {
	write(w, http.StatusPaymentRequired, msg)
}

// Forbidden response.
func Forbidden(w http.ResponseWriter, msg ...interface{}) {
	write(w, http.StatusForbidden, msg)
}

// NotFound response.
func NotFound(w http.ResponseWriter, msg ...interface{}) {
	write(w, http.StatusNotFound, msg)
}

// MethodNotAllowed response.
func MethodNotAllowed(w http.ResponseWriter, msg ...interface{}) {
	write(w, http.StatusMethodNotAllowed, msg)
}

// NotAcceptable response.
func NotAcceptable(w http.ResponseWriter, msg ...interface{}) {
	write(w, http.StatusNotAcceptable, msg)
}

// ProxyAuthRequired response.
func ProxyAuthRequired(w http.ResponseWriter, msg ...interface{}) {
	write(w, http.StatusProxyAuthRequired, msg)
}

// RequestTimeout response.
func RequestTimeout(w http.ResponseWriter, msg ...interface{}) {
	write(w, http.StatusRequestTimeout, msg)
}

// Conflict response.
func Conflict(w http.ResponseWriter, msg ...interface{}) {
	write(w, http.StatusConflict, msg)
}

// Gone response.
func Gone(w http.ResponseWriter, msg ...interface{}) {
	write(w, http.StatusGone, msg)
}

// LengthRequired response.
func LengthRequired(w http.ResponseWriter, msg ...interface{}) {
	write(w, http.StatusLengthRequired, msg)
}

// PreconditionFailed response.
func PreconditionFailed(w http.ResponseWriter, msg ...interface{}) {
	write(w, http.StatusPreconditionFailed, msg)
}

// RequestEntityTooLarge response.
func RequestEntityTooLarge(w http.ResponseWriter, msg ...interface{}) {
	write(w, http.StatusRequestEntityTooLarge, msg)
}

// RequestURITooLong response.
func RequestURITooLong(w http.ResponseWriter, msg ...interface{}) {
	write(w, http.StatusRequestURITooLong, msg)
}

// UnsupportedMediaType response.
func UnsupportedMediaType(w http.ResponseWriter, msg ...interface{}) {
	write(w, http.StatusUnsupportedMediaType, msg)
}

// RequestedRangeNotSatisfiable response.
func RequestedRangeNotSatisfiable(w http.ResponseWriter, msg ...interface{}) {
	write(w, http.StatusRequestedRangeNotSatisfiable, msg)
}

// ExpectationFailed response.
func ExpectationFailed(w http.ResponseWriter, msg ...interface{}) {
	write(w, http.StatusExpectationFailed, msg)
}

// Teapot response.
func Teapot(w http.ResponseWriter, msg ...interface{}) {
	write(w, http.StatusTeapot, msg)
}

// InternalServerError response.
func InternalServerError(w http.ResponseWriter, msg ...interface{}) {
	write(w, http.StatusInternalServerError, msg)
}

// NotImplemented response.
func NotImplemented(w http.ResponseWriter, msg ...interface{}) {
	write(w, http.StatusNotImplemented, msg)
}

// BadGateway response.
func BadGateway(w http.ResponseWriter, msg ...interface{}) {
	write(w, http.StatusBadGateway, msg)
}

// ServiceUnavailable response.
func ServiceUnavailable(w http.ResponseWriter, msg ...interface{}) {
	write(w, http.StatusServiceUnavailable, msg)
}

// GatewayTimeout response.
func GatewayTimeout(w http.ResponseWriter, msg ...interface{}) {
	write(w, http.StatusGatewayTimeout, msg)
}

// HTTPVersionNotSupported response.
func HTTPVersionNotSupported(w http.ResponseWriter, msg ...interface{}) {
	write(w, http.StatusHTTPVersionNotSupported, msg)
}
