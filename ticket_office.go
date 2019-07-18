package objects

type TicketOffice struct {
	amount  float64
	tickets []Ticket
}

func (to *TicketOffice) GetTicket() Ticket {
	var ticket Ticket
	ticket, to.tickets = to.tickets[len(to.tickets)-1], to.tickets[:len(to.tickets)-1]
	return ticket
}

func (to TicketOffice) GetRestTickets() int {
	return len(to.tickets)
}

func (to TicketOffice) MinusAmount(amount float64) {
	to.amount -= amount
}

func (to TicketOffice) PlusAmount(amount float64) {
	to.amount += amount
}

func NewTicketOffice(amount float64, tickets []Ticket) (*TicketOffice, error) {
	return &TicketOffice{
		amount:  amount,
		tickets: tickets,
	}, nil
}
