package notify

import (
	"demoapp/config"
	"io/ioutil"
	"log"

	"github.com/valyala/fasttemplate"
	"gopkg.in/gomail.v2"
)

// EmailNotidy2 is a email notify
type EmailNotidy2 struct {
	config        config.Config
	dialer        *gomail.Dialer
	templateCache map[string]string
}

// NewEailNotidy2 NewEailNotidy2 instance
func NewEailNotidy2() *EmailNotidy2 {
	config := config.New()
	d := gomail.NewDialer(config.Email.ServerHost, config.Email.ServerPort, config.Email.FromEmail, config.Email.FromPasswd)
	bytes, err := ioutil.ReadFile(config.Template.EmailTemplate)
	if err != nil {
		log.Fatalf("init mail instance error:%s", err.Error())
	}
	return &EmailNotidy2{
		config: config,
		dialer: d,
		templateCache: map[string]string{
			config.Template.EmailTemplate: string(bytes),
		},
	}
}

// Send Send
func (e *EmailNotidy2) Send(to string, subject string, datafiles map[string]interface{}) error {
	t := fasttemplate.New(e.templateCache[e.config.Template.EmailTemplate], "{{", "}}")
	htmlBody := t.ExecuteString(datafiles)
	m := gomail.NewMessage()
	m.SetHeader("From", e.config.Email.FromEmail)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", htmlBody)
	sender, err := e.dialer.Dial()
	err = sender.Send(e.config.Email.FromEmail, []string{to}, m)
	if err != nil {
		return err
	}
	return nil
}
