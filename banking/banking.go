package banking

import (
    "sync"
)

type Account struct {
    owner string
    number int
    sync.Mutex
    balance int
}

func (acc *Account) listen(transfers <-chan int, control chan<- bool) {
    for amount := range transfers {
        acc.Lock()
        acc.balance += amount
        acc.Unlock()
    }
    control <- true
}

type transfer struct {
    source *Account
    target *Account
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

func ExecuteTransfers(source, target []*Account, amount int) {
    if len(source) != len(target) {
        panic("different number of source and target accounts")
    }
    accounts := len(source)

    fwdTrans := make([]*transfer, accounts)
    bckTrans := make([]*transfer, accounts)
    for i := 0; i < accounts; i++ {
        fwdTrans[i] = &transfer{source: source[i], target: target[i],
            amount: amount}
        bckTrans[i] = &transfer{source: target[i], target: source[i],
            amount: amount}
    }

    status := make(chan string)
    for i := 0; i < accounts; i++ {
        go fwdTrans[i].execute(status)
        go bckTrans[i].execute(status)
    }

    for i := 0; i < accounts * 2; i++ {
        <-status
    }
    close(status)
}
