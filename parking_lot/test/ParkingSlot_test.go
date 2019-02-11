package test

import (
	lot "../internal/app/parkingLot"
	"../internal/app/vehicle"
	_ "fmt"
	"reflect"
	"testing"
)

func TestNewParkingSlot(t *testing.T) {
	type args struct {
		slotNumber int
	}
	tests := []struct {
		name string
		args args
		want *lot.ParkingSlot
	}{
		{
			"Create Test Parking Slot",
			args{0},
			&lot.ParkingSlot{
				0,
				lot.NOT_OCCUPIED,
				nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lot.NewParkingSlot(tt.args.slotNumber); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewParkingSlot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParkingSlot_IsOccupied(t *testing.T) {
	type fields struct {
		SlotNumber  int
		Status      lot.ParkingSlotStatus
		SlotVehicle *vehicle.Vehicle
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			slot := &lot.ParkingSlot{
				SlotNumber:  tt.fields.SlotNumber,
				Status:      tt.fields.Status,
				SlotVehicle: tt.fields.SlotVehicle,
			}
			if got := slot.IsOccupied(); got != tt.want {
				t.Errorf("ParkingSlot.IsOccupied() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParkingSlot_AllotVehicle(t *testing.T) {
	type fields struct {
		SlotNumber  int
		Status      lot.ParkingSlotStatus
		SlotVehicle *vehicle.Vehicle
	}
	type args struct {
		vehicle *vehicle.Vehicle
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantStatus bool
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			slot := &lot.ParkingSlot{
				SlotNumber:  tt.fields.SlotNumber,
				Status:      tt.fields.Status,
				SlotVehicle: tt.fields.SlotVehicle,
			}
			gotStatus, err := slot.AllotVehicle(tt.args.vehicle)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParkingSlot.AllotVehicle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotStatus != tt.wantStatus {
				t.Errorf("ParkingSlot.AllotVehicle() = %v, want %v", gotStatus, tt.wantStatus)
			}
		})
	}
}

func TestParkingSlot_Deallocate(t *testing.T) {
	type fields struct {
		SlotNumber  int
		Status      lot.ParkingSlotStatus
		SlotVehicle *vehicle.Vehicle
	}
	tests := []struct {
		name       string
		fields     fields
		wantStatus bool
		wantErr    bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			slot := &lot.ParkingSlot{
				SlotNumber:  tt.fields.SlotNumber,
				Status:      tt.fields.Status,
				SlotVehicle: tt.fields.SlotVehicle,
			}
			gotStatus, err := slot.Deallocate()
			if (err != nil) != tt.wantErr {
				t.Errorf("ParkingSlot.Deallocate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotStatus != tt.wantStatus {
				t.Errorf("ParkingSlot.Deallocate() = %v, want %v", gotStatus, tt.wantStatus)
			}
		})
	}
}
