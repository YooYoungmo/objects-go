package objects

import (
	"gotest.tools/assert"
	"testing"
)

func TestEnter_Invited(t *testing.T) {
	// given
	bag, _ := NewBagInvitation(1000.0, &Invitation{})
	audience, _ := NewAudience(bag)

	tickets := []Ticket{
		{
			fee: 10.0,
		},
		{
			fee: 20.0,
		},
	}
	ticketOffice, _ := NewTicketOffice(100.0, tickets)
	ticketSeller, _ := NewTicketSeller(ticketOffice)
	theater, _ := NewTheater(ticketSeller)

	// when
	theater.Enter(audience)

	// then
	assert.Equal(t, audience.GetBag().ticket.getFee(), 20.0)
	assert.Equal(t, audience.GetBag().GetAmount(), 1000.0)
	assert.Equal(t, theater.GetRestTickets(), 1)
}

func TestEnter_Not_Invited(t *testing.T) {
	// given
	bag, _ := NewBag(1000.0)
	audience, _ := NewAudience(bag)

	tickets := []Ticket{
		{
			fee: 10.0,
		},
		{
			fee: 20.0,
		},
	}
	ticketOffice, _ := NewTicketOffice(100.0, tickets)
	ticketSeller, _ := NewTicketSeller(ticketOffice)
	theater, _ := NewTheater(ticketSeller)

	// when
	theater.Enter(audience)

	// then
	assert.Equal(t, audience.GetBag().ticket.getFee(), 20.0)
	assert.Equal(t, audience.GetBag().GetAmount(), 980.0)
	assert.Equal(t, theater.GetRestTickets(), 1)
}
