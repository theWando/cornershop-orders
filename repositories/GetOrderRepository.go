package repositories

import (
	"fmt"
	"github.com/theWando/conershopOrders/rest/client"
	"os"
)

const url = "https://cornershopapp.com/api/v3/orders/"

func GetOrder(id string) (map[string]interface{}, error) {
	var headers = make(map[string]interface{}, 2)
	headers["cookie"] = os.Getenv("CORNERSHOP_COOKIE")
	headers["accept"] = "application/json"
	return client.Get(fmt.Sprint(url, id), headers)
}
