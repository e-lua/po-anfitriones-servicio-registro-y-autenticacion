package repositories

import (
	"encoding/json"

	models "github.com/Aphofisis/po-anfitrion-servicio-registro-y-autenticacion/models"
	"github.com/gomodule/redigo/redis"
)

func Re_Get_Phone(phoneregister int) (models.Re_SetGetCode, error) {

	var code models.Re_SetGetCode

	reply, err := redis.String(models.RedisCN.Do("GET", phoneregister))

	if err != nil {
		return code, err
	}

	err = json.Unmarshal([]byte(reply), &code)

	if err != nil {
		return code, err
	}
	return code, nil
}
