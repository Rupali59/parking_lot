package ticketing_system

//ParkingMap holds a hash like structure to store the occupied slots as well as the details of the vehicles being stored
type ParkingMap struct {
	Entry []ParkingMapEntry
}

//NewParkingMap creates a new parking map
func NewParkingMap() *ParkingMap {
	return &ParkingMap{
		Entry: []ParkingMapEntry{},
	}
}

type ParkingMapEntry struct {
	SlotNumber         int //SlotNumber as visible to user of the system
	RegistrationNumber string //RegistrationNumber of the vehicle parked at the slot number
	Color              string //Color of the vehicle parked at the slot number
}
