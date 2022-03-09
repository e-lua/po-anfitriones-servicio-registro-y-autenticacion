package export

import (

	//MDOELS

	worker_repository "github.com/Aphofisis/po-anfitrion-servicio-registro-y-autenticacion/repositories/worker"
)

func ExportIDDevice_Service(idbusines int) (int, bool, string, []string) {

	//Enviamos la variable instanciada al repository
	list_string, error_list_string := worker_repository.Pg_Find_IDDevice(idbusines)
	if error_list_string != nil {
		return 500, true, "Error interno en el servidor al intentar listar todos los ID de dispostivos, detalle: " + error_list_string.Error(), list_string
	}

	return 201, false, "", list_string
}
