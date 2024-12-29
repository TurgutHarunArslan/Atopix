package events

import "sync"

type EventHandler func(data EventInterface)

type EventBus struct {
	handlers map[string][]EventHandler
	mu       sync.RWMutex
}

func New() *EventBus {
	return &EventBus{
		handlers: make(map[string][]EventHandler),
	}
}

func (eb *EventBus) Subscribe(event string, handler EventHandler) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	eb.handlers[event] = append(eb.handlers[event], handler)
}

func (eb *EventBus) Publish(Event EventInterface) {
	eb.mu.RLock()
	defer eb.mu.RUnlock()

	if handlers, ok := eb.handlers[Event.Type()]; ok {
		for _, handler := range handlers {
			go handler(Event)
		}
	}
}
