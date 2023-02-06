package workdemo

import "testing"

func TestNewWorkerManager(t *testing.T) {
	wm := NewWorkerManager(10)

	wm.StartWorkerPool()
}
