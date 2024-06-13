package main // Create a package to wrap the app code

import (
	"fmt" // The format package has the print function
	"sync"
	"time"
)

var conferenceName = "Mario Bros Conference"
const conferenceTickets = 50 
var remainingTickets uint = 50
var bookings = make([]UserData, 0) // Bookings is an empty list of UserData structs

// Create a struct
type UserData struct {
	firstName string
	lastName string	
	email string 
	numberOfTickets uint
}

// Creates a wait group
var wg = sync.WaitGroup{} 

func main(){  // The entrypoint of the main function

	greetUsers() // You dont need to provide arguments, bc these variables are acessible globally
		
	// Get the user input
	firstName, lastName, email, userTickets := getUserInput() // This function doesnt have arguments bc it uses the user input, but it has outputs

	// User Input Validation function. It comes from helper.go
	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		bookTicket(userTickets, firstName, lastName, email) 
		
		wg.Add(1) // Set the number of go routines to wait for (and increase the counter)
		go sendTicket(userTickets, firstName, lastName, email) // Adding 'go' makes the function concurrent

		// Get the first name from the booking list
		firstNames := getFirstNames() // Bookings is defined globally
		fmt.Printf("The first names of the bookings are: %v\n", firstNames)

		// End the program if the tickets are all booked
		if remainingTickets == 0 {
			fmt.Println("Our conference is sold out. Try again next year!")
		}

		} else { // If the user tries to buy more tickets than available tickets or use invalid data, the program is restarted
			if !isValidName {
				fmt.Println("Invalid first or last name, too short.")
			}
			if !isValidEmail {
				fmt.Println("The email must contain an @ sign.")
			}
			if !isValidTicketNumber{
				fmt.Println("The entered number of tickets is invalid.")
			}
			
		}
	wg.Wait() // It waits till all threads are done
}

func greetUsers() { 
	fmt.Printf("Welcome to the %v\n", conferenceName) 
	fmt.Printf("We have a total of %v tickets and %v tickets are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here!")
}

func getFirstNames() []string { // (input) output
	firstNames := []string{}
	// There is an error because you're declaring a variable that you do not use (index), but you need it because of the range function. So you just use underscore in the place of index to tell Go you know that variable is not used
	for _, booking := range bookings { // range: iterates over elements in different data structures (not only arrays and slices). For arrays and slices, range returns the index and value of each element
		firstNames = append(firstNames, booking.firstName) // Access elements in a struct
	}
	return firstNames
}

// Ask the user for name and number of tickets
func getUserInput () (string, string, string, uint) {
	var firstName string // If you dont assing a value, you need to declare the type
	var lastName string
	var email string
	var userTickets uint
	
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets // Calculations have to use the same data type
	// Create the struct and assign it to userData var
	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets, 
	}

	bookings = append(bookings, userData) // It adds maps to the list of maps (bookings)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. We will send a confirmation email to %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}

// Create and send a ticket. This is a simulation of the task, additional steps should be executed in the background
func sendTicket(userTickets uint, firstName string , lastName string, email string) {
	time.Sleep(10 * time.Second) // It stops the execution of the thread for 10s.
	// Generates a string, but dont print it, save it to another variable
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("################")
	fmt.Printf("Sending ticket:\n %v to email address %v\n", ticket, email)
	fmt.Println("################")
	wg.Done() // Remove the thread from the waiting group because its concluded, so the main thread does not have to wait for it
}