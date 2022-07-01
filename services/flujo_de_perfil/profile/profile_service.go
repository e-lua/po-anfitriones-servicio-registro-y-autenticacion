package profile

import (

	//MDOELS

	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"

	//REPOSITORIES
	"github.com/Aphofisis/po-anfitrion-servicio-registro-y-autenticacion/models"
	worker_repository "github.com/Aphofisis/po-anfitrion-servicio-registro-y-autenticacion/repositories/worker"
)

//FUNCIONES PRIVADAS

func compareToken(inputpassword_comensal_founded string, inputpassword string) error {

	//brypt trabaja con slices de bytes - Esta password no esta encriptada
	passwordBytes := []byte(inputpassword)

	//Password se encripta
	passwordBD := []byte(inputpassword_comensal_founded)

	//Comparamos si el hash encriptado es el password que se escribio en el Login
	error_compare_hash := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if error_compare_hash != nil {
		return error_compare_hash
	}

	return nil
}

func encrypt(input string) (string, error) {
	cost := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(input), cost)
	return string(bytes), err
}

/*-------------------------------FUNCIONES PUBLICAS-------------------------------*/

func UpdateNameLastNameEmail_Service(input_business Entry_Profile, input_idworker int) (int, bool, string, string) {

	//Enviamos la variable instanciada al repository
	error_update_password := worker_repository.Pg_Update_NameLastNameEmail(input_business.Name, input_business.LastName, input_business.Email, input_idworker)
	if error_update_password != nil {
		return 500, true, "Error interno en el servidor al intentar actualizar la contraseña, detalle: " + error_update_password.Error(), ""
	}

	return 201, false, "", "Nombre y apellidos actualizados correctamente"
}

func UpdatePassword_Service(input_entrydata EntryData_Password, idbusiness int, idcountry int) (int, bool, string, string) {

	//Buscamos la existencia del registro en Pg
	pass, idworker, error_findcomensal := worker_repository.Pg_FindPassword_ById(idbusiness)
	if error_findcomensal != nil {
		return 500, true, "Error en el servidor interno al intentar buscar el comensal, detalle: " + error_findcomensal.Error(), ""
	}
	if pass == "" {
		return 404, true, "Este anfitrion no se encuentra registrado", ""
	}

	//Intentamos el login
	error_compareToken := compareToken(pass, input_entrydata.OldPassword)
	if error_compareToken != nil {
		return 403, true, "Contraseña incorrecta, detalle: " + error_compareToken.Error(), ""
	}

	//Creamos un nuevo  codigo de sesion
	hour, minute, sec := time.Now().Clock()
	new_session_code := minute*100 + sec + hour + 1111 + rand.Intn(7483647)

	//Encriptar password
	encrypted_pass, _ := encrypt(input_entrydata.NewPassword)

	//Enviamos la variable instanciada al repository
	error_update_password := worker_repository.Pg_Update_Password(encrypted_pass, idbusiness)
	if error_update_password != nil {
		return 500, true, "Error interno en el servidor al intentar actualizar la contraseña, detalle: " + error_update_password.Error(), ""
	}

	//Registramos en Redis
	_, err_add_re := worker_repository.Re_Set_ID(idworker, idcountry, new_session_code, idbusiness)
	if err_add_re != nil {
		return 500, true, "Error en el servidor interno al intentar registrar el código en cache, detalle: " + err_add_re.Error(), ""
	}

	return 201, false, "", "Contraseña actualizada correctamente"
}

func DeleteAnfitrion_Service(input_idworker int) (int, bool, string, string) {

	//Validamos que no se tengan colaboradores asociados
	_, quantity, error_find_workers := worker_repository.Pg_Find_SubWorkers(input_idworker)
	if error_find_workers != nil {
		return 500, true, "Error interno en el servidor al intentar buscar los colaboradores, detalle: " + error_find_workers.Error(), ""
	}
	if quantity > 0 {
		return 403, true, "No se puede eliminar esta cuenta, ya que cuenta con colaboradores activos", ""
	}

	//Enviamos la variable instanciada al repository
	error_update_password := worker_repository.Pg_Update_IsDeleted(input_idworker)
	if error_update_password != nil {
		return 500, true, "Error interno en el servidor al intentar eliminar al anfitrion, detalle: " + error_update_password.Error(), ""
	}

	return 201, false, "", "Eliminado correctamente"
}

func DeleteColaborador_Service(input_idsubworker int, data_idrol int, data_idcountry int, data_idbusines int) (int, bool, string, string) {

	//Enviamos la variable instanciada al repository
	error_update_password := worker_repository.Pg_Delete_SubWorker(input_idsubworker)
	if error_update_password != nil {
		return 500, true, "Error interno en el servidor al intentar eliminar al anfitrion, detalle: " + error_update_password.Error(), ""
	}

	//Registramos en Redis
	err_add_re := worker_repository.Re_Set_Email(input_idsubworker, 12411451345, data_idrol)
	if err_add_re != nil {
		return 500, true, "Error en el servidor interno al intentar cambiar el codigo para evitar el ingreso del subworker eliminado, detalle: " + err_add_re.Error(), ""
	}

	return 201, false, "", "Colaborador eliminado correctamente"
}

func UpdateIDDevice_Service(input_idworker int, iddevice string) (int, bool, string, string) {

	//Enviamos la variable instanciada al repository
	error_update_iddevice := worker_repository.Pg_Update_IDDevice(input_idworker, iddevice)
	if error_update_iddevice != nil {
		return 500, true, "Error interno en el servidor al intentar actualziar el ID del dispositivo, detalle: " + error_update_iddevice.Error(), ""
	}

	return 201, false, "", "ID de dispositivo actualizado correctamente"
}

func GetColaborador_Service(input_idbusiness int) (int, bool, string, []models.Pg_SubWorker) {

	//Enviamos la variable instanciada al repository
	subworkers, _, error_update_password := worker_repository.Pg_Find_SubWorkers(input_idbusiness)
	if error_update_password != nil {
		return 500, true, "Error interno en el servidor al intentar eliminar al anfitrion, detalle: " + error_update_password.Error(), subworkers
	}

	return 201, false, "", subworkers
}

func GetEmail_Service(input_idbusiness int) (int, bool, string, string) {

	//Enviamos la variable instanciada al repository
	email_found, error_update_password := worker_repository.Pg_Find_Email(input_idbusiness)
	if error_update_password != nil {
		return 500, true, "Error interno en el servidor al intentar obtener el email del anfitrion, detalle: " + error_update_password.Error(), email_found
	}

	return 201, false, "", email_found
}

/*=======================================*/
/*===============VERSION 2===============*/
/*=======================================*/

func V2_GetColaborador_Service(input_idbusiness int) (int, bool, string, []models.V2_Pg_SubWorker) {

	//Enviamos la variable instanciada al repository
	subworkers, _, error_update_password := worker_repository.V2_Pg_Find_SubWorkers(input_idbusiness)
	if error_update_password != nil {
		return 500, true, "Error interno en el servidor al intentar obtener los datos del anfitrion, detalle: " + error_update_password.Error(), subworkers
	}

	return 201, false, "", subworkers
}

func V2_GetColaboradorToExport_Service(input_idsubworker int) (int, bool, string, models.V2_Pg_SubWorker) {

	//Enviamos la variable instanciada al repository
	subworker, error_update_password := worker_repository.V2_Pg_Find_SubWorkers_ToWorker(input_idsubworker)
	if error_update_password != nil {
		return 500, true, "Error interno en el servidor al intentar exportar los datos del anfitrion, detalle: " + error_update_password.Error(), subworker
	}

	return 201, false, "", subworker
}
