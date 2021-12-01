package worker

type Worker interface {
	Do(rawLog []byte)
}