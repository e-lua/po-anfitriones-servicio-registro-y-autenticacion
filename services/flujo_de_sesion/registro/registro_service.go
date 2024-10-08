package registro

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"

	//MDOELS
	models "github.com/Aphofisis/po-anfitrion-servicio-registro-y-autenticacion/models"
	"github.com/labstack/gommon/log"
	"github.com/streadway/amqp"

	//REPOSITORIES
	code_repository "github.com/Aphofisis/po-anfitrion-servicio-registro-y-autenticacion/repositories/codes"
	worker_repository "github.com/Aphofisis/po-anfitrion-servicio-registro-y-autenticacion/repositories/worker"

	//TWILIO
	twilio "github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

//FUNCIONES PRIVADAS
func encrypt(input string) (string, error) {
	cost := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(input), cost)
	return string(bytes), err
}

func notificacionWABA(country int, phone int, code string) (int, bool, string, string) {

	//Enviamos el codigo al anfitrion
	client := twilio.NewRestClientWithParams(twilio.RestClientParams{
		Username: "ACeb643456bb0e06813948315b95c3aa98",
		Password: "b6febb18bf85369763c4a303937137d9",
	})

	params := &openapi.CreateMessageParams{}
	params.SetTo("whatsapp:+" + strconv.Itoa(country) + strconv.Itoa(phone))
	params.SetFrom("whatsapp:+17816503313")
	params.SetBody(`Su codigo de Restoner es ` + code)

	_, err := client.ApiV2010.CreateMessage(params)
	if err != nil {
		return 200, false, err.Error(), ""
	}

	return 200, false, "", "Solicitud enviada correctamente"
}

func notificacion_registor_waba(idbusiness int, timezone string) (int, bool, string, string) {

	ahora := time.Now()
	time_zones, _ := strconv.Atoi(timezone)
	fecha := ahora.Add(time.Hour * time.Duration(time_zones))

	//Enviamos el codigo al anfitrion
	client := twilio.NewRestClientWithParams(twilio.RestClientParams{
		Username: "ACeb643456bb0e06813948315b95c3aa98",
		Password: "b6febb18bf85369763c4a303937137d9",
	})
	params := &openapi.CreateMessageParams{}
	params.SetTo("whatsapp:+51938488229")
	params.SetFrom("whatsapp:+17816503313")
	params.SetBody("*Restoner Bot:* Acaba de iniciar sesión en una cuenta de *" + " REGISTRO DEL NEGOCIO " + strconv.Itoa(idbusiness) + "* el " + fecha.Format("2006-01-02 3:4:5 pm") + ", si no fue usted, haga clic aquí https://youtu.be/JDcPVzSP8-M")

	_, err := client.ApiV2010.CreateMessage(params)
	if err != nil {
		return 200, false, "", "Solicitud enviada correctamente"
	}

	return 200, false, "", "Solicitud enviada correctamente"
}

//FUNCIONES PUBLICAS

func AvailableRegister_Service() (int, bool, string, bool) {

	available, err_update := worker_repository.Pg_Find_IfIsAvailable()
	if err_update != nil {
		return 500, true, "Error en el servidor interno al intentar actualizar la cantidad de codigos requeridos por este comensal, detalle: " + err_update.Error(), available
	}

	//Si todo ha ido bien envia un Status 200
	return 201, false, "", available
}

