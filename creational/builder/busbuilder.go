package builder

// BusBuilder is a
type BusBuilder struct {
	v VehicleProduct
}

// SetWheels is a
func (b *BusBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 4 * 2
	return b
}

// SetSeats is a
func (b *BusBuilder) SetSeats() BuildProcess {
	b.v.Seats = 30
	return b
}

// SetStructure is a
func (b *BusBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Bus"
	return b
}

// GetVehicle is a
func (b *BusBuilder) GetVehicle() VehicleProduct {
	return b.v
}

// Build is a
func (b *BusBuilder) Build() VehicleProduct {
	return VehicleProduct{}
}
