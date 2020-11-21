package notify

import (
	"demoapp/config"
	"io/ioutil"
	"log"
	"time"

	"github.com/valyala/fasttemplate"
	mail "github.com/xhit/go-simple-mail/v2"
)

// EmailNotidy is a email notify
type EmailNotidy struct {
	config     config.Config
	smtpClient *mail.SMTPClient
}

// NewEailNotidy NewEailNotidy instance
func NewEailNotidy() *EmailNotidy {
	config := config.New()
	server := mail.NewSMTPClient()
	// SMTP Server
	server.Host = config.Email.ServerHost
	server.Port = config.Email.ServerPort
	server.Username = config.Email.FromEmail
	server.Password = config.Email.FromPasswd
	server.Encryption = mail.EncryptionNone
	// Since v2.3.0 you can specified authentication type:
	// - PLAIN (default)
	// - LOGIN
	// - CRAM-MD5
	server.Authentication = mail.AuthPlain
	// Variable to keep alive connection
	server.KeepAlive = true
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second
	smtpClient, err := server.Connect()
	if err != nil {
		log.Fatalf("init mail instance error:%s", err.Error())
	}
	return &EmailNotidy{
		config:     config,
		smtpClient: smtpClient,
	}
}

// Send Send
func (e *EmailNotidy) Send(to string, subject string, datafiles map[string]interface{}) error {
	bytes, err := ioutil.ReadFile(e.config.Template.EmailTemplate)
	if err != nil {
		log.Fatalf("read file error:%s", err.Error())
	}
	t := fasttemplate.New(string(bytes), "{{", "}}")
	htmlBody := t.ExecuteString(datafiles)
	email := mail.NewMSG()
	email.SetFrom(e.config.Email.FromEmail).
		AddTo(to).
		AddCc([]string{"dalongdemo@qq.com"}...).
		SetSubject(subject)

	email.SetBody(mail.TextHTML, htmlBody)
	err = email.Send(e.smtpClient)
	if err != nil {
		return err
	}
	return nil
}
