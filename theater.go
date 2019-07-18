package objects

type Theater struct {
	ticketSeller *TicketSeller
}

func (t *Theater) Enter(audience *Audience) {
	if audience.GetBag().HasInvitation() {
		ticket := t.ticketSeller.GetTicketOffice().GetTicket()
		audience.GetBag().SetTicket(ticket)
	} else {
		ticket := t.ticketSeller.GetTicketOffice().GetTicket()
		audience.GetBag().MinusAmount(ticket.getFee())
		t.ticketSeller.GetTicketOffice().PlusAmount(ticket.getFee())
		audience.GetBag().SetTicket(ticket)
	}
}

func (t *Theater) GetRestTickets() int {
	return t.ticketSeller.GetTicketOffice().GetRestTickets()
}

func NewTheater(ticketSeller *TicketSeller) (*Theater, error) {
	return &Theater{
		ticketSeller: ticketSeller,
	}, nil
}
