package messages

func Resp(val any, op string, respCode int) map[string]any {
	var payload = map[string]any{
		"Data":   val,
		"Op":     op,
		"Status": respCode,
	}
	return payload
}
