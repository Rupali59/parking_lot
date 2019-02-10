package test

import (
	"reflect"
	"testing"
	"time"
	system "../internal/app/ticketing_system"
)

func TestNewTicket(t *testing.T) {
	type args struct {
		slotNumber    int
		vehicleNumber string
		color         string
	}
	tests := []struct {
		name string
		args args
		want *system.Ticket
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := system.NewTicket(tt.args.slotNumber, tt.args.vehicleNumber, tt.args.color); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTicket() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTicket_delete(t *testing.T) {
	type fields struct {
		TicketNumber  int
		SlotNumber    int
		VehicleNumber string
		Color         string
		CreatedOn     time.Time
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ticket := &system.Ticket{
				TicketNumber:  tt.fields.TicketNumber,
				SlotNumber:    tt.fields.SlotNumber,
				VehicleNumber: tt.fields.VehicleNumber,
				Color:         tt.fields.Color,
				CreatedOn:     tt.fields.CreatedOn,
			}
			ticket.Inactivate()
		})
	}
}
