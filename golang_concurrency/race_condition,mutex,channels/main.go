package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func main() {
	//variable for bank balance
	var bankBalance int
	var balance sync.Mutex
	//starting values
	fmt.Printf("initial account balance:%d.00", bankBalance)
	fmt.Println()

	//weekly revenue
	incomes := []Income{
		{"mainjob", 500},
		{"gift", 10},
		{"parttimejob", 50},
		{"investment", 100},
	}
	wg.Add(len(incomes))
	//loop through 52 weks and print how much is maade and total
	for i, income := range incomes {

		go func(i int, income Income) {
			defer wg.Done()
			for week := 1; week <= 52; week++ {
				balance.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				balance.Unlock()
				fmt.Printf("on week %d , you earned %d.00 from %s\n", week, income.Amount, income.Source)
			}

		}(i, income)

		//printout fianl balance

	}
	wg.Wait()
	fmt.Printf("final amount=%d.00", bankBalance)
	fmt.Println()
}

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var msg string
// var wg sync.WaitGroup

// func updateMessage(s string, m *sync.Mutex) {
// 	defer wg.Done()
// 	m.Lock()

// 	msg = s
// 	m.Unlock()

// }

// func main() {
// 	var mutex sync.Mutex
// 	msg = "hello world"
// 	wg.Add(2)
// 	go updateMessage("hello universe", &mutex)
// 	go updateMessage("hello cosmos", &mutex)
// 	wg.Wait()
// 	fmt.Println(msg)

// }
