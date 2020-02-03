package repositories

import "os"

func DefaultHeaders() map[string]interface{} {
	var headers = make(map[string]interface{}, 2)
	headers["cookie"] = os.Getenv("CORNERSHOP_COOKIE")
	headers["accept"] = "application/json"
	return headers
}
