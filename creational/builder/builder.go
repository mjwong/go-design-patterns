package builder

// VehicleProduct is a
type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

// BuildProcess is a
type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}
