package main

import (
    "fmt"
    "math"
    "sync"
)

type account struct {
    owner string
    number int
    sync.Mutex
    balance int
}

func (acc *account) listen(transfers <-chan int, control chan<- bool) {
    for amount := range transfers {
        acc.Lock()
        acc.balance += amount
        acc.Unlock()
    }
    control <- true
}

type transfer struct {
    source *account
    target *account
    amount int
}

func (trans transfer) execute(status chan<- string) {
    const PAYMENT = 1
    sourceChannel := make(chan int)
    targetChannel := make(chan int)
    controlChannel := make (chan bool)
    go trans.source.listen(sourceChannel, controlChannel)
    go trans.target.listen(targetChannel, controlChannel)
    for paid := 0; paid < trans.amount; paid += PAYMENT {
        sourceChannel <- -PAYMENT
        targetChannel <- +PAYMENT
    }
    close(sourceChannel)
    close(targetChannel)
    if <- controlChannel && <- controlChannel {
        status <- "transfer done"
    }
}

func main() {
    const ACCOUNTS = 10
    const TRANSFERS = ACCOUNTS * 2
    const AMOUNT = 100
    const BALANCE = 1000

    fooBalance := 0
    barBalance := 0
    foo := [ACCOUNTS]*account{}
    bar := [ACCOUNTS]*account{}
    for i := 0; i < ACCOUNTS; i++ {
        foo[i] = &account{owner: "foo", number: i, balance: BALANCE}
        bar[i] = &account{owner: "bar", number: i, balance: BALANCE}
        fooBalance += foo[i].balance
        barBalance += bar[i].balance
    }

    fooToBar := [ACCOUNTS]*transfer{}
    barToFoo := [ACCOUNTS]*transfer{}
    for i := 0; i < ACCOUNTS; i++ {
        fooToBar[i] = &transfer{source: foo[i], target: bar[i], amount: AMOUNT}
        barToFoo[i] = &transfer{source: bar[i], target: foo[i], amount: AMOUNT}
    }

    status := make(chan string)
    for i := 0; i < ACCOUNTS; i++ {
        go fooToBar[i].execute(status)
        go barToFoo[i].execute(status)
    }

    for i := 0; i < TRANSFERS; i++ {
        fmt.Printf("%d. %s\n", i + 1, <-status)
    }
    close(status)

    for i := 0; i < ACCOUNTS; i++ {
        fooBalance -= foo[i].balance
        barBalance -= bar[i].balance
    }
    if (fooBalance != 0 || barBalance != 0) {
        difference := math.Abs(float64(fooBalance)) + math.Abs(float64(barBalance))
        fmt.Println("Error: difference detected: ", difference)
    } else {
        fmt.Println("Success: no difference detected")
    }
}
