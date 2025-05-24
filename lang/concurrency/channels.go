package concurrency

import (
	"sync"
	"time"
)

func nums1(channel chan int) {
	channel <- 1
	channel <- 2
	channel <- 3
	close(channel)
}

func nums2(channel chan int) {
	channel <- 4
	channel <- 5
	channel <- 6
	close(channel)
}

func SimpleChannelsExamples() {
	channel1 := make(chan int)
	channel2 := make(chan int)

	go nums1(channel1)
	go nums2(channel2)

	// loop 6 times
	for i := 0; i < 3; i++ {
		pl(<-channel1)
		pl(<-channel2)
	}
}

type Account struct {
	balance int
	lock    sync.Mutex
}

func (a *Account) Deposit(amount int) {
	a.lock.Lock()
	defer a.lock.Unlock()
	a.balance += amount
	pl("Deposited", amount, "to the account")
}

func (a *Account) Withdraw(amount int) {
	a.lock.Lock()
	defer a.lock.Unlock()
	if a.balance < amount {
		pl("Not enough money in the account")
		return
	}
	a.balance -= amount
	pl("Withdrew", amount, "from the account")
}

func (a *Account) GetBalance() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	pl("Balance:", a.balance)
	return a.balance
}

func ChannelsWithMutexExamples() {
	account := Account{balance: 1000}

	go account.Deposit(100)
	go account.Withdraw(200)
	go account.GetBalance()

	for i := 0; i < 100; i++ {
		go account.Withdraw(200)
	}

	time.Sleep(1 * time.Second)
}
