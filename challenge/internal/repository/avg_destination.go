package repository

func (r *repository) AverageDestination(destination string) (float64, error) {
	totalDest, err := r.GetTotalTickets(destination)
	if err != nil {
		return float64(totalDest), err
	}

	return float64(totalDest) / float64(len(r.Tickets)), nil
}
