package main

import (
	"fmt"
	"time"
)

func greetUser(conferenceName string, remainingTickets int) {
	fmt.Println("Welcome to", conferenceName, "booking app!")
	fmt.Println("Get your tickets here. Only", remainingTickets, "tickets are available.")
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email: ")
	fmt.Scan(&email)
	fmt.Println("Enter the no. of seats youy want to book: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(firstName string, lastName string, userTickets uint, email string) {
	remainingTickets -= userTickets;

	var userdata = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userdata)

	fmt.Printf("Thanks you %v %v for booking %v tickets. You will recieve a confirmation mail at %v\n", firstName, lastName, userTickets, email)
	
}

func sendTicket(firstName string, lastName string, userTickets uint, email string) {
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Printf("Sending ticket:\n%v\nto email address %v", ticket, email)
	wg.Done()
}

func getUsers(bookings []UserData) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}