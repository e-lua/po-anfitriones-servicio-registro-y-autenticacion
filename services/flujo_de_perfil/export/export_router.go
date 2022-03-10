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

	idbusiness_string := c.Request().URL.Query().Get("idbusiness")
	idbusiness, _ := strconv.Atoi(idbusiness_string)

	type_string := c.Request().URL.Query().Get("type")
	type_int, _ := strconv.Atoi(type_string)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := ExportIDDevice_Service(idbusiness, type_int)
	results := Response_ListString{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}
