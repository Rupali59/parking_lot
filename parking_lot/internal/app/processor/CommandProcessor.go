//package processor has the processor which handles the commands
package processor

import (
	"../parkingLot"
	ticketingSystem "../ticketing_system"
	"../vehicle"
	"errors"
	"strconv"
	"strings"
)

type CommandProcessor struct {
	TicketingSystem *ticketingSystem.TicketingSystem //Has an instance of the ticketing system
}

var (
	ErrorInvalidCommand       = errors.New("Invalid Command")
	ErrorInvalidArguments     = errors.New("Invalid Number of arguments")
	ErrorParkingLotNotCreated = errors.New("Parking Lot Not Created")
)


//Process processes the command passed to it
func (processor *CommandProcessor) Process(commandString string) (output string, err error) {
	args := strings.Split(commandString, " ")

	command := args[0]
	switch command {
	case "create_parking_lot":
		return processor.CreateParkingLotCommand(args)
	case "park":
		return processor.ParkCommand(args)
	case "leave":
		return processor.LeaveCommand(args)
	case "status":
		return processor.GetStatusCommand()
	case "registration_numbers_for_cars_with_colour":
		return processor.GetRegistrationNumberGivenColorCommand(args)
	case "slot_numbers_for_cars_with_colour":
		return processor.GetSlotNumbersGivenColorCommand(args)
	case "slot_number_for_registration_number":
		return processor.GetSlotsWithRegistrationNumberCommand(args)
	default:
		return "", ErrorInvalidCommand
	}
	return "", ErrorInvalidCommand
}

//CreateParkingLotCommand processes the create parking command
func (processor *CommandProcessor) CreateParkingLotCommand(args []string) (output string, err error) {
	if len(args) == 2 {
		n, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			return "", err
		}
		parkingLot, response, err := parkingLot.CreateParkingLot(int(n))
		if err != nil {
			return "", err
		}
		processor.TicketingSystem = ticketingSystem.NewTicketingSystem(parkingLot)
		return response, err
	}
	return "", ErrorInvalidArguments
}

//ParkCommand creates the vehicle that is supposed to be parked
func (processor *CommandProcessor) ParkCommand(args []string) (output string, err error) {
	if processor.TicketingSystem == nil {
		return "", ErrorParkingLotNotCreated
	}
	if len(args) == 3 {
		registrationNumber := args[1]
		color := args[2]
		car := vehicle.GetVehicle("car", registrationNumber, color)
		return processor.TicketingSystem.GenerateTicket(car)
	}
	return "", ErrorInvalidArguments
}

//LeaveCommand empties out the slot given the slot number
func (processor *CommandProcessor) LeaveCommand(args []string) (output string, err error) {
	if processor.TicketingSystem == nil {
		return "", ErrorParkingLotNotCreated
	}
	if len(args) == 2 {
		slotNumber, err := strconv.ParseInt(args[1], 10, 16)
		if err != nil {
			return "", err
		}
		return processor.TicketingSystem.InvalidateTicket(int(slotNumber))
	}
	return "", ErrorInvalidArguments
}

//GetStatusCommand gets the current status of the parking lot
func (processor *CommandProcessor) GetStatusCommand() (output string, err error) {
	if processor.TicketingSystem == nil {
		return "", ErrorParkingLotNotCreated
	}
	return processor.TicketingSystem.GetStatus()
}

//GetRegistrationNumberGivenColorCommand gets the registration number of all the cars currently parked in the lot with the given color
func (processor *CommandProcessor) GetRegistrationNumberGivenColorCommand(args []string) (slotNumber string, err error) {
	if processor.TicketingSystem == nil {
		return "", ErrorParkingLotNotCreated
	}
	if len(args) == 2 {
		color := args[1]
		return processor.TicketingSystem.GetRegistrationNumberWithColor(color)
	}
	return "", ErrorInvalidArguments
}

//GetSlotsWithRegistrationNumberCommand gets the slot number where the car with the given registration number is parked
func (processor *CommandProcessor) GetSlotsWithRegistrationNumberCommand(args []string) (slotNumber string, err error) {
	if processor.TicketingSystem == nil {
		return "", ErrorParkingLotNotCreated
	}
	if len(args) == 2 {
		registrationNumber := args[1]
		return processor.TicketingSystem.GetSlotNumberGivenRegistrationNumber(registrationNumber)
	}
	return "", ErrorInvalidArguments
}

//GetSlotNumbersGivenColorCommand gets all the slots where the vehicle having the color is parked
func (processor *CommandProcessor) GetSlotNumbersGivenColorCommand(args []string) (slotNumber string, err error) {
	if processor.TicketingSystem == nil {
		return "", ErrorParkingLotNotCreated
	}
	if len(args) == 2 {
		color := args[1]
		return processor.TicketingSystem.GetSlotsWithColor(color)
	}
	return "", ErrorInvalidArguments
}
