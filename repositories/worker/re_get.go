package repositories

import (
	"strconv"

	models "github.com/Aphofisis/po-anfitrion-servicio-registro-y-autenticacion/models"
	"github.com/gomodule/redigo/redis"
)

func Re_Get_Id(idworker int, idcountry int, idbusiness int) (string, error) {

	reply, err := redis.String(models.RedisCN.Get().Do("GET", strconv.Itoa(idworker)+strconv.Itoa(idcountry)+strconv.Itoa(idbusiness)))

	if err != nil {
		return reply, err
	}

	return reply, nil
}
