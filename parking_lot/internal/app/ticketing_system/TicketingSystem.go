package ticketing_system

import (
	"errors"
	parkinglot "../parkingLot"
	vehicle "../vehicle"
	"fmt"
	"strconv"
	"strings"
)

var (
	ErrorParkingLotFull =  errors.New("Sorry, parking lot is full");
	ErrorNotFound = errors.New("Not found");
	ErrorSameRegistrationNumberFound = errors.New("Multiple Slots with same registration number found!");
	ErrorInProcessing = errors.New("Internal Server Error");
)

var (
	ResponseSlotAllocatedSuccessfully = "Allocated slot number: %d";
	ResponseSlotDeallocatedSuccesfully = "Slot number %d is free"
)

type TicketingSystem struct {
	ValidTickets []Ticket
	Parkinglot *parkinglot.ParkingLot
	ParkingMap *ParkingMap
}

func NewTicketingSystem(lot *parkinglot.ParkingLot) *TicketingSystem {
	return &TicketingSystem{
		ValidTickets:[]Ticket{},
		Parkinglot:lot,
		ParkingMap: NewParkingMap(),
	}
}

func (system *TicketingSystem)GenerateTicket(car *vehicle.Vehicle) (response string,err error){
	slotNumber := system.Parkinglot.GetEmptySlot();
	if(slotNumber == -1){
		return "",ErrorParkingLotFull
	}else {
		status, err := system.Parkinglot.Park(slotNumber, car);
		if err != nil{
			return "",err
		}
		if(!status){
			return "",ErrorInProcessing
		}else{
			newEntry:=ParkingMapEntry{
				SlotNumber:slotNumber+1,
				Color:car.Color,
				RegistrationNumber:car.RegistrationNumber,
			}
			system.ParkingMap.Entry = append(system.ParkingMap.Entry, newEntry)
			system.ValidTickets = append(system.ValidTickets, *NewTicket(slotNumber+1, car.RegistrationNumber, car.Color))
			return fmt.Sprintf(ResponseSlotAllocatedSuccessfully,slotNumber+1),nil
		}
	}
}


func (system *TicketingSystem)InvalidateTicket(slotNumber int)(response string, err error){
	_,err = system.Parkinglot.Leave(slotNumber-1)
	if(err !=nil){
		return "", err
	}
	var tickets []Ticket
	for _,ticket:= range system.ValidTickets{
		if(ticket.SlotNumber != slotNumber-1){
			tickets = append(tickets, ticket)
		}else{
			ticket.Inactivate()
		}
	}
	system.ValidTickets = tickets
	var validEntry []ParkingMapEntry
	for _,entry := range system.ParkingMap.Entry {
		if (entry.SlotNumber != slotNumber) {
			validEntry = append(validEntry, entry)
		}
	}
	system.ParkingMap.Entry = validEntry
	return fmt.Sprintf(ResponseSlotDeallocatedSuccesfully, slotNumber), nil
}


func (system *TicketingSystem) GetStatus() (outputString string, err error){
	var output []string
	for _, entry := range system.ParkingMap.Entry{
		output = append(
			output,
			fmt.Sprintf(
				"%-12v%-20v%-10v",
				entry.SlotNumber,
				entry.RegistrationNumber,
				entry.Color,
			),
		)
	}
	return strings.Join(output, "\n"),nil
}

func (system TicketingSystem) GetRegistrationNumberWithColor(color string) (registrationNumbersString string, err error) {
	var registrationNumbers []string
	for _,entry := range system.ParkingMap.Entry{
		if(entry.Color== color){
			registrationNumbers=append(registrationNumbers, entry.RegistrationNumber)
		}
	}
	registrationNumbersString = strings.Trim(strings.Replace(fmt.Sprint(registrationNumbers), " ", ", ", -1), "[]")
	return registrationNumbersString,nil
}


func (system *TicketingSystem) GetSlotsWithColor(color string) (slotNumbersString string, err error){
	var slotNumbers []int
	for _,entry := range system.ParkingMap.Entry{
		if(entry.Color == color){
			slotNumbers= append(slotNumbers, entry.SlotNumber)
		}
	}
	slotNumbersString = strings.Trim(strings.Replace(fmt.Sprint(slotNumbers), " ", ", ", -1), "[]")
	return slotNumbersString,nil
}


func (system *TicketingSystem) GetSlotNumberGivenRegistrationNumber(registrationNumber string) (slotNumber string, err error){
	var slotNumbers []int;
	for _,entry := range system.ParkingMap.Entry{
		if strings.Compare(entry.RegistrationNumber,registrationNumber)==0{
			slotNumbers = append(slotNumbers, entry.SlotNumber)
		}
	}
	switch len(slotNumbers) {
	case 0: return "",ErrorNotFound
	case 1: return strconv.Itoa(slotNumbers[0]),nil
	default: return "", ErrorSameRegistrationNumberFound
	}
}


