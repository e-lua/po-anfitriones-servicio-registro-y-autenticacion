package repositories

import (
	"strconv"

	models "github.com/Aphofisis/po-anfitrion-servicio-registro-y-autenticacion/models"
	"github.com/gomodule/redigo/redis"
)

func Re_Get_Id(idworker int, idcountry int, idbusiness int) (string, error) {
	//CONNECTION TO MASTER
	reply_master, err_master := redis.String(models.RedisCN.Get().Do("GET", strconv.Itoa(idworker)+strconv.Itoa(idcountry)+strconv.Itoa(idbusiness)))
	if err_master != nil {
		return reply_master, err_master
	}

	return reply_master, nil
}

func Re_Get_Email(idworker int, sessioncode int, idrol int) (string, error) {

	reply_master, err_master := redis.String(models.RedisCN.Get().Do("GET", strconv.Itoa(idworker)+strconv.Itoa(idrol)))
	if err_master != nil {
		return reply_master, err_master
	}

	return reply_master, nil
}
