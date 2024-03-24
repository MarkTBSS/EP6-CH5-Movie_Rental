package rental

const (
	_ = iota
	CHILDRENS
	NEW_RELEASE
	REGULAR
)

type Charger interface {
	Charge(daysRented int) float64
	PriceCode() int
}

type Childrens struct {
}

func (children Childrens) Charge(daysRented int) float64 {
	result := 1.5
	if daysRented > 3 {
		result += float64(daysRented-3) * 1.5
	}
	return result
}

type Movie struct {
	title     string
	priceCode int
	Charger   Charger
}

func NewMovie(title string, priceCode int) (movie Movie) {
	movie = Movie{
		title:     title,
		priceCode: priceCode,
	}
	return movie
}
func (movie Movie) PriceCode() int {
	return movie.priceCode
}
func (movie Movie) Title() string {
	return movie.title
}
