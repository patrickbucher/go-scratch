package main

import "fmt"

const consumers = 10
const messages = 50

func consume(number int, queue <-chan string, consumed chan<- bool) {
for message := range queue {
fmt.Printf("consumer %2d consumes message \"%s\"\n", number, message)
}
consumed <- true
}

func produce(queue chan<- string) {
	for i := 0; i < messages; i++ {
		queue <- fmt.Sprintf("Yooo-Hooo! %2d", i+1)
	}
	close(queue)
}

func main() {
	queue := make(chan string)
	control := make(chan bool)
	for i := 0; i < consumers; i++ {
		go consume(i+1, queue, control)
	}
	go produce(queue)
	<-control
}
