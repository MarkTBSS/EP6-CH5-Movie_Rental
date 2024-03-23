package rental

import "fmt"

type Customer struct {
	name    string
	rentals []Rental
}

func NewCustomer(name string) (rcvr Customer) {
	rcvr = Customer{
		name:    name,
		rentals: []Rental{},
	}
	return rcvr
}
func (rcvr Customer) AddRental(arg Rental) {
	rcvr.rentals = append(rcvr.rentals, arg)
}
func (rcvr Customer) Name() string {
	return rcvr.name
}
func (rental Rental) Charge() float64 {
	result := 0.0
	switch rental.Movie().PriceCode() {
	case REGULAR:
		result += 2
		if rental.DaysRented() > 2 {
			result += float64(rental.DaysRented()-2) * 1.5
		}
	case NEW_RELEASE:
		result += float64(rental.DaysRented()) * 3.0
	case CHILDRENS:
		result += 1.5
		if rental.DaysRented() > 3 {
			result += float64(rental.DaysRented()-3) * 1.5
		}
	}
	return result
}

func getPoints(rental Rental) int {
	if rental.Movie().PriceCode() == NEW_RELEASE && rental.DaysRented() > 1 {
		return 2
	}
	return 1
}
func (rcvr Customer) Statement() string {
	totalAmount := 0.0
	frequentRenterPoints := 0
	result := fmt.Sprintf("Rental Record for %v\n", rcvr.Name())
	for _, rental := range rcvr.rentals {
		frequentRenterPoints += getPoints(rental)
		result += fmt.Sprintf("\t%v\t%.1f\n", rental.Movie().Title(), rental.Charge())
		totalAmount += rental.Charge()
	}
	result += fmt.Sprintf("Amount owed is %.1f\n", totalAmount)
	result += fmt.Sprintf("You earned %v frequent renter points", frequentRenterPoints)
	return result
}
