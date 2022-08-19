package main

import (
	"booking-app/helper"
	"fmt"
	"time"
	"sync"
)

var conferenceName = "Go Conference"
const conferenceTickets = 50
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct{
	firstName string
	lastName string
	email string
	numberOfTickets uint
}
var wg= sync.WaitGroup{}
func main() {
	    greetUsers()

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTickets := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)
		if isValidName && isValidEmail && isValidTickets {

			bookTicket(userTickets, firstName, lastName, email)

			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)
			firstNames := returnFirstNames()
			fmt.Printf("First names of our bookings are: %v\n", firstNames)
			if remainingTickets == 0 {
				fmt.Println("All tickets sold out")
			}
		} else {
			if !isValidName {
				fmt.Println("first Name or Last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email ou entered is invalid")
			}

			fmt.Println("Your input data is invalid")
		}

	}
func greetUsers() {
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here!!")
}

func returnFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames

}
func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address:")
	fmt.Scan(&email)

	fmt.Println("Enter no of tickets:")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}
func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets -= userTickets

	var userData = UserData{
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}
    
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets, you will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for the conference\n", remainingTickets)
}
func sendTicket(userTickets uint,firstName string, lastName string,  email string){
	var ticket = fmt.Sprintf("%v tickets for %v %v \n",userTickets, firstName, lastName)
	time.Sleep(50* time.Second)
	fmt.Println("##################")
	fmt.Printf("sending ticket %v to email address:%v\n",ticket,email)
    fmt.Println("##################")
}