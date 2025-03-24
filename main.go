package main

import (
	"fmt"
	"strings"
	"sync"
)


type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}


var wg = sync.WaitGroup{}
var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)			// slice(vector) of struct


func main() {
	greetUser(conferenceName, int(remainingTickets))

	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName := len(firstName) > 2 && len(lastName) > 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(firstName, lastName, userTickets, email)

			wg.Add(1)
			go sendTicket(firstName, lastName, userTickets, email)	 // this creates a new thread to perform this func

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name and last name length should be greater than 2.")
			}
			
			if !isValidEmail {
				fmt.Println("Invalid Email")
			}
			
			if !isValidTicketNumber {
				fmt.Printf("We only have %v tickets available. So you can't book %v tickets.", remainingTickets, userTickets)
			}
		}
	}

	firstNames := getUsers(bookings)
	fmt.Printf("These are all the bookings: %v", firstNames)

	wg.Wait()
}