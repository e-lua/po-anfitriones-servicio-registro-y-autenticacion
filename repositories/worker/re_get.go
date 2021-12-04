package repositories

import (
	models "github.com/Aphofisis/po-anfitrion-servicio-registro-y-autenticacion/models"
	"github.com/gomodule/redigo/redis"
)

func Re_Get_Id(idbusiness int) (string, error) {

	reply, err := redis.String(models.RedisCN.Do("GET", idbusiness))

	if err != nil {
		return reply, err
	}

	if err != nil {
		return reply, err
	}
	return reply, nil
}
