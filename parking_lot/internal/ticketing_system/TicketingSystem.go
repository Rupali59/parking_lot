package ticketing_system

import (
	"errors"
	parkinglot "../parking_lot"
	vehicle "../vehicle"
	"fmt"
)

var (
	ErrorParkingLotFull =  errors.New("Sorry, parking lot is full");
	ErrorNotFound = errors.New("Not found");
	ErrorSameRegistrationNumberFound = errors.New("Multiple Slots with same registration number found!");
	ErrorInProcessing = errors.New("Internal Server Error");
)

var (
	ResponseSlotAllocatedSuccessfully = "Allocated slot number: %d";
)

type TicketingSystem struct {
	validTickets []Ticket
	parkinglot *parkinglot.ParkingLot
	parkingMap *ParkingMap
}

func NewTicketingSystem(lot *parkinglot.ParkingLot) *TicketingSystem {
	return &TicketingSystem{
		validTickets:nil,
		parkinglot:lot,
		parkingMap:nil,
	}
}

func (system *TicketingSystem)GenerateTicket(car *vehicle.Vehicle) (response string,err error){
	slotNumber := system.parkinglot.GetEmptySlot();
	if(slotNumber == -1){
		return "",ErrorParkingLotFull
	}else {
		status, err := system.parkinglot.Park(slotNumber, car);
		if err != nil{
			return "",err
		}
		if(!status){
			return "",ErrorInProcessing
		}else{
			system.parkingMap.Entry = append(system.parkingMap.Entry, ParkingMapEntry{
				SlotNumber:slotNumber+1,
				Color:car.Color,
				RegistrationNumber:car.RegistrationNumber,
			})
			system.validTickets = append(system.validTickets, *NewTicket(slotNumber+1, car.RegistrationNumber, car.Color))
			return fmt.Sprintf(ResponseSlotAllocatedSuccessfully,slotNumber+1),nil
		}
	}
}


func (system *TicketingSystem)InvalidateTicket(slotNumber int)(status bool, err error){
	status,err = system.parkinglot.Leave(slotNumber-1)
	if(err !=nil){
		return false, err
	}
	var tickets []Ticket
	for _,ticket:= range system.validTickets{
		if(ticket.SlotNumber != slotNumber-1){
			tickets = append(tickets, ticket)
		}
	}
	system.validTickets = tickets
	var validEntry []ParkingMapEntry
	for _,entry := range system.parkingMap.Entry {
		if (entry.SlotNumber != slotNumber) {
			validEntry = append(validEntry, entry)
		}
	}
	system.parkingMap.Entry = validEntry
	return true, nil
}


func (system *TicketingSystem) GetStatus() (*ParkingMap){
	return system.parkingMap;
}

func (system TicketingSystem) GetRegistrationNumberWithColor(color string) (registrationNumbers []string, err error) {
	for _,entry := range system.parkingMap.Entry{
		if(entry.Color== color){
			registrationNumbers=append(registrationNumbers, entry.RegistrationNumber)
		}
	}
	return registrationNumbers,nil
}


func (system *TicketingSystem) GetSlotsWithColor(color string) (slotNumbers[]int, err error){
	for _,entry := range system.parkingMap.Entry{
		if(entry.Color == color){
			slotNumbers= append(slotNumbers, entry.SlotNumber)
		}
	}
	return slotNumbers,nil
}


func (system *TicketingSystem) GetSlotNumberGivenRegistrationNumber(registrationNumber string) (slotNumber int, err error){
	var slotNumbers []int;
	for _,entry := range system.parkingMap.Entry{
		if entry.RegistrationNumber== registrationNumber{
			slotNumbers = append(slotNumbers, entry.SlotNumber)
		}
	}
	switch len(slotNumbers) {
	case 0: return -1,ErrorNotFound
	case 1: return slotNumbers[0],nil
	default: return -1, ErrorSameRegistrationNumberFound
	}
}


