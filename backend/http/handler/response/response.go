package response

func SuccessResult(result interface{}) map[string]interface{} {
	return map[string]interface{} {
		"status": "success",
		"payload": result,
	}
}

func ErrorResult(message string) map[string]interface{} {
	return map[string]interface{} {
		"status": "error",
		"payload": map[string]interface{} {
			"message": message,
		},
	}
}