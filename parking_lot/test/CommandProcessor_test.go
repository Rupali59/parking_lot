package test

import (
	"testing"
	processor "../internal/app/processor"
	ticketingSystem "../internal/app/ticketing_system"
)

func TestCommandProcessor_Process(t *testing.T) {
	type fields struct {
		TicketingSystem *ticketingSystem.TicketingSystem
	}
	type args struct {
		commandString string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantOutput string
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			processor := &processor.CommandProcessor{
				TicketingSystem: tt.fields.TicketingSystem,
			}
			gotOutput, err := processor.Process(tt.args.commandString)
			if (err != nil) != tt.wantErr {
				t.Errorf("CommandProcessor.Process() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOutput != tt.wantOutput {
				t.Errorf("CommandProcessor.Process() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func TestCommandProcessor_CreateParkingLot(t *testing.T) {
	type fields struct {
		TicketingSystem *ticketingSystem.TicketingSystem
	}
	type args struct {
		args []string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantOutput string
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			processor := &processor.CommandProcessor{
				TicketingSystem: tt.fields.TicketingSystem,
			}
			gotOutput, err := processor.CreateParkingLot(tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("CommandProcessor.CreateParkingLot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOutput != tt.wantOutput {
				t.Errorf("CommandProcessor.CreateParkingLot() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func TestCommandProcessor_ParkCommand(t *testing.T) {
	type fields struct {
		TicketingSystem *ticketingSystem.TicketingSystem
	}
	type args struct {
		args []string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantOutput string
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			processor := &processor.CommandProcessor{
				TicketingSystem: tt.fields.TicketingSystem,
			}
			gotOutput, err := processor.ParkCommand(tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("CommandProcessor.ParkCommand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOutput != tt.wantOutput {
				t.Errorf("CommandProcessor.ParkCommand() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func TestCommandProcessor_LeaveCommand(t *testing.T) {
	type fields struct {
		TicketingSystem *ticketingSystem.TicketingSystem
	}
	type args struct {
		args []string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantOutput string
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			processor := &processor.CommandProcessor{
				TicketingSystem: tt.fields.TicketingSystem,
			}
			gotOutput, err := processor.LeaveCommand(tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("CommandProcessor.LeaveCommand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOutput != tt.wantOutput {
				t.Errorf("CommandProcessor.LeaveCommand() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func TestCommandProcessor_GetStatusCommand(t *testing.T) {
	type fields struct {
		TicketingSystem *ticketingSystem.TicketingSystem
	}
	tests := []struct {
		name       string
		fields     fields
		wantOutput string
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			processor := &processor.CommandProcessor{
				TicketingSystem: tt.fields.TicketingSystem,
			}
			gotOutput, err := processor.GetStatusCommand()
			if (err != nil) != tt.wantErr {
				t.Errorf("CommandProcessor.GetStatusCommand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOutput != tt.wantOutput {
				t.Errorf("CommandProcessor.GetStatusCommand() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func TestCommandProcessor_GetRegistrationNumberGivenColorCommand(t *testing.T) {
	type fields struct {
		TicketingSystem *ticketingSystem.TicketingSystem
	}
	type args struct {
		args []string
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
			processor := &processor.CommandProcessor{
				TicketingSystem: tt.fields.TicketingSystem,
			}
			gotSlotNumber, err := processor.GetRegistrationNumberGivenColorCommand(tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("CommandProcessor.GetRegistrationNumberGivenColorCommand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSlotNumber != tt.wantSlotNumber {
				t.Errorf("CommandProcessor.GetRegistrationNumberGivenColorCommand() = %v, want %v", gotSlotNumber, tt.wantSlotNumber)
			}
		})
	}
}

func TestCommandProcessor_GetSlotsWithRegistrationNumberCommand(t *testing.T) {
	type fields struct {
		TicketingSystem *ticketingSystem.TicketingSystem
	}
	type args struct {
		args []string
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
			processor := &processor.CommandProcessor{
				TicketingSystem: tt.fields.TicketingSystem,
			}
			gotSlotNumber, err := processor.GetSlotsWithRegistrationNumberCommand(tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("CommandProcessor.GetSlotsWithRegistrationNumberCommand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSlotNumber != tt.wantSlotNumber {
				t.Errorf("CommandProcessor.GetSlotsWithRegistrationNumberCommand() = %v, want %v", gotSlotNumber, tt.wantSlotNumber)
			}
		})
	}
}

func TestCommandProcessor_GetSlotNumbersGivenColor(t *testing.T) {
	type fields struct {
		TicketingSystem *ticketingSystem.TicketingSystem
	}
	type args struct {
		args []string
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
			processor := &processor.CommandProcessor{
				TicketingSystem: tt.fields.TicketingSystem,
			}
			gotSlotNumber, err := processor.GetSlotNumbersGivenColor(tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("CommandProcessor.GetSlotNumbersGivenColor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSlotNumber != tt.wantSlotNumber {
				t.Errorf("CommandProcessor.GetSlotNumbersGivenColor() = %v, want %v", gotSlotNumber, tt.wantSlotNumber)
			}
		})
	}
}
