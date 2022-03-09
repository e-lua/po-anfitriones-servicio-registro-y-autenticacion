package export

import (
	"strconv"

	"github.com/labstack/echo/v4"
	//MDOELS
)

var ExportRouter *exportRouter

type exportRouter struct {
}

func (er *exportRouter) ExportIDDevice(c echo.Context) error {

	idbusiness := c.Param("idbusiness")
	idbusiness_int, _ := strconv.Atoi(idbusiness)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := ExportIDDevice_Service(idbusiness_int)
	results := Response_ListString{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}
