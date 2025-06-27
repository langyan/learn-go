package eventbus

import "sync"

type Event struct {
	Payload any
}
type (
	EventChan chan Event
)
type EventBus struct {
	mu          sync.RWMutex
	subscribers map[string][]EventChan
}

func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[string][]EventChan),
	}
}

func(eb * EventBus) Publish(topic string,event Event){
	eb.mu.RLock()
	defer eb.mu.RUnlock()

	// Copy a new subscriber list to avoid modifying the list while publishing
	subscribers := append([]EventChan{}, eb.subscribers[topic]...)
	go func() {
	  for _, subscriber := range subscribers {
		subscriber <- event
	  }
	}()

}

func (eb *EventBus) Subscribe(topic string) EventChan {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	ch := make(EventChan)
	eb.subscribers[topic] = append(eb.subscribers[topic], ch)
	return ch
  }

  func (eb *EventBus) Unsubscribe(topic string, ch EventChan) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	if subscribers, ok := eb.subscribers[topic]; ok {
	  for i, subscriber := range subscribers {
		if ch == subscriber {
		  eb.subscribers[topic] = append(subscribers[:i], subscribers[i+1:]...)
		  close(ch)
		  // Drain the channel
		  for range ch {
		  }
		  return
		}
	  }
	}
  }


