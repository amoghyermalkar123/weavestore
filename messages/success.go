package messages

func Resp(val any, op string, respCode int) map[string]interface{} {
	var Success = map[string]any{
		"Data":   val,
		"Op":     op,
		"Status": respCode,
	}
	return Success
}
