package notify

import (
	"demoapp/config"
	"io/ioutil"
	"log"
	"sync"
	"time"

	"github.com/valyala/fasttemplate"
	mail "github.com/xhit/go-simple-mail/v2"
)

// EmailNotidy is a email notify
type EmailNotidy struct {
	sync.Mutex
	config        config.Config
	smtpClient    *mail.SMTPClient
	templateCache map[string]string
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
	server.SendTimeout = 0
	smtpClient, err := server.Connect()
	if err != nil {
		log.Fatalf("init mail instance error:%s", err.Error())
	}
	bytes, err := ioutil.ReadFile(config.Template.EmailTemplate)
	if err != nil {
		log.Fatalf("init mail instance error:%s", err.Error())
	}
	return &EmailNotidy{
		config:     config,
		smtpClient: smtpClient,
		templateCache: map[string]string{
			config.Template.EmailTemplate: string(bytes),
		},
	}
}

// Send Send
func (e *EmailNotidy) Send(to string, subject string, datafiles map[string]interface{}) error {
	// add lock for goroutine
	e.Lock()
	defer e.Unlock()
	t := fasttemplate.New(e.templateCache[e.config.Template.EmailTemplate], "{{", "}}")
	htmlBody := t.ExecuteString(datafiles)
	email := mail.NewMSG()
	from := e.config.Email.FromEmail
	email.SetFrom(from).
		AddTo(to).
		AddCc([]string{"dalongdemo@qq.com"}...).
		SetSubject(subject)

	email.SetBody(mail.TextHTML, htmlBody)
	err := email.Send(e.smtpClient)
	if err != nil {
		return err
	}
	return nil
}
