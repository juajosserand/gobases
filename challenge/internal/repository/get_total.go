package repository

func (r *repository) GetTotalTickets(destination string) (total int, err error) {
	if destination == "" {
		return total, errEmptyDestination
	}

	for _, t := range r.Tickets {
		if t.DestinationCountry == destination {
			total++
		}
	}

	return total, nil
}
