package main

import (
	"fmt"
	"strings"
	"reservations-app/common"
)

const concertTickets = 50
var concertName = "Go Concert"
var remainingTickets uint = 50
var bookings = []string {} //Declaring an array (slice) to store the names of booking clients

func main() {
	// Call greetUsers function and include parameters
	greetUsers()

	for {
		// Call function for getting user input
		firstName, lastName, email, userTickets := getUserInput()

		// Call function for validating user input
		isValidName, isValidEmail, isValidTicketNumber := common.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets) 

			// BOOKING TICKETS!
			if isValidName && isValidEmail && isValidTicketNumber {
				
				// Call function for booking tickets
				bookTicket(userTickets, firstName, lastName, email)

				// Call function to print first names
				firstNames := getFirstNames()
				fmt.Printf ("The first names of bookings are: %v\n", firstNames)

							// exit application if no tickets are left
							noTicketsRemaining := remainingTickets == 0
							if noTicketsRemaining {
								//end programme by using break
								fmt.Println ("All tickets soldout")
								break
							}

			} else {
				if !isValidName {
					fmt.Println("First name or last name you entered is too short")
				}
				if !isValidEmail {
					fmt.Println("Email address you entered doesn't contain @ sign")
				}
				if !isValidTicketNumber {
					fmt.Println("Number of tickets you entered is invalid")
				}
				continue
			}
	}
}

func greetUsers() {
	fmt.Printf ("You are now in the %v bookings.\n", concertName)
	fmt.Printf ("We have the total of %v tickets, to which %v are still available.\n", concertTickets, remainingTickets)
	fmt.Println ("Please book your tickets here.")
} 

// Function for getting first names of concert goers
func getFirstNames() []string {
					//Print only first names
					firstNames := []string {}
					// For arrays and slices, range provides undex and value for each element we iterate through.
					// Range allows for iterating through different data structures - not only arrays and slices.
					// Use undescore as a blank identifier as we dont need index value
					for _, booking := range bookings {
						//strings.Fields splits the booking string with white space separator
						var names = strings.Fields(booking)
						firstNames = append(firstNames, names[0])
					}
					return firstNames
}

// Function for getting user input
func getUserInput () (string, string, string,uint) {
	var firstName string
	var lastName string
	var email string 
	var userTickets uint

	//Asking user for their name (user input)
	fmt.Println ("Enter your first name:")
	fmt.Scan (&firstName)

	fmt.Println ("Enter your last name:")
	fmt.Scan (&lastName)

	fmt.Println ("Enter your email address name:")
	fmt.Scan (&email)

	fmt.Println ("Enter number of tickets:")
	fmt.Scan (&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	//Logic for booking tickets
	remainingTickets = concertTickets - userTickets

	/*The bookings slice will consist of the following:
	This will dynamically add (append) entries to the slice without having to specify min/max or index position.
	*/
	bookings = append(bookings, firstName + " " + lastName)
	//Lets see whats inside the bookings array
	fmt.Printf ("Thank you %v %v for buying %v tickets, confirmation will be sent to %v soon.\n", firstName, lastName, userTickets, email)
	fmt.Printf ("Only %v remaining tickets for %v\n", remainingTickets, concertName)
}


