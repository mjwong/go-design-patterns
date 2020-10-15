package builder

// ManufacturingDirector is a
type ManufacturingDirector struct {
	builder BuildProcess
}

// SetBuilder is a
func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}

// Construct is a
func (f *ManufacturingDirector) Construct() {
	f.builder.SetSeats().SetStructure().SetWheels()
}
