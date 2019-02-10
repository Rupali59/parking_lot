package ticketing_system

type ParkingMap struct {
	Entry []ParkingMapEntry
}

type ParkingMapEntry struct {
	SlotNumber int
	RegistrationNumber string
	Color string
}

