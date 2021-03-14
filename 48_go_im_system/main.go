package main

func main() {
	var Server = NewServer("127.0.0.1", 8888)
	Server.Start()
}
