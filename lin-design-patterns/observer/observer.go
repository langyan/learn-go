package observer

type EventBus struct {
	subscribers []chan string
}

func (b *EventBus) Subscribe() chan string {
	ch := make(chan string)
	b.subscribers = append(b.subscribers, ch)
	return ch
}
func (b *EventBus) Publish(msg string) {
	for _, ch := range b.subscribers {
		ch <- msg
	}
}