func SignUpNumber_Service(inputcode models.Re_SetGetCode) (int, bool, string, SignInFirstStep) {

	var phone_and_code SignInFirstStep

	/*---------------------------VALIDAMOS QUE NO ESTEN SPAMEANDO CODIGOS---------------------------*/
	quantity_request, error_request := code_repository.Re_Get_Request(inputcode.PhoneRegister_Key, inputcode.Country)
	if error_request != nil {
		/*---------------------------GUARDAMOS LOS NUMEROS QUE PIDEN CODIGO DIARIO---------------------------*/
		log.Print("Error al obtener los request ", error_request.Error())
		error_set_quantity := code_repository.Re_Set_Request(inputcode.PhoneRegister_Key, inputcode.Country, 0)
		if error_set_quantity != nil {
			return 500, true, "Error en el servidor interno al intentar actualizar la cantidad de codigos soliticidos, detalle: " + error_set_quantity.Error(), phone_and_code
		}
		/*---------------------------------------------------------------------------------------------------*/
	}
	if quantity_request > 2 {
		return 406, true, "Limite diario excedido", phone_and_code
	}
	/*---------------------------------------------------------------------------------------------*/

	if inputcode.Country != 51 && inputcode.Country != 52 {
		return 406, true, "El codigo de pais ingresado no esta incluido en la lista países disponibles de Restoner", phone_and_code
	}

	/*if inputcode.PhoneRegister_Key == 938198374 || inputcode.PhoneRegister_Key == 938488229 {
		return 406, true, "Su equipo ha sido identificado como sospechoso", phone_and_code
	}*/

	//Verificamos que no sea spam
	quantity, _ := worker_repository.Pg_Find_QtyCodesRegistered(inputcode.PhoneRegister_Key, inputcode.Code)
	if quantity > 9 {
		return 406, true, "Este número ha sido bloqueado debido a multiples intentos de ingreso", phone_and_code
	}

	//Generamos un codigo random
	num_alea := rand.Intn(999999-100000) + 100000
	//Asignamos el codigo al modelo Code
	inputcode.Code = num_alea

	//Enviamos el codigo al anfitrion
	/*client := twilio.NewRestClientWithParams(twilio.RestClientParams{
		Username: "ACeb643456bb0e06813948315b95c3aa98",
		Password: "b6febb18bf85369763c4a303937137d9",
	})
	params := &openapi.CreateMessageParams{}
	params.SetTo("+" + strconv.Itoa(inputcode.Country) + strconv.Itoa(inputcode.PhoneRegister_Key))
	params.SetFrom("+18455793864")
	params.SetBody("Su codigo de Restoner es: " + strconv.Itoa(num_alea))
	_, error_sendcode := client.ApiV2010.CreateMessage(params)
	if error_sendcode != nil {
		return 500, true, "Error en el servidor interno al intentar enviar el codigo", phone_and_code
	}*/

	//Notificación WABA
	notificacionWABA(inputcode.Country, inputcode.PhoneRegister_Key, strconv.Itoa(num_alea))

	//Buscamos si el numero ya ha sido registrado en el modelo Code
	phoneregister, err_add := code_repository.Re_Set_Phone(inputcode)
	if err_add != nil {
		return 500, true, "Error en el servidor interno al intentar registrar el código, detalle: " + err_add.Error(), phone_and_code
	}

	err_update := worker_repository.Pg_Update_QtyCodesRegistered(phoneregister, inputcode.Code)
	if err_update != nil {
		return 500, true, "Error en el servidor interno al intentar actualizar la cantidad de codigos requeridos por este comensal, detalle: " + err_update.Error(), phone_and_code
	}

	/*---------------------------GUARDAMOS LOS NUMEROS QUE PIDEN CODIGO DIARIO---------------------------*/
	error_set_quantity := code_repository.Re_Set_Request(inputcode.PhoneRegister_Key, inputcode.Country, quantity_request+1)
	if error_set_quantity != nil {
		return 500, true, "Error en el servidor interno al intentar actualizar la cantidad de codigos soliticidos, detalle: " + error_set_quantity.Error(), phone_and_code
	}
	/*---------------------------------------------------------------------------------------------------*/

	phone_and_code.Phone = phoneregister
	phone_and_code.Country = inputcode.Country

	//Si todo ha ido bien envia un Status 200
	return 201, false, "", phone_and_code
}

