# mutex相关demo
## 使用chan实现mutex的两种方式
无非就是设置一个buffer为1的channel，将这个1作为资源
谁先获得了这个资源，谁就获得了锁（同一时间里只能有一个goroutine能获得这个锁）
或者是谁先把资源发送进去占住为了，谁就获得了锁