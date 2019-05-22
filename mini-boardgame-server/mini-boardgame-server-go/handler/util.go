package handler

type J map[string]interface{}

func M(msg string) interface{} {
	return J{
		"message": msg,
	}
}
