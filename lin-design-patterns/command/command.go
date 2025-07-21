package command

// Encapsulates actions, makes retrying and queuing easy.
type Command interface {
	Execute() error
}

type SendEmailCommand struct {
	To string
}

func (c SendEmailCommand) Execute() error {
	// send email
	return nil
}
