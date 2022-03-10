package export

import (

	//MDOELS

	worker_repository "github.com/Aphofisis/po-anfitrion-servicio-registro-y-autenticacion/repositories/worker"
)

func ExportIDDevice_Service(idbusines int, type_int int) (int, bool, string, []string) {

	var list_string []string
	var error_find error

	if type_int == 1 {
		//Enviamos la variable instanciada al repository
		list_string, error_find = worker_repository.Pg_Find_IDDevice(idbusines)
		if error_find != nil {
			return 500, true, "Error interno en el servidor al intentar listar todos los ID de dispostivos, detalle: " + error_find.Error(), list_string
		}
	}

	if type_int == 2 {
		//Enviamos la variable instanciada al repository
		list_string, error_find = worker_repository.Pg_Find_IDDevice_All()
		if error_find != nil {
			return 500, true, "Error interno en el servidor al intentar listar todos los ID de dispostivos, detalle: " + error_find.Error(), list_string
		}
	}

	return 201, false, "", list_string
}
