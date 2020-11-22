# go-simple-mail+fasttemplate+mailhog

just for test golang email client


## some note 

can't send email message with goroutines to share the same mail.SMTPClient
with `KeepAlive=true`