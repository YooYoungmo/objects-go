package objects

type Audience struct {
	bag *Bag
}

func NewAudience(bag *Bag) (*Audience, error) {
	return &Audience{
		bag: bag,
	}, nil
}

func (a Audience) GetBag() *Bag {
	return a.bag
}
