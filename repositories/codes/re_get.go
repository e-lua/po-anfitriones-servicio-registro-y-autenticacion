package repositories

import (
	"encoding/json"
	"math/rand"
	"strconv"

	models "github.com/Aphofisis/po-anfitrion-servicio-registro-y-autenticacion/models"
	"github.com/gomodule/redigo/redis"
)

func Re_Get_Phone(phoneregister int, idcountry int) (models.Re_SetGetCode, error) {

	var code models.Re_SetGetCode
	var reply string
	var err error

	random := rand.Intn(4)
	if random%2 == 0 {
		reply, err = redis.String(models.RedisCN.Get().Do("GET", strconv.Itoa(phoneregister)+strconv.Itoa(idcountry)))
	} else {
		reply, err = redis.String(models.RedisCN_Slave.Get().Do("GET", strconv.Itoa(phoneregister)+strconv.Itoa(idcountry)))
	}

	if err != nil {
		return code, err
	}

	err = json.Unmarshal([]byte(reply), &code)

	if err != nil {
		return code, err
	}

	return code, nil
}
