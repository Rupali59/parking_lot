package test

import (
	parkinglot "../internal/app/parkingLot"
	system "../internal/app/ticketing_system"
	vehicle "../internal/app/vehicle"
	"fmt"
	"reflect"
	"testing"
)

func TestNewTicketingSystem(t *testing.T) {
	type args struct {
		lot *parkinglot.ParkingLot
	}
	tests := []struct {
		name string
		args args
		want *system.TicketingSystem
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := system.NewTicketingSystem(tt.args.lot); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("system.NewTicketingSystem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTicketingSystem_GenerateTicket(t *testing.T) {
	type fields struct {
		ValidTickets []system.Ticket
		Parkinglot   *parkinglot.ParkingLot
		ParkingMap   *system.ParkingMap
	}
	type args struct {
		car *vehicle.Vehicle
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantResponse string
		wantErr      bool
	}{
		{
			name: "Best case Generate Ticket",
			fields: fields{
				[]system.Ticket{},
				&parkinglot.ParkingLot{
					Slots: []parkinglot.ParkingSlot{
						{
							SlotNumber:  0,
							Status:      parkinglot.NOT_OCCUPIED,
							SlotVehicle: nil,
						},
						{
							SlotNumber:  1,
							Status:      parkinglot.NOT_OCCUPIED,
							SlotVehicle: nil,
						},
						{
							SlotNumber:  2,
							Status:      parkinglot.NOT_OCCUPIED,
							SlotVehicle: nil,
						},
						{
							SlotNumber:  3,
							Status:      parkinglot.NOT_OCCUPIED,
							SlotVehicle: nil,
						},
						{
							SlotNumber:  4,
							Status:      parkinglot.NOT_OCCUPIED,
							SlotVehicle: nil,
						},
						{
							SlotNumber:  5,
							Status:      parkinglot.NOT_OCCUPIED,
							SlotVehicle: nil,
						},
					},
					EmptySlots: []int{0, 1, 2, 3, 4, 5},
					Capacity:   6,
				},
				system.NewParkingMap(),
			},
			args: args{
				car: &vehicle.Vehicle{
					Type:               "car",
					RegistrationNumber: "KA-01-HH-1234",
					Color:              "White",
				},
			},
			wantResponse: fmt.Sprintf(system.ResponseSlotAllocatedSuccessfully, 1),
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			system := &system.TicketingSystem{
				ValidTickets: tt.fields.ValidTickets,
				Parkinglot:   tt.fields.Parkinglot,
				ParkingMap:   tt.fields.ParkingMap,
			}
			gotResponse, err := system.GenerateTicket(tt.args.car)
			if (err != nil) != tt.wantErr {
				t.Errorf("TicketingSystem.GenerateTicket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResponse != tt.wantResponse {
				t.Errorf("TicketingSystem.GenerateTicket() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestTicketingSystem_InvalidateTicket(t *testing.T) {
	type fields struct {
		ValidTickets []system.Ticket
		Parkinglot   *parkinglot.ParkingLot
		ParkingMap   *system.ParkingMap
	}
	type args struct {
		slotNumber int
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantResponse string
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			system := &system.TicketingSystem{
				ValidTickets: tt.fields.ValidTickets,
				Parkinglot:   tt.fields.Parkinglot,
				ParkingMap:   tt.fields.ParkingMap,
			}
			gotResponse, err := system.InvalidateTicket(tt.args.slotNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("TicketingSystem.InvalidateTicket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResponse != tt.wantResponse {
				t.Errorf("TicketingSystem.InvalidateTicket() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestTicketingSystem_GetStatus(t *testing.T) {
	type fields struct {
		ValidTickets []system.Ticket
		Parkinglot   *parkinglot.ParkingLot
		ParkingMap   *system.ParkingMap
	}
	tests := []struct {
		name             string
		fields           fields
		wantOutputString string
		wantErr          bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			system := &system.TicketingSystem{
				ValidTickets: tt.fields.ValidTickets,
				Parkinglot:   tt.fields.Parkinglot,
				ParkingMap:   tt.fields.ParkingMap,
			}
			gotOutputString, err := system.GetStatus()
			if (err != nil) != tt.wantErr {
				t.Errorf("TicketingSystem.GetStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOutputString != tt.wantOutputString {
				t.Errorf("TicketingSystem.GetStatus() = %v, want %v", gotOutputString, tt.wantOutputString)
			}
		})
	}
}

func TestTicketingSystem_GetRegistrationNumberWithColor(t *testing.T) {
	type fields struct {
		ValidTickets []system.Ticket
		Parkinglot   *parkinglot.ParkingLot
		ParkingMap   *system.ParkingMap
	}
	type args struct {
		color string
	}
	tests := []struct {
		name                          string
		fields                        fields
		args                          args
		wantRegistrationNumbersString string
		wantErr                       bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			system := system.TicketingSystem{
				ValidTickets: tt.fields.ValidTickets,
				Parkinglot:   tt.fields.Parkinglot,
				ParkingMap:   tt.fields.ParkingMap,
			}
			gotRegistrationNumbersString, err := system.GetRegistrationNumberWithColor(tt.args.color)
			if (err != nil) != tt.wantErr {
				t.Errorf("TicketingSystem.GetRegistrationNumberWithColor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRegistrationNumbersString != tt.wantRegistrationNumbersString {
				t.Errorf("TicketingSystem.GetRegistrationNumberWithColor() = %v, want %v", gotRegistrationNumbersString, tt.wantRegistrationNumbersString)
			}
		})
	}
}

func TestTicketingSystem_GetSlotsWithColor(t *testing.T) {
	type fields struct {
		ValidTickets []system.Ticket
		Parkinglot   *parkinglot.ParkingLot
		ParkingMap   *system.ParkingMap
	}
	type args struct {
		color string
	}
	tests := []struct {
		name                  string
		fields                fields
		args                  args
		wantSlotNumbersString string
		wantErr               bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			system := &system.TicketingSystem{
				ValidTickets: tt.fields.ValidTickets,
				Parkinglot:   tt.fields.Parkinglot,
				ParkingMap:   tt.fields.ParkingMap,
			}
			gotSlotNumbersString, err := system.GetSlotsWithColor(tt.args.color)
			if (err != nil) != tt.wantErr {
				t.Errorf("TicketingSystem.GetSlotsWithColor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSlotNumbersString != tt.wantSlotNumbersString {
				t.Errorf("TicketingSystem.GetSlotsWithColor() = %v, want %v", gotSlotNumbersString, tt.wantSlotNumbersString)
			}
		})
	}
}

func TestTicketingSystem_GetSlotNumberGivenRegistrationNumber(t *testing.T) {
	type fields struct {
		ValidTickets []system.Ticket
		Parkinglot   *parkinglot.ParkingLot
		ParkingMap   *system.ParkingMap
	}
	type args struct {
		registrationNumber string
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantSlotNumber string
		wantErr        bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			system := &system.TicketingSystem{
				ValidTickets: tt.fields.ValidTickets,
				Parkinglot:   tt.fields.Parkinglot,
				ParkingMap:   tt.fields.ParkingMap,
			}
			gotSlotNumber, err := system.GetSlotNumberGivenRegistrationNumber(tt.args.registrationNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("TicketingSystem.GetSlotNumberGivenRegistrationNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSlotNumber != tt.wantSlotNumber {
				t.Errorf("TicketingSystem.GetSlotNumberGivenRegistrationNumber() = %v, want %v", gotSlotNumber, tt.wantSlotNumber)
			}
		})
	}
}
