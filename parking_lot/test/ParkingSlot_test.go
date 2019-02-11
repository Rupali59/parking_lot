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
	{
	name:"Unoccupied Case",
	fields:fields{
			0,
			lot.NOT_OCCUPIED,
			nil,
	},
	want:false,
	},
		{
			name:"Occupied Case",
			fields:fields{
				0,
				lot.OCCUPIED,
				nil,
			},
			want:true,
		},
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
		{
			name :"Valid Case",
			fields:fields{
				SlotNumber:0,
				Status:lot.NOT_OCCUPIED,
				SlotVehicle:nil,
			},
			args:args{
				vehicle:vehicle.GetVehicle("Car","KA-01-HH-1234","White"),
			},
			wantStatus:true,
			wantErr:false,
		},
		{
			name :"Invalid Case",
			fields:fields{
				SlotNumber:0,
				Status:lot.OCCUPIED,
				SlotVehicle:nil,
			},
			args:args{
				vehicle:vehicle.GetVehicle("Car","KA-01-HH-1234","White"),
			},
			wantStatus:false,
			wantErr:true,
		},
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
				t.Errorf("ParkingSlot.AllotVehicle() name=%v error = %v, wantErr %v",tt.name, err, tt.wantErr)
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
	}{
		{
			name :"Valid Case",
			fields:fields{
				SlotNumber:0,
				Status:lot.OCCUPIED,
				SlotVehicle:vehicle.GetVehicle("Car","KA-01-HH-1234","White"),
			},
			wantStatus:true,
			wantErr:false,
		},
		{
			name :"Valid Case",
			fields:fields{
				SlotNumber:0,
				Status:lot.NOT_OCCUPIED,
				SlotVehicle:nil,
			},
			wantStatus:true,
			wantErr:true,
		},
	}
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
