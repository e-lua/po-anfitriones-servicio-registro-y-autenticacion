package repositories

import (
	"strconv"

	models "github.com/Aphofisis/po-anfitrion-servicio-registro-y-autenticacion/models"
	"github.com/gomodule/redigo/redis"
)

func Re_Get_Id(idbusiness int, idcountry int) (string, error) {

	reply, err := redis.String(models.RedisCN.Get().Do("GET", strconv.Itoa(idbusiness)+strconv.Itoa(idcountry)))

	if err != nil {
		models.RedisCN.Get().Close()
		return reply, err
	}

	models.RedisCN.Get().Close()
	return reply, nil
}
