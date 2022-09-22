package repositories

import (
	"strconv"

	models "github.com/Aphofisis/po-anfitrion-servicio-registro-y-autenticacion/models"
	"github.com/gomodule/redigo/redis"
)

func Re_Get_Request(phoneregister int, idcountry int) (int, error) {

	reply, _ := redis.String(models.RedisCN.Get().Do("GET", strconv.Itoa(phoneregister)+strconv.Itoa(idcountry)+"REQUEST"))
	quantity_int, _ := strconv.Atoi(reply)

	return quantity_int, nil
}
