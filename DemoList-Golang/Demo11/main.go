package main

import "fmt"

func main() {
	message := make(chan string, 2)

	message <- "buffered"
	message <- "rakuten.com"

	fmt.Println(<-message)
	fmt.Println(<-message)
}
