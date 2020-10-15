package prototype

import "errors"

// Maintain a set of objects that will be cloned to create new instances
// Provide a default value of some type to start working on top of it
// Free CPU of complex object initialization to take more memory resources

// ShirtCloner is a
type ShirtCloner interface {
	GetClone(s int) (ItemInfoGetter, error)
}

const (
	// White shirt
	White = 1
	// Black shirt
	Black = 2
	// Blue shirt
	Blue = 3
)

// GetShirtsCloner retrieves a new instance of the cloner
func GetShirtsCloner() ShirtCloner {
	return &ShirtsCache{}
}

// ShirtsCache is a
type ShirtsCache struct{}

// GetClone is a
func (s *ShirtsCache) GetClone(m int) (ItemInfoGetter, error) {
	switch m {
	case White:
		newItem := *whitePrototype
		return &newItem, nil
	case Black:
		newItem := *blackPrototype
		return &newItem, nil
	case Blue:
		newItem := *bluePrototype
		return &newItem, nil
	default:
		return nil, errors.New("Shirt model not recognized")
	}
}
