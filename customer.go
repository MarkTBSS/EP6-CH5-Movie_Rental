package rental

import (
	"fmt"
)

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
func (customer Customer) AddRental(arg Rental) {
	customer.rentals = append(customer.rentals, arg)
}
func (rcvr Customer) Name() string {
	return rcvr.name
}

func RegularCharge(rental Rental) float64 {
	result := 2.0
	if rental.DaysRented() > 2 {
		result += float64(rental.DaysRented()-2) * 1.5
	}
	return result
}

func NewReleaseCharge(rental Rental) float64 {
	return float64(rental.DaysRented()) * 3.0
}

func ChildrenCharge(rental Rental) float64 {
	result := 1.5
	if rental.DaysRented() > 3 {
		result += float64(rental.DaysRented()-3) * 1.5
	}
	return result
}
func (rental Rental) Charge() float64 {
	switch rental.Movie().PriceCode() {
	case REGULAR:
		return RegularCharge(rental)
	case NEW_RELEASE:
		return NewReleaseCharge(rental)
	case CHILDRENS:
		return ChildrenCharge(rental)
	}
	return 0
}

func getPoints(rental Rental) int {
	if rental.Movie().PriceCode() == NEW_RELEASE && rental.DaysRented() > 1 {
		return 2
	}
	return 1
}

func getTotalPoints(rentals []Rental) int {
	points := 0
	for _, rental := range rentals {
		points += getPoints(rental)
	}
	return points
}

func getTotalAmount(rentals []Rental) float64 {
	result := 0.0
	for _, rental := range rentals {
		result += rental.Charge()
	}
	return result
}

func (customer Customer) Statement() string {
	frequentRenterPoints := getTotalPoints(customer.rentals)
	totalAmount := getTotalAmount(customer.rentals)
	result := fmt.Sprintf("Rental Record for %v\n", customer.Name())
	for _, rental := range customer.rentals {
		title := rental.Movie().Title()
		amount := rental.Charge()
		result += fmt.Sprintf("\t%v\t%.1f\n", title, amount)
	}
	result += fmt.Sprintf("Amount owed is %.1f\n", totalAmount)
	result += fmt.Sprintf("You earned %v frequent renter points", frequentRenterPoints)
	return result
}
