package main

// 带方向的channel不能被关闭
func Stop(stop <-chan bool)  {
	close(stop)
}