package banking

import (
    "testing"
    "math"
)

const ACCOUNTS = 100
const AMOUNT = 1000
const BALANCE = 10000
const OWNER_A = "foo"
const OWNER_B = "bar"

func TestExecuteTransfers(t *testing.T) {
    foo, bar := createAccounts(ACCOUNTS, BALANCE)

    fooBalance := 0
    barBalance := 0
    for i := 0; i < ACCOUNTS; i++ {
        fooBalance += foo[i].balance
        barBalance += bar[i].balance
    }

    ExecuteTransfers(foo, bar, AMOUNT)

    for i := 0; i < ACCOUNTS; i++ {
        fooBalance -= foo[i].balance
        barBalance -= bar[i].balance
    }

    diff := math.Abs(float64(fooBalance)) + math.Abs(float64(barBalance))
    if diff != 0 {
        t.Errorf("detected difference of %f", diff)
    }
}

func BenchmarkExecuteTransfers(b *testing.B) {
    foo, bar := createAccounts(ACCOUNTS, BALANCE)
    for i := 0; i < b.N; i++ {
        ExecuteTransfers(foo, bar, AMOUNT)
    }
}

func createAccounts(n, balance int) ([]*Account, []*Account) {
    a := make([]*Account, n)
    b := make([]*Account, n)
    for i := 0; i < ACCOUNTS; i++ {
        a[i] = &Account{owner: OWNER_A, number: i, balance: balance}
        b[i] = &Account{owner: OWNER_B, number: i, balance: balance}
    }
    return a, b
}
