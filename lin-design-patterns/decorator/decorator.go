package decorator

import "log"

// Why it matters: You can add functionality without modifying original code.
type Service interface {
	DoWork() error
}

type LoggingService struct {
	next Service
}

func (l LoggingService) DoWork() error {
	log.Println("before work")
	err := l.next.DoWork()
	log.Println("after work")
	return err
}
