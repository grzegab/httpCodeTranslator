package httpCodeTranslator

type Translator struct {
	httpCode    int
	grpcCode    int
	grpcMessage string
	httpMessage string
}

func (t *Translator) grpcToHttp() (int, string) {
	t.httpCode, t.httpMessage = grpcDictionary(t.grpcCode)

	return t.httpCode, t.httpMessage
}

func (t *Translator) httpToGRPC() (int, string) {
	t.grpcCode, t.grpcMessage = httpDictionary(t.httpCode)

	return t.grpcCode, t.grpcMessage
}

func httpDictionary(code int) (int, string) {
	switch code {
	case 200:
		return 0, "OK"
	case 400:
		return 3, "INVALID_ARGUMENT"
	case 401:
		return 16, "UNAUTHENTICATED"
	case 403:
		return 7, "PERMISSION_DENIED"
	case 404:
		return 5, "NOT_FOUND"
	case 409:
		return 6, "ALREADY_EXISTS"
	case 429:
		return 8, "RESOURCE_EXHAUSTED"
	case 499:
		return 1, "CANCELLED"
	case 500:
		return 13, "INTERNAL"
	case 501:
		return 12, "UNIMPLEMENTED"
	case 503:
		return 14, "UNAVAILABLE"
	case 504:
		return 4, "DEADLINE_EXCEEDED"
	default:
		return 2, "UNKNOWN"
	}
}

func grpcDictionary(code int) (int, string) {
	switch code {
	case 0:
		return 200, "OK"
	case 1:
		return 499, "Client Closed Request"
	case 2:
		return 500, "Internal Server Error"
	case 3:
		return 400, "Bad Request"
	case 4:
		return 504, "Gateway Timeout"
	case 5:
		return 404, "Not Found"
	case 6:
		return 409, "Conflict"
	case 7:
		return 403, "Forbidden"
	case 8:
		return 429, "Too Many Requests"
	case 9:
		return 400, "Bad Request"
	case 10:
		return 409, "Conflict"
	case 11:
		return 400, "Bad Request"
	case 12:
		return 501, "Not Implemented"
	case 13:
		return 500, "Internal Server Error"
	case 14:
		return 503, "Service Unavailable"
	case 15:
		return 500, "Internal Server Error"
	case 16:
		return 401, "Unauthorized"
	default:
		return 500, "Internal Server Error"
	}
}
