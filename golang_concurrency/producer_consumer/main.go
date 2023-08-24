package main

import (
	"fmt"
	"math/rand"
	"time"

	color "github.com/fatih/color"
)

const NumberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}
type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}
func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber += 1
	if pizzaNumber <= NumberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("recieve order %d\n", pizzaNumber)
		rnd := rand.Intn(12) + 1
		msg := ""
		success := false
		if rnd < 5 {
			pizzasFailed += 1
		} else {
			pizzasMade += 1
		}
		total += 1
		fmt.Printf("making pizza %d,it will taake %d seconds", pizzaNumber, delay)
		fmt.Println()
		time.Sleep(time.Duration(delay) * time.Second)
		if rnd < 2 {
			fmt.Println()
			msg = fmt.Sprintf("***We ran out of incrediants for pizza %d\n", pizzaNumber)

		} else if rnd <= 4 {
			fmt.Println()
			msg = fmt.Sprintf("***The cook quit whilemaking pizzanumber %d\n", pizzaNumber)
		} else {
			success = true
			fmt.Println()
			msg = fmt.Sprintf("***pizza order %d is ready", pizzaNumber)

		}
		p := PizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}
		return &p

	}
	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
	}

}

func pizzeria(pizzaMaker *Producer) {
	//keep track of which pizza we are making
	var i = 0
	// run forever or until we recieve we recieve quit notification

	//try to make pizzas
	for {
		currentPizza := makePizza(i)
		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select {
			case pizzaMaker.data <- *currentPizza:

			case quitChan := <-pizzaMaker.quit:
				//close channels
				close(pizzaMaker.data)
				close(quitChan)
				return
			}

		}

		//try to make a pizza
		//decision

	}
}

func main() {

	//seed the random number generator
	rand.Seed(time.Now().UnixNano())
	//printout a mesage
	color.Cyan("The Pizzeria is open for bussiness")
	color.Cyan("..................................")

	// create a producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	//run the producer in background
	go pizzeria(pizzaJob)

	// create and run consumer

	for i := range pizzaJob.data {
		if i.pizzaNumber <= NumberOfPizzas {
			if i.success {
				color.Green(i.message)
				color.Green("order number %d is out of delivery", i.pizzaNumber)

			} else {
				color.Red(i.message)
				color.Red("the customer is really made")
			}

		} else {
			color.Cyan("done making pizas")
			err := pizzaJob.Close()
			if err != nil {
				color.Red("error cclosing channel", err)
			}
		}
	}
	color.Cyan("******************* ")
	color.Cyan("Done for the day ")
	color.Cyan("We made %d pizzas ,but failed to make %d,with %d attempt in total", pizzasMade, pizzasFailed, total)
	switch {
	case pizzasFailed > 9:
		color.Red("Bad Day")
	case pizzasFailed >= 6:
		color.Red("Average Day")
	case pizzasFailed >= 4:
		color.Yellow("okkish Day")
	case pizzasFailed >= 2:
		color.Yellow("pretty good day")
	default:
		color.Green("Wonderful day ")

	}

}
