package worker

import (
	"context"
	"grape/logtrans/logs"
	"grape/pkg/logger"
	"time"
)

const (
	defaultBatchMaxSize = 100
	defaultInterval     = 3
)

type Message = logs.Message

type Writer interface {
	Write([]Message)
}

type runner struct {
	W            Writer
	BatchMaxSize int
	Interval     time.Duration

	cache chan Message
	l     logger.Logger
}

func NewRunner(w Writer, batchMaxSize, interval int, l logger.Logger) *runner {
	if batchMaxSize <= 0 {
		batchMaxSize = defaultBatchMaxSize
	}
	if interval <= 0 {
		interval = defaultInterval
	}
	return &runner{
		W:            w,
		BatchMaxSize: batchMaxSize,
		Interval:     time.Second * time.Duration(interval),

		cache: make(chan Message),
		l:     l,
	}
}

func (r *runner) Receive(msg Message) {
	r.cache <- msg
}

func (r *runner) RefreshLoop(ctx context.Context) {
	buf := make([]Message, 0, r.BatchMaxSize)
	tk := time.NewTicker(r.Interval)
	defer tk.Stop()
	for {
		select {
		case <-tk.C:
			r.W.Write(buf)
			buf = buf[:0]
		case m := <-r.cache:
			buf = append(buf, m)
			if len(buf) >= r.BatchMaxSize {
				r.W.Write(buf)
				buf = buf[:0]
			}
		case <-ctx.Done():
			r.l.Warnf("wait for write logs %d", len(buf))
			close(r.cache)
			r.W.Write(buf)
			return
		}
	}
}
