package main

import "fmt"

// Athlete is a
type Athlete struct{}

// Train is a
func (a *Athlete) Train() {
	fmt.Println("Training")
}

// CompositeSwimmerA is a
type CompositeSwimmerA struct {
	MyAthlete Athlete
	MySwim    func()
}

// Swim is a
func Swim() {
	fmt.Println("Swimming!")
}

// Animal is a
type Animal struct{}

// Eat is a
func (r *Animal) Eat() {
	println("Eating")
}

// Shark is a
type Shark struct {
	Animal // embedded object
	Swim   func()
}

// Alternate way

// Swimmer is a
type Swimmer interface {
	Swim()
}

// Trainer is a
type Trainer interface {
	Train()
}

// SwimmerImpl is a
type SwimmerImpl struct{}

// Swim is a
func (s *SwimmerImpl) Swim() {
	println("Swimming!")
}

// CompositeSwimmerB is a
type CompositeSwimmerB struct {
	Trainer
	Swimmer
}

func main() {
	swimmer := CompositeSwimmerA{
		MySwim: Swim,
	}

	swimmer.MyAthlete.Train()
	swimmer.MySwim()

	fish := Shark{
		Swim: Swim,
	}

	// call the Eat method of the Animal struct without creating it
	// or using the intermediate field name.
	fish.Eat()
	fish.Swim()

	// The Swimmer field is embedded, but won't be zero-initialized
	// as it is a pointer to an interface.
	swimmerB := CompositeSwimmerB{
		&Athlete{},
		&SwimmerImpl{},
	}

	swimmerB.Train()
	swimmerB.Swim()

}
