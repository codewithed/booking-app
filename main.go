package main

import (
	"fmt"
)

type booking struct {
	firstName   string
	lastName    string
	userTickets uint
	email       string
}

const noOfTickets uint = 100

var conferenceName string = "Go Conference"
var remainingTickets uint = 100
var bookings = make([]booking, 0)

func main() {

	greetUsers()
	firstName, lastName, usertickets, email := getUserInput()

	bookTicket(firstName, lastName, usertickets, email)

	fmt.Printf("%v tickets remaining\n", remainingTickets)

}

func greetUsers() {
	fmt.Printf("Welcome to %v \n", conferenceName)
	fmt.Printf("We have %v left \n", remainingTickets)
	fmt.Println("Get your tickets to attend")
}

func getUserInput() (string, string, uint, string) {
	//ask user for input
	var firstName string
	var lastName string
	var userTickets uint
	var email string

	fmt.Println("Please enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("How many tickets are you buying? ")
	fmt.Scan(&userTickets)

	fmt.Println("Please enter your email address: ")
	fmt.Scan(&email)

	return firstName, lastName, userTickets, email
}

func bookTicket(firstName string, lastName string, userTickets uint, email string) {
	remainingTickets -= 1

	var newBOOKING = booking{
		firstName:   firstName,
		lastName:    lastName,
		userTickets: userTickets,
		email:       email,
	}
	bookings = append(bookings, newBOOKING)

	fmt.Printf("Thanks %v %v for buying %v tickets\n", firstName, lastName, userTickets)
}
