package repositories

import (
	"math/rand"
	"strconv"

	models "github.com/Aphofisis/po-anfitrion-servicio-registro-y-autenticacion/models"
	"github.com/gomodule/redigo/redis"
)

func Re_Get_Id(idworker int, idcountry int, idbusiness int) (string, error) {

	var reply string
	var err error

	random := rand.Intn(4)
	if random%2 == 0 {
		reply, err = redis.String(models.RedisCN.Get().Do("GET", strconv.Itoa(idworker)+strconv.Itoa(idcountry)+strconv.Itoa(idbusiness)))
	} else {
		reply, err = redis.String(models.RedisCN_Slave.Get().Do("GET", strconv.Itoa(idworker)+strconv.Itoa(idcountry)+strconv.Itoa(idbusiness)))
	}

	if err != nil {
		return reply, err
	}

	return reply, nil
}

func Re_Get_Email(idworker int, sessioncode int, idrol int) (string, error) {

	var reply string
	var err error

	random := rand.Intn(4)
	if random%2 == 0 {
		reply, err = redis.String(models.RedisCN.Get().Do("GET", strconv.Itoa(idworker)+strconv.Itoa(idrol)))
	} else {
		reply, err = redis.String(models.RedisCN_Slave.Get().Do("GET", strconv.Itoa(idworker)+strconv.Itoa(idrol)))
	}

	if err != nil {
		return reply, err
	}

	return reply, nil
}
