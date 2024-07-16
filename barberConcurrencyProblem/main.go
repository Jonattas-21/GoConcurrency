package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

type barbershop struct {
	ShopCapacity    int
	HairCutDuration time.Duration
	NumberOfBarbers int
	BarberDone      chan bool
	ClientsChan     chan string
	Open            bool
}

func (shop *barbershop) addBarber(barber string) {
	shop.NumberOfBarbers++

	go func() {
		isSleeping := false
		color.Cyan("%s goes to the waiting room check for clients.", barber)

		for {
			//if there is no clients the barber will sleep
			if len(shop.ClientsChan) == 0 {
				color.Cyan("%s goes to sleep, there is nothing to do", barber)
				isSleeping = true
			}

			client, shopOpen := <-shop.ClientsChan

			if shopOpen {
				if isSleeping {
					color.Cyan("%s wakes up %s", client, barber)
					isSleeping = false
				}
				//cut hair
				shop.cutHair(barber, client)
			} else {
				//shop is close send barber to home
				shop.sendBarberHome(barber)
				return
			}
		}
	}()
}

func (shop *barbershop) cutHair(barber string, client string) {
	color.Cyan("%s is cutting %s hair\n", barber, client)
	time.Sleep(shop.HairCutDuration)
	color.Cyan("%s is done cutting %s hair", barber, client)
}

func (shop *barbershop) sendBarberHome(barber string) {
	color.Red("%s is going home", barber)
	shop.BarberDone <- true
}

func (shop *barbershop) addClient(client string) {
	color.Green("**** %s enters the shop", client)

	if shop.Open {
		select {
		case shop.ClientsChan <- client:
			color.Yellow("%s takes a seat in the waiting room", client)
		default:
			color.Red("%s leaves the shop, there are no seats available", client)
		}
	} else {
		color.Red("%s leaves the shop, it is closed", client)
	}
}

func (shop *barbershop) closeShopForDay() {
	color.Cyan("The shop is closing for the day")
	close(shop.ClientsChan)
	shop.Open = false

	for a := 0; a < shop.NumberOfBarbers; a++ {
		<-shop.BarberDone
	}

	close(shop.BarberDone)
	color.Green("--------------------------------")
	color.Green("The shop is closed for the day, and everyone has gone home.")
}

var seatscapacity = 4
var arrivalRate = 100
var cutDuration = 1000 * time.Millisecond
var timeOpen = 10 * time.Second

func main() {

	//seed our number generartor
	rand.Seed(time.Now().UnixNano())

	//print wellcome message
	color.Yellow("Welcome to the barbershop problem")
	color.Yellow("--------------------------------")

	//create chennels
	clienthannel := make(chan string, seatscapacity)
	doneChannel := make(chan bool)

	//create the barbershop
	shop := barbershop{
		ShopCapacity:    seatscapacity,
		HairCutDuration: cutDuration,
		NumberOfBarbers: 0,
		BarberDone:      doneChannel,
		ClientsChan:     clienthannel,
		Open:            true,
	}

	color.Green("The shop is open for the day")

	//add barbers
	shop.addBarber("Fausto")
	shop.addBarber("Luis")
	shop.addBarber("Juan")
	shop.addBarber("Raul")
	shop.addBarber("Pedro")

	//start the barbershop as a go routine
	shopClosing := make(chan bool)
	close := make(chan bool)

	go func() {
		<-time.After(timeOpen)
		shopClosing <- true
		shop.closeShopForDay()
		close <- true
	}()

	//add clients
	i := 1

	go func() {
		for {
			//get a random with avaredge arrival rate
			randomMillisecond := rand.Int() % (2 * arrivalRate)

			select {
			case <-shopClosing:
				return
			case <-time.After(time.Millisecond * time.Duration(randomMillisecond)):
				shop.addClient(fmt.Sprintf("Client %d", i))
				i++
			}
		}
	}()

	//block until the barbershop is close
	<-close
}
