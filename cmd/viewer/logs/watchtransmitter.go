package logs

import "sync"

type WatchTransmitter struct {
	watchers map[string][]Receiver
	mux      sync.Mutex
}

func NewWatchTransmitter() *WatchTransmitter {
	return &WatchTransmitter{
		watchers: map[string][]Receiver{},
		mux:      sync.Mutex{},
	}
}

func (t *WatchTransmitter) Receive(msg Message) {
	wats := t.watchers[msg.MessageType]
	for _, w := range wats {
		w.Receive(msg)
	}
}

func (t *WatchTransmitter) Distribute(types []string, rec Receiver) {
	t.mux.Lock()
	defer t.mux.Unlock()
	for _, mt := range types {
		if t.watchers[mt] == nil {
			t.watchers[mt] = []Receiver{}
		}
		t.watchers[mt] = append(t.watchers[mt], rec)
	}
}
