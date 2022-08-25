package utils

import "sync"

type WithEventLoop interface {
	EventLoop()
}

func WaitOne(one WithEventLoop, wg *sync.WaitGroup) {
	defer wg.Done()
	one.EventLoop()
}

func WaitAll(el ...WithEventLoop) {
	var wg sync.WaitGroup
	wg.Add(len(el))

	for _, e := range el {
		go WaitOne(e, &wg)
	}

	wg.Wait()
}
