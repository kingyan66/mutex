package main

type Mutex struct {
	ch chan struct{}
}

func NewMutex() *Mutex {
	m := &Mutex{make(chan struct{}, 1)}
	m.ch <- struct{}{}
	return m
}

func (m *Mutex) Lock() {
	<-m.ch
}

func (m *Mutex) Unlock() {
	select {
	case m.ch <- struct{}{}:
	default:
		panic("unlock of unlock mutex")
	}
}
