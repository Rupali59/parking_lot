package processor

import (
	parkingLot "../parking_lot"
	"errors"
	"strconv"
	"strings"
	vehicle "../vehicle"
	ticketingSystem "../ticketing_system"
)

type CommandProcessor struct {
	TicketingSystem *ticketingSystem.TicketingSystem
}

var (
	ErrorInvalidCommand = errors.New("Invalid Command")
	ErrorInvalidArguments = errors.New("Invalid Number of arguments")
)

func (processor *CommandProcessor) Process(commandString string) (output string, err error) {
	args := strings.Split(commandString, " ")

	command := args[0]
	switch command {
	case "create_parking_lot":
		return processor.CreateParkingLot(args)
	case "park":

	case "leave":
	case "status":
		return processor.GetStatusCommand()
	case "registration_numbers_for_cars_with_colour":
		return  processor.GetRegistrationNumberGivenColorCommand(args)
	case "slot_numbers_for_cars_with_colour":
		return processor.GetSlotNumbersGivenColor(args)
	case "slot_number_for_registration_number":
		return processor.GetSlotsWithRegistrationNumberCommand(args)
	default:
		return "", ErrorInvalidCommand
	}
	return "", ErrorInvalidCommand
}


func (processor *CommandProcessor) CreateParkingLot(args []string) (output string, err error) {
	if(len(args) ==2){
		n, err := strconv.ParseInt(args[1],10,64)
		if err!=nil{
			return "", err
		}
		parkingLot,err,response := parkingLot.CreateParkingLot(n)
		if err!=nil{
			return "", err
		}
		processor.TicketingSystem = ticketingSystem.NewTicketingSystem(parkingLot)
		return response,err
	}
	return "",ErrorInvalidArguments
}

func (processor *CommandProcessor) ParkCommand(args []string) (output string, err error) {
	if(len(args) ==3){
		registrationNumber := args[1]
		color := args[2]
		car := vehicle.GetVehicle("car",registrationNumber,color)
		return processor.TicketingSystem.GenerateTicket(car)
	}
	return "",ErrorInvalidArguments
}

func (processor *CommandProcessor) LeaveCommand(args []string) (output string, err error) {
	if(len(args) ==2){
		slotNumber,err :=strconv.ParseInt(args[1],10,16)
		if(err != nil){
			return "", err
		}
		return processor.TicketingSystem.InvalidateTicket(int(slotNumber))
	}
	return "",ErrorInvalidArguments
}

func (processor *CommandProcessor) GetStatusCommand() (output string, err error) {
	return processor.TicketingSystem.GetStatus()
}

func (processor *CommandProcessor) GetRegistrationNumberGivenColorCommand(args []string) (slotNumber string, err error) {
	if (len(args) == 2) {
		color := args[1]
		return processor.TicketingSystem.GetRegistrationNumberWithColor(color)
	}
	return "", ErrorInvalidArguments
}

func (processor *CommandProcessor) GetSlotsWithRegistrationNumberCommand(args []string) (slotNumber string, err error) {
	if (len(args) == 2) {
		registrationNumber := args[1]
		return processor.TicketingSystem.GetSlotNumberGivenRegistrationNumber(registrationNumber)
	}
	return "", ErrorInvalidArguments
}


func (processor *CommandProcessor) GetSlotNumbersGivenColor(args []string) (slotNumber string, err error) {
	if (len(args) == 2) {
		color := args[1]
		return processor.TicketingSystem.GetSlotNumberGivenRegistrationNumber(color)
	}
	return "", ErrorInvalidArguments
}