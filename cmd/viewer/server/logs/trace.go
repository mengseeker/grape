package logs

type Trace struct {
	Log string
}

func (e *Trace) Marshaler() []byte {
	return []byte(e.Log)
}
