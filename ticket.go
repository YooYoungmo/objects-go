package objects

type Ticket struct {
	fee float64
}

func (t Ticket) getFee() float64 {
	return t.fee
}