func UpdateWithCode_Service(input_phoneregister int, input models.Re_SetGetCode, input_country int) (int, bool, string, PhoneCountryCode) {

	//Instanciamos la variable del help
	var resp PhoneCountryCode

	//Validamos si esta registrado en el modelo Code
	codigo, error_r := code_repository.Re_Get_Phone(input_phoneregister, input_country)
	if error_r != nil {
		return 500, true, "Error en el servidor interno al intentar buscar el número, detalles: " + error_r.Error(), resp
	}
	if codigo.PhoneRegister_Key < 8 {
		return 404, true, "Este numero no se encuentra registrado, numero: " + strconv.Itoa(codigo.PhoneRegister_Key), resp
	}
	if input.Code != codigo.Code {
		return 403, true, "Codigo inválido", resp
	}

	//Validamos si esta registrado en el modelo
	anfitrion_found, _ := worker_repository.Pg_FindByPhone(input_phoneregister, input_country)
	if anfitrion_found.IdBusiness > 8 && !anfitrion_found.IsDeleted {
		return 403, true, "999" + "Este número ya se ha registrado", resp
	}

	resp.Country = input_country
	resp.Phone = input_phoneregister
	resp.Code = codigo.Code

	return 201, false, "", resp
}

func UpdateWithCodeRecovery_Service(input_phoneregister int, input models.Re_SetGetCode, input_country int) (int, bool, string, PhoneCountryCode) {

	//Instanciamos la variable del help
	var resp PhoneCountryCode

	//Validamos si esta registrado en el modelo Code
	codigo, _ := code_repository.Re_Get_Phone(input_phoneregister, input_country)
	if codigo.PhoneRegister_Key < 8 {
		return 404, true, "Este numero no se encuentra registrado", resp
	}
	if input.Code != codigo.Code {
		return 403, true, "Codigo inválido", resp
	}

	//Validamos si esta registrado en el modelo
	anfitrion_found, _ := worker_repository.Pg_FindByPhone(input_phoneregister, input_country)
	if anfitrion_found.IdBusiness < 8 && anfitrion_found.IsDeleted {
		return 403, true, "Este número no se encuentra registrado", resp
	}

	resp.Country = input_country
	resp.Phone = input_phoneregister
	resp.Code = codigo.Code

	return 201, false, "", resp
}

func RegisterAnfitrion_Service(input_anfitrion models.Pg_BusinessWorker) (int, bool, string, string) {

	//Validamos si esta registrado en el modelo Code
	codigo, _ := code_repository.Re_Get_Phone(input_anfitrion.Phone, input_anfitrion.IdCountry)
	if codigo.PhoneRegister_Key < 6 {
		return 404, true, "Este numero no se encuentra registrado", ""
	}
	if input_anfitrion.CodeRedis != codigo.Code {
		return 403, true, "Codigo inválido", ""
	}

	//Validamos si esta registrado en el modelo
	anfitrion_found, _ := worker_repository.Pg_FindByPhone(input_anfitrion.Phone, input_anfitrion.IdCountry)

	if anfitrion_found.Phone > 2 && !anfitrion_found.IsDeleted {
		return 403, true, "Este número ya se ha registrado", ""
	}

	//Creamos un codigo de sesion
	hour, minute, sec := time.Now().Clock()

	//Encriptar password
	encrypted_pass, _ := encrypt(input_anfitrion.Password)
	input_anfitrion.Password = encrypted_pass
	input_anfitrion.UpdatedDate = time.Now()
	input_anfitrion.SessionCode = minute*100 + sec + hour + 1111 + rand.Intn(7483647)

	//Enviamos la variable instanciada al repository
	idworker_business, error_insert_anfitrion := worker_repository.Pg_Add(input_anfitrion)
	if error_insert_anfitrion != nil {
		return 500, true, "Error interno en el servidor al intentar registrar al anfitrion, detalle: " + error_insert_anfitrion.Error(), ""
	}

	go func() {
		worker_repository.Pg_Update_IdBusiness(idworker_business)
	}()

	//Registramos en Redis
	_, err_add_re := worker_repository.Re_Set_ID(idworker_business, input_anfitrion.IdCountry, input_anfitrion.SessionCode, idworker_business)
	if err_add_re != nil {
		return 500, true, "Error en el servidor interno al intentar registrar el código en cache, detalle: " + err_add_re.Error(), ""
	}

	//Notificación WABA
	notificacion_registor_waba(idworker_business, "-5")

	//Enviamos a actualizar la URL en el banner
	go func() {
		//Comienza el proceso de MQTT
		ch, error_conection := models.MqttCN.Channel()
		if error_conection != nil {
			log.Error(error_conection)
		}

		//Enviamos a serializar los datos
		input_anfitrion.IdBusiness = idworker_business
		input_anfitrion.IdWorker = idworker_business
		input_anfitrion.Phone = 0
		input_anfitrion.Password = ""
		input_anfitrion.SubsidiaryOf = idworker_business
		input_anfitrion.IsSubsidiary = false
		bytes, error_serializar := serialize(input_anfitrion)
		if error_serializar != nil {
			log.Error(error_serializar)
		}
		error_publish := ch.Publish("", "anfitrion/businessdata", false, false,
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "text/plain",
				Body:         bytes,
			})
		if error_publish != nil {
			log.Error(error_publish)
		}

	}()

	return 201, false, "", "Registro exitoso"
}

