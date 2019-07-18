package objects

type TicketSeller struct {
	ticketOffice *TicketOffice
}

func (ts TicketSeller) GetTicketOffice() *TicketOffice {
	return ts.ticketOffice
}

func NewTicketSeller(ticketOffice *TicketOffice) (*TicketSeller, error) {
	return &TicketSeller{
		ticketOffice: ticketOffice,
	}, nil
}
