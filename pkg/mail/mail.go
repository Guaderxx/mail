package mail

import (
	"crypto/tls"
	"log"
	"time"

	"github.com/Guaderxx/mail/pkg/model"
	"github.com/toorop/go-dkim"
	mail "github.com/xhit/go-simple-mail/v2"
)

const privateKey = ""

func initClient(u model.User) *mail.SMTPClient {
	server := mail.NewSMTPClient()
	server.Host = u.Host
	server.Port = u.Port
	server.Username = u.Username
	server.Password = u.Password
	// default EncryptionSTARTTLS
	server.Encryption = u.EncryptionType

	server.KeepAlive = false
	server.ConnectTimeout = time.Second * time.Duration(u.ConnectTimeout)
	server.SendTimeout = time.Second * time.Duration(u.SendTimeout)
	server.TLSConfig = &tls.Config{
		InsecureSkipVerify: true,
	}

	client, err := server.Connect()
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func Send() {
	client := initClient(model.U)
	msg := mail.NewMSG()
	msg.AddTo(model.M.To).
		SetSubject(model.M.Subject)

	msg.SetBody(mail.TextHTML, model.M.Content)

	if privateKey != "" {
		options := dkim.NewSigOptions()
		options.PrivateKey = []byte(privateKey)
		options.Domain = ""
		options.Selector = ""
		options.SignatureExpireIn = 3600
		options.Headers = []string{"from", "date", "mime-version", "received"}
		options.AddSignatureTimestamp = true
		options.Canonicalization = "relaxed/relaxed"

		msg.SetDkim(options)
	}

	err := msg.Send(client)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Email sent")
}
