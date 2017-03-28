package main

import "fmt"

const consumers = 5
const messages = 40

func consume(number int, channel <-chan string, consumed chan<- bool) {
    for message := range channel {
        fmt.Printf("consumer %d consumes message \"%s\"\n", number, message)
    }
    consumed<- true
}

func main() {
    queue := make(chan string)
    control := make(chan bool)
    for i := 0; i < consumers; i++ {
        go consume(i + 1, queue, control)
    }
    for i := 0; i < messages; i++ {
        queue<- fmt.Sprintf("hello %2d", i + 1)
    }
    close(queue)
    <-control
}
