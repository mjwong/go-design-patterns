package factory

import (
	"errors"
	"fmt"
)

// Delegating the creation of new instances of structures to a different part of the program
// Working at the interface level instead of with concrete implementations
// Grouping families of objects to obtain a family object creator.

// PaymentMethod is a
type PaymentMethod interface {
	Pay(amount float32) string
}

const (
	// Cash is a cash payment method
	Cash = 1
	// DebitCard is a debit card payment method
	DebitCard = 2
)

// CashPM is a type of payment method
type CashPM struct{}

// DebitCardPM is a type of payment method
type DebitCardPM struct{}

// CreditCardPM is a type of payment method
type CreditCardPM struct{}

// Pay methtod implements the PaymentMethod interface
func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using cash\n", amount)
}

// Pay methtod implements the PaymentMethod interface
func (c *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%#0.2f paid using debit card\n", amount)
}

// Pay methtod implements the PaymentMethod interface
func (d *CreditCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%#0.2f paid using debit card (new)\n", amount)
}

// GetPaymentMethod accepts a payment type as arg.
// The preceding code is the function that will create the objects for us.
// It returns a pointer, which must have an object that implements the PaymentMethod interface,
// and an error if asked for a method which is not registered.
func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(CreditCardPM), nil // changed to new CreditCardPM
	default:
		return nil, errors.New(fmt.Sprintf("Payment method %d not recognized\n", m))
	}
}
