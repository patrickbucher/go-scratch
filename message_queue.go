package main

import "fmt"

const CONSUMERS = 5
const MESSAGES = 40

func consume(number int, channel <-chan string, drained chan<- bool) {
    for message := range channel {
        fmt.Printf("consumer %d consumes message \"%s\"\n", number, message)
    }
    drained<- true
}

func main() {
    queue := make(chan string)
    control := make(chan bool)
    for i := 0; i < CONSUMERS; i++ {
        go consume(i + 1, queue, control)
    }
    for i := 0; i < MESSAGES / 2; i++ {
        queue<- fmt.Sprintf("hello %2d", i + 1)
        queue<- fmt.Sprintf("goodbye %2d", i + 1)
    }
    close(queue)
    <-control
}
