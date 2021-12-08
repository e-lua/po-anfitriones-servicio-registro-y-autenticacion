package repositories

import (
	"encoding/json"
	"strconv"

	models "github.com/Aphofisis/po-anfitrion-servicio-registro-y-autenticacion/models"
	"github.com/gomodule/redigo/redis"
)

func Re_Get_Phone(phoneregister int, idcountry int) (models.Re_SetGetCode, error) {

	var code models.Re_SetGetCode

	reply, err := redis.String(models.RedisCN.Get().Do("GET", strconv.Itoa(phoneregister)+strconv.Itoa(idcountry)))

	if err != nil {
		models.RedisCN.Get().Close()
		return code, err
	}

	err = json.Unmarshal([]byte(reply), &code)

	if err != nil {
		models.RedisCN.Get().Close()
		return code, err
	}

	models.RedisCN.Get().Close()
	return code, nil
}
