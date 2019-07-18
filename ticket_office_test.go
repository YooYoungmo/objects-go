package objects

import (
	"gotest.tools/assert"
	"testing"
)

func TestGetTicket(t *testing.T) {
	// given
	tickets := []Ticket{
		{
			fee: 10.0,
		},
		{
			fee: 20.0,
		},
	}
	ticketOffice, _ := NewTicketOffice(100.0, tickets)

	// when
	ticket := ticketOffice.GetTicket()

	// then
	assert.Equal(t, ticket.fee, 20.0)
	assert.Equal(t, ticketOffice.GetRestTickets(), 1)
}
