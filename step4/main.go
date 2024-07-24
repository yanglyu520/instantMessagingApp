package main

// test with nc 127.0.0.1 8888, which simulate a client sending http request
func main() {
	server := NewServer("127.0.0.1", 8888)
	server.Start()
}
