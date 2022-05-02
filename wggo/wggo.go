package wggo

import "sync"

var (
	Wg sync.WaitGroup
)

func WgGo(f func()) {
	Wg.Add(1)
	go func() {
		defer Wg.Done()
		f()
	}()
}
