package main

import (
	"time"

	color "github.com/fatih/color"
)

type BarberShop struct {
	ShopCapacity    int
	HairCutDuration time.Duration
	NumberOfBabers  int
	BarberDoneChan  chan bool
	ClientsChan     chan string
	Open            bool
}

func (shop *BarberShop) addClient(client string) {

	color.Green("***%s arrives,client")
	if shop.Open {
		select {
		case shop.ClientsChan <- client:
			color.Blue("%s takes a seat in waiting room", client)
		default:
			color.Red("The waiting room is full so %s leaves ", client)
		}
	} else {
		color.Red("The shop is already closes ,so %s leaves !", client)
	}
}

func (shop *BarberShop) ClosedShopforDay() {
	color.Blue("closing shop for day")
	close(shop.ClientsChan)
	shop.Open = false
	for a := 1; a <= shop.NumberOfBabers; a++ {
		<-shop.BarberDoneChan

	}
	close(shop.BarberDoneChan)
	color.Green("__________________________")
	color.Green("the barber shop is now close for the day ")

}

func (shop *BarberShop) addBarber(barber string) {
	shop.NumberOfBabers++
	go func() {
		isSleeping := false
		color.Yellow("%s goes to waiting room to check for clients,", barber)
		for {
			if len(shop.ClientsChan) == 0 {
				color.Yellow("theree is nothing to do ,so %s takes nap", barber)
				isSleeping = true
			}
			client, shopOpen := <-shop.ClientsChan
			if shopOpen {
				if isSleeping {
					color.Yellow("%s wakes %s up", client, barber)
					isSleeping = false
				}
				shop.cutHair(barber, client)

			} else {
				shop.sendBrberHome(barber)
				return

			}

		}

	}()
}
func (shop *BarberShop) cutHair(barber, client string) {
	color.Green("%s is cutting %s's hair", barber, client)
	time.Sleep(shop.HairCutDuration)
}

func (shop *BarberShop) sendBrberHome(barber string) {
	color.Blue("%s is going home", barber)
	shop.BarberDoneChan <- true
}
