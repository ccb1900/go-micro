package email

type Email interface {
	Send(content string, receivers []string, cc []string)
}
