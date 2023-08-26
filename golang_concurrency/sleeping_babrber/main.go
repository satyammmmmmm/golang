package main

import (
	"fmt"
	"math/rand"
	"time"

	color "github.com/fatih/color"
)

// variables
var seatingCapacity = 10
var arrivalRate = 100
var cutDuration = 1000 * time.Millisecond
var timeOpen = 10 * time.Second

func main() {
	// sedd our randon no, generator
	rand.Seed(time.Now().UnixNano())

	// welcome messgae
	color.Yellow("the sleeping barber problem")
	color.Yellow("****************")

	//create channels if we need
	clientChan := make(chan string, seatingCapacity)
	doneChan := make(chan bool)

	//create barbershop
	shop := BarberShop{
		seatingCapacity,
		cutDuration,
		0,
		doneChan,
		clientChan,
		true,
	}
	color.Green("the shop is open for day")

	//add barbers
	shop.addBarber("Frank")

	//strt barbershop as goroutine
	shopClosing := make(chan bool)
	closed := make(chan bool)
	go func() {
		<-time.After(timeOpen)
		shopClosing <- true
		shop.ClosedShopforDay()
		closed <- true

	}()

	// add clients
	i := 1
	go func() {
		for {
			randomMilliseconds := rand.Int() % (2 * arrivalRate)
			select {
			case <-shopClosing:
				return
			case <-time.After(time.Millisecond * time.Duration(randomMilliseconds)):
				shop.addClient(fmt.Sprinf("client %d", i))
			}

		}
	}()

	// block until barbershop is closed
	<-closed
}

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func server1(ch chan string) {
// 	for {
// 		time.Sleep(6 * time.Second)
// 		ch <- "this is from server1"
// 	}
// }

// func server2(ch chan string) {
// 	time.Sleep(3 * time.Second)
// 	ch <- "this is from server2"
// }

// func main() {
// 	ch1 := make(chan string)
// 	ch2 := make(chan string)
// 	go server1(ch1)
// 	go server2(ch2)
// 	for {
// 		select {
// 		case s1 := <-ch1:
// 			fmt.Println(s1)
// 		case s2 := <-ch2:
// 			fmt.Println(s2)
// 		}
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func shout(ping chan string, pong chan string) {
// 	for {
// 		//fmt.Println("executing loop")
// 		s ,ok:= <-ping
// 		if !ok{

// 		}
// 		pong <- fmt.Sprintf("%s!!!", strings.ToUpper(s))
// 	}
// }

// // channel basic
// func main() {

// 	ping := make(chan string)
// 	pong := make(chan string)

// 	go shout(ping, pong)
// 	fmt.Printf("type something and press enter (enter q to quit)")
// 	for {
// 		fmt.Printf("-->")
// 		var userInput string
// 		_, _ = fmt.Scanln(&userInput)
// 		if userInput == strings.ToLower(("q")) {
// 			break
// 		}
// 		ping <- userInput
// 		response := <-pong
// 		fmt.Println("response:", response)
// 	}
// 	fmt.Println("All done closing channels")
// 	close(ping)
// 	close(pong)

// }