func UpdatePassword_Recover_Service(input_entrydata EntryData_Password) (int, bool, string, string) {

	//Validamos si esta registrado en el modelo Code
	codigo, _ := code_repository.Re_Get_Phone(input_entrydata.Phone, input_entrydata.Country)
	if codigo.PhoneRegister_Key < 8 {
		return 404, true, "Este numero no se encuentra registrado", ""
	}
	if input_entrydata.Code != codigo.Code {
		return 403, true, "Codigo inválido", ""
	}

	//Creamos un nuevo  codigo de sesion
	hour, minute, sec := time.Now().Clock()
	new_session_code := minute*100 + sec + hour + 1111 + rand.Intn(7483647)

	//Encriptar password
	encrypted_pass, _ := encrypt(input_entrydata.NewPassword)

	//Enviamos la variable instanciada al repository
	error_update_password := worker_repository.Pg_Update_Password_Recovery(encrypted_pass, input_entrydata.Phone, input_entrydata.Country, new_session_code)
	if error_update_password != nil {
		return 500, true, "Error interno en el servidor al intentar actualizar la contraseña, detalle: " + error_update_password.Error(), ""
	}

	return 201, false, "", "Contraseña actualizada correctamente"
}

func RegisterColaborador_Service(data_idbusiness int, input_anfitrion models.Pg_BusinessWorker) (int, bool, string, string) {

	//Validamos la cantidad de Colaboradores
	_, quantity_subworkers, _ := worker_repository.Pg_Find_Qty_SubWorkers(data_idbusiness)
	if quantity_subworkers > 2 {
		return 404, true, "Solo se puede registrar como máximo 2 colaboradores", ""
	}

	//Validamos si esta registrado en el modelo Code
	codigo, _ := code_repository.Re_Get_Phone(input_anfitrion.Phone, input_anfitrion.IdCountry)
	if codigo.PhoneRegister_Key < 6 {
		return 404, true, "Este numero no se encuentra registrado", ""
	}
	if input_anfitrion.CodeRedis != codigo.Code {
		return 403, true, "Codigo inválido", ""
	}

	//Validamos si esta registrado en el modelo
	anfitrion_found, _ := worker_repository.Pg_FindByPhone(input_anfitrion.Phone, input_anfitrion.IdCountry)

	if anfitrion_found.Phone > 2 && !anfitrion_found.IsDeleted {
		return 403, true, "Este número ya se ha registrado", ""
	}

	//Creamos un codigo de sesion
	hour, minute, sec := time.Now().Clock()

	//Variable para asignar el rol
	var rol int

	//Encriptar password
	encrypted_pass, _ := encrypt(input_anfitrion.Password)
	input_anfitrion.Password = encrypted_pass
	input_anfitrion.UpdatedDate = time.Now()
	input_anfitrion.SessionCode = minute*100 + sec + hour + 1111 + rand.Intn(7483647)
	input_anfitrion.IdBusiness = data_idbusiness

	if input_anfitrion.IdRol == 0 {
		rol = 2
	} else {
		rol = 3
	}

	//Enviamos la variable instanciada al repository
	idsubworker, error_insert_anfitrion := worker_repository.Pg_Add_Subworker(rol, input_anfitrion)
	if error_insert_anfitrion != nil {
		return 500, true, "Error interno en el servidor al intentar registrar al colaborador, detalle: " + error_insert_anfitrion.Error(), ""
	}

	//Registramos en Redis
	_, err_add_re := worker_repository.Re_Set_ID(idsubworker, input_anfitrion.IdCountry, input_anfitrion.SessionCode, input_anfitrion.IdBusiness)
	if err_add_re != nil {
		return 500, true, "Error en el servidor interno al intentar registrar el código en cache, detalle: " + err_add_re.Error(), ""
	}

	return 201, false, "", "Registro exitoso"
}

