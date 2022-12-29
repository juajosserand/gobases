package ticket

const (
	Dawning   = "dawning"
	Morning   = "morning"
	Afternoon = "afternoon"
	Night     = "night"
)

type Ticket struct {
	ID                 uint64
	Name               string
	Email              string
	DestinationCountry string
	DepartureTime      string
	Price              float64
}
