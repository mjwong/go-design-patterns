package builder

// CarBuilder is a
type CarBuilder struct {
	v VehicleProduct
}

// SetWheels is a
func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c
}

// SetSeats is a
func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = 5
	return c
}

// SetStructure is a
func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Car"
	return c
}

// GetVehicle is a
func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}

// Build is a
func (c *CarBuilder) Build() VehicleProduct {
	return VehicleProduct{}
}
