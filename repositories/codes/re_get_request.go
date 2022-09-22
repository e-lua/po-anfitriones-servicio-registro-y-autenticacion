package repositories

import (
	"encoding/json"
	"strconv"

	models "github.com/Aphofisis/po-anfitrion-servicio-registro-y-autenticacion/models"
	"github.com/gomodule/redigo/redis"
)

func Re_Get_Request(phoneregister int, idcountry int) (int, error) {

	var quantity_string string

	reply, _ := redis.String(models.RedisCN.Get().Do("GET", strconv.Itoa(phoneregister)+strconv.Itoa(idcountry)+"REQUEST"))

	err := json.Unmarshal([]byte(reply), &quantity_string)
	quantity_int, _ := strconv.Atoi(quantity_string)
	if err != nil {
		return quantity_int, err
	}

	return quantity_int, nil
}
