package messages

func Resp(val any, respCode int) map[string]any {
	var payload = map[string]any{
		"Data":   val,
		"Status": respCode,
	}
	return payload
}
