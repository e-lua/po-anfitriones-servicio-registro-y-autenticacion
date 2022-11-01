package subsidiary

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"strconv"
	"time"

	"github.com/Aphofisis/po-anfitrion-servicio-registro-y-autenticacion/models"
	subsidiary_repository "github.com/Aphofisis/po-anfitrion-servicio-registro-y-autenticacion/repositories/subsidiary"
	worker_repository "github.com/Aphofisis/po-anfitrion-servicio-registro-y-autenticacion/repositories/worker"
	"github.com/labstack/gommon/log"
	"github.com/streadway/amqp"
	"golang.org/x/crypto/bcrypt"

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
	params.SetBody("*Restoner Bot:* Acaba de iniciar sesión en una cuenta de *" + " REGISTRO DE SUCURSAL " + strconv.Itoa(idbusiness) + "* el " + fecha.Format("2006-01-02 3:4:5 pm") + ", si no fue usted, haga clic aquí https://youtu.be/JDcPVzSP8-M")

	_, err := client.ApiV2010.CreateMessage(params)
	if err != nil {
		return 200, false, "", "Solicitud enviada correctamente"
	}

	return 200, false, "", "Solicitud enviada correctamente"
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

//FUNCIONES PUBLICAS
func AddSubsidiary_Service(idbusiness int, subsidiary models.Pg_BusinessWorker) (int, bool, string, string) {

	//Validamos si esta registrado en el modelo Code
	qty, _ := subsidiary_repository.Pg_Find_Qty_Subsidiary(idbusiness)
	if qty > 5 {
		return 404, true, "Limite de sucursales alcanzadas", ""
	}

	//Creamos un codigo de sesion
	hour, minute, sec := time.Now().Clock()

	//Encriptar password
	encrypted_pass, _ := encrypt(subsidiary.Password)
	subsidiary.Password = encrypted_pass
	subsidiary.UpdatedDate = time.Now()
	subsidiary.SessionCode = minute*100 + sec + hour + 1111 + rand.Intn(7483647)
	subsidiary.IsSubsidiary = true
	subsidiary.SubsidiaryOf = idbusiness

	//Enviamos la variable instanciada al repository
	idworker_business, error_insert_anfitrion := subsidiary_repository.Pg_Add(subsidiary)
	if error_insert_anfitrion != nil {
		return 500, true, "Error interno en el servidor al intentar registrar la subsidiaria, detalle: " + error_insert_anfitrion.Error(), ""
	}

	go func() {
		worker_repository.Pg_Update_IdBusiness(idworker_business)
	}()

	//Registramos en Redis
	_, err_add_re := worker_repository.Re_Set_ID(idworker_business, subsidiary.IdCountry, subsidiary.SessionCode, idworker_business)
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
		subsidiary.IdBusiness = idworker_business
		subsidiary.IdWorker = idworker_business
		subsidiary.Phone = 0
		subsidiary.Password = ""
		subsidiary.SubsidiaryOf = 0
		subsidiary.IsSubsidiary = false
		bytes, error_serializar := serialize(subsidiary)
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
