package main

import "fmt"

const consumers = 10
const producers = 5
const messages = 30

func consume(number int, queue <-chan string, consumed chan<- bool) {
	for message := range queue {
		fmt.Printf("consumer %2d consumes message \"%s\"\n", number, message)
	}
	consumed<- true
}

func produce(number int, queue chan<- string, produced chan<- bool) {
	for i := 0; i < messages / producers; i++ {
		queue <- fmt.Sprintf("Yooo-Hooo! %2d from producer %2d", i+1, number)
	}
    produced<- true
}

func main() {
	queue := make(chan string)
	producerControl := make(chan bool)
	consumerControl := make(chan bool)
	for i := 0; i < consumers; i++ {
		go consume(i + 1, queue, consumerControl)
	}
    for i := 0; i < producers; i++ {
        go produce(i + 1, queue, producerControl)
    }
    for i := 0; i < producers; i++ {
        <-producerControl
    }
    close(queue)
    for i := 0; i < consumers; i++ {
        <-consumerControl
    }
}
