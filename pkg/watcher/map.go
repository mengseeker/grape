package watcher

// import "sync"

// type MapWatcher struct {
// 	l     sync.RWMutex
// 	m     map[string]map[int]bool
// 	p     map[int]Processer
// 	idSeq int
// }

// func NewMapWatcher() *MapWatcher {
// 	return &MapWatcher{
// 		l:     sync.RWMutex{},
// 		m:     map[string]map[int]bool{},
// 		p:     map[int]Processer{},
// 		idSeq: 1000,
// 	}
// }

// func (w *MapWatcher) Registry(p Processer) int {
// 	w.l.Lock()
// 	defer w.l.Unlock()
// 	id := w.idSeq
// 	w.idSeq++
// 	key := p.WatchKey()
// 	if w.m[key] == nil {
// 		w.m[key] = map[int]bool{id: true}
// 	} else {
// 		w.m[key][id] = true
// 	}
// 	w.p[id] = p
// 	return id
// }

// func (w *MapWatcher) UnRegistry(id int) {
// 	w.l.Lock()
// 	defer w.l.Unlock()
// 	p := w.p[id]
// 	delete(w.p, id)
// 	delete(w.m[p.WatchKey()], id)
// 	if len(w.m[p.WatchKey()]) == 0 {
// 		delete(w.m, p.WatchKey())
// 	}
// }

// func (w *MapWatcher) Notify(n Notifier) {
// 	w.l.RLock()
// 	defer w.l.RUnlock()
// 	for id := range w.m[n.NotifyKey()] {
// 		w.p[id].Process(n)
// 	}
// }
