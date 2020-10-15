package prototype

import "fmt"

// ItemInfoGetter is a
type ItemInfoGetter interface {
	GetInfo() string
}

// ShirtColor is a
type ShirtColor byte

// Shirt is a
type Shirt struct {
	Price float32
	SKU   string
	Color ShirtColor
}

// GetInfo is a
func (s *Shirt) GetInfo() string {
	return fmt.Sprintf("Shirt with SKU '%s' and Color id %d that costs %f\n", s.SKU, s.Color, s.Price)
}

// GetPrice is a
func (s *Shirt) GetPrice() float32 {
	return s.Price
}

var whitePrototype = &Shirt{
	Price: 15.00,
	SKU:   "empty",
	Color: White,
}

var blackPrototype = &Shirt{
	Price: 16.00,
	SKU:   "empty",
	Color: Black,
}

var bluePrototype = &Shirt{
	Price: 17.00,
	SKU:   "empty",
	Color: Blue,
}
