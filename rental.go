package rental

type Rental struct {
	movie      Movie
	daysRented int
}

func NewRental(movie Movie, daysRented int) (rcvr Rental) {
	rcvr = Rental{
		movie:      movie,
		daysRented: daysRented,
	}
	return rcvr
}
func (rental Rental) DaysRented() int {
	return rental.daysRented
}
func (rental Rental) Movie() Movie {
	return rental.movie
}
