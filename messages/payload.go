package messages

// Resp creates json response payloads for the http api
func Resp(val any, respCode int) map[string]any {
	var payload = map[string]any{
		"Data":   val,
		"Status": respCode,
	}
	return payload
}
