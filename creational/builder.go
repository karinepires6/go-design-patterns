package creational

type ManufacturingDirector struct {
	builder BuildProcess
}

type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

func (f *ManufacturingDirector) Construct() {
	f.builder.SetSeats().SetStructure().SetWheels()
}

func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}

type CarBuilder struct {
	vehicle VehicleProduct
}

func (c *CarBuilder) Build() VehicleProduct {
	return VehicleProduct{}
}

func (c *CarBuilder) SetWheels() BuildProcess {
	c.vehicle.Wheels = 4
	return c
}

func (c *CarBuilder) SetSeats() BuildProcess {
	c.vehicle.Seats = 5
	return c
}

func (c *CarBuilder) SetStructure() BuildProcess {
	c.vehicle.Structure = "Car"
	return c
}

func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.vehicle
}

type BikeBuilder struct {
	vehicle VehicleProduct
}

func (b *BikeBuilder) Build() VehicleProduct {
	return VehicleProduct{}
}

func (b *BikeBuilder) SetWheels() BuildProcess {
	b.vehicle.Wheels = 2
	return b
}

func (b *BikeBuilder) SetSeats() BuildProcess {
	b.vehicle.Seats = 2
	return b
}

func (b *BikeBuilder) SetStructure() BuildProcess {
	b.vehicle.Structure = "Motorbike"
	return b
}

func (b *BikeBuilder) GetVehicle() VehicleProduct {
	return b.vehicle
}
