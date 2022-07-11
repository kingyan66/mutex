package main

import (
	"fmt"
	"sync"
)

type Mutex struct {
	ch chan struct{}
}

type AnotherMutex struct {
	cn chan struct{}
}

// 第一种实现方式
// 谁先取谁取得了这个元素，就相当于获取了这把锁。
func NewMutex() *Mutex {
	m := &Mutex{ch: make(chan struct{}, 1)}
	m.ch <- struct{}{}
	return m
}

// 第二种实现方式
// 谁能成功地把元素发送到这个 Channel，谁就获取了这把锁。
func NewAnotherMutex() *AnotherMutex {
	m := &AnotherMutex{cn: make(chan struct{}, 1)}
	return m
}

func main() {
	var wg sync.WaitGroup
	var count1 int
	var count2 int
	m := NewMutex()
	am := NewAnotherMutex()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			am.Lock()
			count1++
			am.Unlock()
			wg.Done()
		}()
	}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			m.Lock()
			count2++
			m.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Print(count1, count2)

func (m *Mutex) Lock() {
	<-m.ch
}

func (m *Mutex) Unlock() {
	// select {
	// case m.ch <- struct{}{}:
	// default:
	// 	fmt.Println("fatal")
	// }
	m.ch <- struct{}{}
}

func (am *AnotherMutex) Lock() {
	am.cn <- struct{}{}
}

func (am *AnotherMutex) Unlock() {
	<-am.cn
}
