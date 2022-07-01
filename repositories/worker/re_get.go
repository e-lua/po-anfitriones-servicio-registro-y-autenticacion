package repositories

import (
	"math/rand"
	"strconv"

	models "github.com/Aphofisis/po-anfitrion-servicio-registro-y-autenticacion/models"
	"github.com/gomodule/redigo/redis"
)

func Re_Get_Id(idworker int, idcountry int, idbusiness int) (string, error) {

	random_num := rand.Intn(10)
	var reply string

	if random_num > 5 {
		//CONNECTION TO MASTER
		reply_master, err_master := redis.String(models.RedisCN.Get().Do("GET", strconv.Itoa(idworker)+strconv.Itoa(idcountry)+strconv.Itoa(idbusiness)))
		if err_master != nil {
			return reply_master, err_master
		}
		reply = reply_master
	} else {
		//CONNECTION TO SLAVE
		reply_slave, err_slave := redis.String(models.RedisCN_Slave.Get().Do("GET", strconv.Itoa(idworker)+strconv.Itoa(idcountry)+strconv.Itoa(idbusiness)))
		if err_slave != nil {
			return reply_slave, err_slave
		}
		reply = reply_slave
	}

	return reply, nil
}

func Re_Get_Email(idworker int, sessioncode int, idrol int) (string, error) {

	random_num := rand.Intn(10)
	var reply string

	if random_num > 5 {
		//CONNECTION TO MASTER
		reply_master, err_master := redis.String(models.RedisCN.Get().Do("GET", strconv.Itoa(idworker)+strconv.Itoa(idrol)))
		if err_master != nil {
			return reply_master, err_master
		}
		reply = reply_master
	} else {
		//CONNECTION TO SLAVE
		reply_slave, err_slave := redis.String(models.RedisCN_Slave.Get().Do("GET", strconv.Itoa(idworker)+strconv.Itoa(idrol)))
		if err_slave != nil {
			return reply_slave, err_slave
		}
		reply = reply_slave
	}

	return reply, nil
}
