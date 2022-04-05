package export

import (
	"github.com/labstack/echo/v4"
)

var ExportRouter *exportRouter

type exportRouter struct {
}

func (er *exportRouter) ExportIDDevice(c echo.Context) error {

	//Instanciamos una variable del modelo B_Name
	var export_notification Request_Export_Notifications

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&export_notification)
	if err != nil {
		results := Response_ListString{Error: true, DataError: "Se debe enviar todos los negocios, revise la estructura o los valores", Data: nil}
		return c.JSON(400, results)
	}

	//Validamos los valores enviados
	if export_notification.Type != 1 && export_notification.Type != 2 && export_notification.Type != 3 {
		results := Response_ListString{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio", Data: nil}
		return c.JSON(403, results)
	}

	status, boolerror, dataerror, data := ExportIDDevice_Service(export_notification.Idbusiness, export_notification.Type, export_notification.ArrayBusinesses)
	results := Response_ListString{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}
