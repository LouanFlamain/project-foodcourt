package request

type EmailRequest struct {
	To       string
	Subject  string
	Template string
	Data     interface{}
}
