package worker

import "time"

const (
	defaultBatchMaxSize = 100
	defaultInterval     = 3
)

type Consume func(*Message)

type Writer interface {
	Write([]*Message)
}

type Runner struct {
	W            Writer
	BatchMaxSize int
	Interval     time.Duration

	done  chan struct{}
	cache chan *Message
}

func NewRunner(w Writer, batchMaxSize, interval int) *Runner {
	if batchMaxSize <= 0 {
		batchMaxSize = defaultBatchMaxSize
	}
	if interval <= 0 {
		interval = defaultInterval
	}
	return &Runner{
		W:            w,
		BatchMaxSize: batchMaxSize,
		Interval:     time.Second * time.Duration(interval),

		done:  make(chan struct{}),
		cache: make(chan *Message),
	}
}

func (r *Runner) NewConsume() Consume {
	return func(m *Message) {
		r.cache <- m
	}
}

func (r *Runner) RefreshLoop() {
	buf := make([]*Message, 0, r.BatchMaxSize)
	tk := time.NewTicker(r.Interval)
	defer tk.Stop()
	for {
		select {
		case <-r.done:
			r.W.Write(buf)
			return
		case <-tk.C:
			r.W.Write(buf)
		case m := <-r.cache:
			buf = append(buf, m)
			if len(buf) >= r.BatchMaxSize {
				r.W.Write(buf)
			}
		}
	}
}
