package builder

// BikeBuilder is a
type BikeBuilder struct {
	v VehicleProduct
}

// SetWheels is a
func (b *BikeBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 2
	return b
}

// SetSeats is a
func (b *BikeBuilder) SetSeats() BuildProcess {
	b.v.Seats = 2
	return b
}

// SetStructure is a
func (b *BikeBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Motorbike"
	return b
}

// GetVehicle is a
func (b *BikeBuilder) GetVehicle() VehicleProduct {
	return b.v
}

// Build is a
func (b *BikeBuilder) Build() VehicleProduct {
	return VehicleProduct{}
}
