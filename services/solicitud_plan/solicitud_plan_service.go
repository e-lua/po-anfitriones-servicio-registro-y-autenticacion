package solicitud_plan

import (
	"log"
	"strconv"
	"time"

	//TWILIO
	twilio "github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func Anfitriones_SendRequest_Service(idbusiness int, timezone string) (int, bool, string, string) {

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
	params.SetBody("*Restoner Bot:* Acaba de iniciar sesión en una cuenta de *" + strconv.Itoa(idbusiness) + "* el " + fecha.Format("2006-01-02 3:4:5 pm") + ", si no fue usted, haga clic aquí https://youtu.be/JDcPVzSP8-M")

	_, err := client.ApiV2010.CreateMessage(params)
	if err != nil {
		log.Println("Error Twilio---->", err.Error())
		return 200, false, "", "Solicitud enviada correctamente"
	}

	return 200, false, "", "Solicitud enviada correctamente"
}
