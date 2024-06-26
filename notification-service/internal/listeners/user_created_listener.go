package listeners

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"

	"github.com/droquedev/e-commerce/notification-service/pkg/email"
	"github.com/droquedev/e-commerce/pkg/nats"
	"github.com/nats-io/stan.go"
)

type UserCreatedListener struct {
	baseListener *nats.BaseListener
	emailSender  *email.EmailSender
}

func NewUserCreatedListener(client stan.Conn, emailSender *email.EmailSender, queGroupName string) nats.Listener {
	baseListener := nats.NewBaseListener(client, "user:created", queGroupName)
	return &UserCreatedListener{
		baseListener: baseListener,
		emailSender:  emailSender,
	}
}

func (l *UserCreatedListener) OnMessage(message *stan.Msg) {
	log.Printf("Received user.created event: %s", string(message.Data))

	var err error

	var data map[string]interface{}

	err = json.Unmarshal([]byte(message.Data), &data)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	const welcomeMessage = `
		<html>
			<body>
				<h1>Bienvenido {{.USERNAME}}</h1>
				<p>Gracias por registrarte en nuestra plataforma</p>
			</body>
		</html>
	`

	template, err := template.New("test").Parse(welcomeMessage)

	if err != nil {
		panic(err)
	}

	variables := struct {
		USERNAME string
	}{
		USERNAME: data["username"].(string),
	}

	var buf bytes.Buffer
	err = template.Execute(&buf, variables)
	if err != nil {
		log.Fatal(err)
	}

	_, err = l.emailSender.SendEmail(data["email"].(string), "Welcome to our platform", buf.String())

	if err != nil {
		log.Printf("Error sending email: %s", err.Error())
		return
	}

	message.Ack()
}

func (l *UserCreatedListener) Listen() {
	l.baseListener.Listen(l.OnMessage)
}
