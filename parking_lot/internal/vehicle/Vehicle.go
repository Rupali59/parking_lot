package vehicle

type Vehicle struct {
	RegistrationNumber string
	Color string
	Type string
}

type VehicleFactory interface {
	CreateVehicle(registrationNumber string, color string) *Vehicle
}

func GetVehicle(vehicleType string, registrationNumber string, color string) *Vehicle{
	switch vehicleType {
	case "car": return CarFactory{}.CreateVehicle(registrationNumber,color)
	}
}