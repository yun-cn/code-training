package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan string)

	select {
	case msg := <-messages:
		fmt.Println("msg: ", msg)
	default:
		fmt.Println("No message received")
	}

	msg := "h1"
	select {
	case messages <- msg:
		fmt.Println("send message: ", msg)
	default:
		fmt.Println("No message send.")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

}