//SERIALIZADORA
func serialize(anfitrion models.Pg_BusinessWorker) ([]byte, error) {
	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	err := encoder.Encode(anfitrion)
	if err != nil {
		return b.Bytes(), err
	}
	return b.Bytes(), nil
}

/*=======================================*/
/*===============VERSION 2===============*/
/*=======================================*/

func V2_RegisterColaborador_Service(data_idbusiness int, input_anfitrion models.Pg_BusinessWorker, data_country int) (int, bool, string, string) {

	//Validamos si esta registrado en el modelo
	anfitrion_found, _ := worker_repository.Pg_FindByEmail(input_anfitrion.Email)
	if len(anfitrion_found.Email) > 2 {
		return 403, true, "Este email ya se ha registrado", ""
	}

	//variable del rol
	var rol int

	//Creamos un codigo de sesion
	hour, minute, sec := time.Now().Clock()

	//Encriptar password
	encrypted_pass, _ := encrypt(input_anfitrion.Password)
	input_anfitrion.Password = encrypted_pass
	input_anfitrion.UpdatedDate = time.Now()
	input_anfitrion.SessionCode = minute*100 + sec + hour + 1111 + rand.Intn(7483647)
	input_anfitrion.IdBusiness = data_idbusiness
	input_anfitrion.Phone = 000000000
	input_anfitrion.IdCountry = data_country

	if input_anfitrion.IdRol == 0 || input_anfitrion.IdRol == 2 {
		rol = 2
	} else {
		rol = 3
	}

	//Enviamos la variable instanciada al repository
	idsubworker, error_insert_anfitrion := worker_repository.V2_Pg_Add_Subworker(rol, input_anfitrion)
	if error_insert_anfitrion != nil {
		return 500, true, "Error interno en el servidor al intentar registrar al colaborador, detalle: " + error_insert_anfitrion.Error(), ""
	}

	//Registramos en Redis
	err_add_re := worker_repository.Re_Set_Email(idsubworker, input_anfitrion.SessionCode, input_anfitrion.IdRol)
	if err_add_re != nil {
		return 500, true, "Error en el servidor interno al intentar registrar el código en cache, detalle: " + err_add_re.Error(), ""
	}

	/*--SENT NOTIFICATION--*/
	notification := map[string]interface{}{
		"message":  "Dale una calurosa bienvenida al nuevo integrate del equipo: " + input_anfitrion.Name + " " + input_anfitrion.LastName,
		"iduser":   data_idbusiness,
		"typeuser": 1,
		"priority": 1,
		"title":    "Restoner anfitriones",
	}
	json_data, _ := json.Marshal(notification)
	http.Post("http://c-a-notificacion-tip.restoner-api.fun:5800/v1/notification", "application/json", bytes.NewBuffer(json_data))
	/*---------------------*/

	return 201, false, "", "Registro exitoso"
}
