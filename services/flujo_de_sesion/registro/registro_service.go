package registro

import (
	"bytes"
	"encoding/json"
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
)

//FUNCIONES PRIVADAS
func encrypt(input string) (string, error) {
	cost := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(input), cost)
	return string(bytes), err
}

//FUNCIONES PUBLICAS

func SignUpNumber_Service(inputcode models.Re_SetGetCode) (int, bool, string, SignInFirstStep) {

	var phone_and_code SignInFirstStep

	if inputcode.Country != 51 && inputcode.Country != 52 {
		return 406, true, "El codigo de pais ingresado no esta incluido en la lista países disponibles de Restoner", phone_and_code
	}

	//Verificamos que no sea spam
	quantity, _ := worker_repository.Pg_Find_QtyCodesRegistered(inputcode.PhoneRegister_Key, inputcode.Code)
	if quantity > 9 {
		return 406, true, "Este número ha sido bloqueado debido a multiples intentos de ingreso", phone_and_code
	}

	//Generamos un codigo random
	//num_alea := rand.Intn(999999-100000) + 100000
	num_alea := 123456
	//Asignamos el codigo al modelo Code
	inputcode.Code = num_alea

	//Enviamos el codigo al anfitrion
	/*client := twilio.NewRestClientWithParams(twilio.RestClientParams{
		Username: "ACaef214e389677f1f21534dd1dd77c609",
		Password: "95fe92994dbd9d82f2aa47fb9dc94daa",
	})
	params := &openapi.CreateMessageParams{}
	params.SetTo("+" + strconv.Itoa(inputcode.PhoneRegister_Key))
	params.SetFrom("+17244143326")
	params.SetBody("Codigo de Restoner es: " + strconv.Itoa(num_alea))
	_, error_sendcode := client.ApiV2010.CreateMessage(params)
	if error_sendcode != nil {
		return 500, true, "Error en el servidor interno al intentar enviar el codigo", 0
	}*/

	//Buscamos si el numero ya ha sido registrado en el modelo Code
	phoneregister, err_add := code_repository.Re_Set_Phone(inputcode)
	if err_add != nil {
		return 500, true, "Error en el servidor interno al intentar registrar el código, detalle: " + err_add.Error(), phone_and_code
	}

	err_update := worker_repository.Pg_Update_QtyCodesRegistered(phoneregister, inputcode.Code)
	if err_update != nil {
		return 500, true, "Error en el servidor interno al intentar actualizar la cantidad de codigos requeridos por este comensal, detalle: " + err_update.Error(), phone_and_code
	}

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
		return 500, true, "Error en els ervidor interno al intentar buscar el número, detalles: " + error_r.Error(), resp
	}
	if codigo.PhoneRegister_Key < 8 {
		return 404, true, "Este numero no se encuentra registrado, numero: " + strconv.Itoa(codigo.PhoneRegister_Key), resp
	}
	if input.Code != codigo.Code {
		return 403, true, "Codigo inválido", resp
	}

	//Validamos si esta registrado en el modelo
	anfitrion_found, _ := worker_repository.Pg_FindByPhone(input_phoneregister, input_country)
	if anfitrion_found.IdBusiness < 8 {
		return 403, true, "Este número no se encuentra registrado", resp
	}
	if anfitrion_found.IdBusiness > 8 {
		return 403, true, "Este número ya se ha registrado", resp
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
	if anfitrion_found.IdBusiness < 8 {
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

	if anfitrion_found.Phone > 2 {
		return 403, true, "Este número ya se ha registrado", ""
	}

	//Encriptar password
	encrypted_pass, _ := encrypt(input_anfitrion.Password)
	input_anfitrion.Password = encrypted_pass
	input_anfitrion.UpdatedDate = time.Now()

	//Enviamos la variable instanciada al repository
	idworker_business, error_insert_anfitrion := worker_repository.Pg_Add(input_anfitrion)
	if error_insert_anfitrion != nil {
		return 500, true, "Error interno en el servidor al intentar registrar al anfitrion, detalle: " + error_insert_anfitrion.Error(), ""
	}

	go func() {
		worker_repository.Pg_Update_IdBusiness(idworker_business)
	}()

	//Registramos en Redis
	_, err_add_re := worker_repository.Re_Set_Id(idworker_business, input_anfitrion.IdCountry)
	if err_add_re != nil {
		return 500, true, "Error en el servidor interno al intentar registrar el código en cache, detalle: " + err_add_re.Error(), ""
	}

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

	time.Sleep(2 * time.Second)

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

	//Encriptar password
	encrypted_pass, _ := encrypt(input_entrydata.NewPassword)

	//Enviamos la variable instanciada al repository
	error_update_password := worker_repository.Pg_Update_Password_Recovery(encrypted_pass, input_entrydata.Phone, input_entrydata.Country)
	if error_update_password != nil {
		return 500, true, "Error interno en el servidor al intentar actualizar la contraseña, detalle: " + error_update_password.Error(), ""
	}

	return 201, false, "", "Contraseña actualizada correctamente"
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
