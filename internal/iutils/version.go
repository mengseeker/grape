package iutils

import "time"

func NewVersion() string {
	return time.Now().String()
}
