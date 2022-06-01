package smtp

import "net/smtp"

type Email struct {
	option *Option
}
type Option struct {
	Username  string
	Cc        []string
	Receivers []string
	Bcc       []string
}

func (e *Email) Send(content string, receivers []string, cc []string) error {
	//auth := smtp.PlainAuth("", "ccb1900@qq.com", "123456", "127.0.0.1:1026")
	err := smtp.SendMail(
		"127.0.0.1:1026",
		nil,
		"ccb1900@qq.com",
		[]string{"ccb1900@qq.com"},
		[]byte("To: recipient@example.net\r\n"+
			"Subject: discount Gophers!\r\n"+
			"\r\n"+
			"<h1>This is the email body.</h1>\r\n"),
	)
	if err != nil {
		return err
	}
	return nil
}

func New(option *Option) *Email {
	return &Email{
		option: option,
	}
}
