package objects

type Bag struct {
	amount     float64
	invitation *Invitation
	ticket     Ticket
}

func (b Bag) HasInvitation() bool {
	return b.invitation != nil
}

func (b *Bag) SetTicket(ticket Ticket) {
	b.ticket = ticket
}

func (b *Bag) MinusAmount(amount float64) {
	b.amount -= amount
}

func (b *Bag) PlusAmount(amount float64) {
	b.amount += amount
}

func (b *Bag) GetAmount() float64 {
	return b.amount
}

func NewBag(amount float64) (*Bag, error) {
	return &Bag{amount: amount}, nil
}

func NewBagInvitation(amount float64, invitation *Invitation) (*Bag, error) {
	return &Bag{invitation: invitation, amount: amount}, nil
}
