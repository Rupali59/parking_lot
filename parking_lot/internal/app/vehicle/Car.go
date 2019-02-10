package vehicle

type CarFactory struct {

}

func (carFactory *CarFactory)CreateVehicle(registrationNumber string, color string) *Vehicle{
	return &Vehicle{
		RegistrationNumber:registrationNumber,
		Color:color,
		Type: "Car",
	}
}
