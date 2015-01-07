package quest

// Methods.
//  HTTP method definitions.
//  See http://tools.ietf.org/html/rfc7231#section-4.3
const (
	OPTIONS = "OPTIONS"
	GET     = "GET"
	HEAD    = "HEAD"
	POST    = "POST"
	PUT     = "PUT"
	PATCH   = "PATCH"
	DELETE  = "DELETE"
	TRACE   = "TRACE"
	CONNECT = "CONNECT"
)

// Method map.
var Methods = map[string]string{
	OPTIONS: OPTIONS,
	GET:     GET,
	HEAD:    HEAD,
	POST:    POST,
	PUT:     PUT,
	PATCH:   PATCH,
	DELETE:  DELETE,
	TRACE:   TRACE,
	CONNECT: CONNECT,
}
