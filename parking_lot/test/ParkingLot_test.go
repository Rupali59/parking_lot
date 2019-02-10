package test

import (
	"fmt"
	"reflect"
	"testing"
	lot "../internal/app/parkingLot"
	"../internal/app/vehicle"
)

func TestCreateParkingLot(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name           string
		args           args
		wantParkingLot *lot.ParkingLot
		wantResponse   string
		wantErr        bool
	}{
		{
			name:"Create Parking Lot Test",
			args:args{
				n:6,
			},
			wantParkingLot: &lot.ParkingLot{
				Slots:[]lot.ParkingSlot{
					{
						SlotNumber:0,
						Status:lot.NOT_OCCUPIED,
						SlotVehicle:nil,
					},
					{
						SlotNumber:1,
						Status:lot.NOT_OCCUPIED,
						SlotVehicle:nil,
					},
					{
						SlotNumber:2,
						Status:lot.NOT_OCCUPIED,
						SlotVehicle:nil,
					},
					{
						SlotNumber:3,
						Status:lot.NOT_OCCUPIED,
						SlotVehicle:nil,
					},
					{
						SlotNumber:4,
						Status:lot.NOT_OCCUPIED,
						SlotVehicle:nil,
					},
					{
						SlotNumber:5,
						Status:lot.NOT_OCCUPIED,
						SlotVehicle:nil,
					},
				},
				EmptySlots:[]int{0,1,2,3,4,5},

				Capacity:6,
			},
			wantResponse:fmt.Sprintf(lot.ResponseCreatedparkingLot,6),
			wantErr:false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotParkingLot, gotResponse, err := lot.CreateParkingLot(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateParkingLot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotParkingLot, tt.wantParkingLot) {
				t.Errorf("CreateParkingLot() gotParkingLot = %v, want %v", gotParkingLot, tt.wantParkingLot)
			}
			if gotResponse != tt.wantResponse {
				t.Errorf("CreateParkingLot() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestParkingLot_GetEmptySlot(t *testing.T) {
	type fields struct {
		Slots         []lot.ParkingSlot
		EmptySlots    []int
		Capacity      int
	}
	tests := []struct {
		name           string
		fields         fields
		wantSlotNumber int
	}{
		{
			name:"Testing Best Case",
			fields:fields{
				Slots:[]lot.ParkingSlot{
					{
						SlotNumber:0,
						Status:lot.NOT_OCCUPIED,
						SlotVehicle:nil,
					},
					{
						SlotNumber:1,
						Status:lot.NOT_OCCUPIED,
						SlotVehicle:nil,
					},
					{
						SlotNumber:2,
						Status:lot.NOT_OCCUPIED,
						SlotVehicle:nil,
					},
					{
						SlotNumber:3,
						Status:lot.NOT_OCCUPIED,
						SlotVehicle:nil,
					},
					{
						SlotNumber:4,
						Status:lot.NOT_OCCUPIED,
						SlotVehicle:nil,
					},
					{
						SlotNumber:5,
						Status:lot.NOT_OCCUPIED,
						SlotVehicle:nil,
					},
				},
				EmptySlots:[]int{0,1,2,3,4,5},
				Capacity:6,
			},
			wantSlotNumber:0,
		},
		{
			name:"Testing Best Case",
			fields:fields{
				Slots:[]lot.ParkingSlot{
					{
						SlotNumber:0,
						Status:lot.NOT_OCCUPIED,
						SlotVehicle:nil,
					},
					{
						SlotNumber:1,
						Status:lot.NOT_OCCUPIED,
						SlotVehicle:nil,
					},
					{
						SlotNumber:2,
						Status:lot.NOT_OCCUPIED,
						SlotVehicle:nil,
					},
					{
						SlotNumber:3,
						Status:lot.NOT_OCCUPIED,
						SlotVehicle:nil,
					},
					{
						SlotNumber:4,
						Status:lot.NOT_OCCUPIED,
						SlotVehicle:nil,
					},
					{
						SlotNumber:5,
						Status:lot.NOT_OCCUPIED,
						SlotVehicle:nil,
					},
				},
				EmptySlots:[]int{1,2,3,4,5},
				Capacity:6,
			},
			wantSlotNumber:1,
		},
		{
			name:"Testing Best Case",
			fields:fields{
				Slots:[]lot.ParkingSlot{
					{
						SlotNumber:0,
						Status:lot.NOT_OCCUPIED,
						SlotVehicle:nil,
					},
					{
						SlotNumber:1,
						Status:lot.NOT_OCCUPIED,
						SlotVehicle:nil,
					},
					{
						SlotNumber:2,
						Status:lot.NOT_OCCUPIED,
						SlotVehicle:nil,
					},
					{
						SlotNumber:3,
						Status:lot.NOT_OCCUPIED,
						SlotVehicle:nil,
					},
					{
						SlotNumber:4,
						Status:lot.NOT_OCCUPIED,
						SlotVehicle:nil,
					},
					{
						SlotNumber:5,
						Status:lot.NOT_OCCUPIED,
						SlotVehicle:nil,
					},
				},
				EmptySlots:[]int{2,3,4,5},
				Capacity:6,
			},
			wantSlotNumber:2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parkingLot := &lot.ParkingLot{
				Slots:         tt.fields.Slots,
				EmptySlots:    tt.fields.EmptySlots,
				Capacity:      tt.fields.Capacity,
			}
			if gotSlotNumber := parkingLot.GetEmptySlot(); gotSlotNumber != tt.wantSlotNumber {
				t.Errorf("ParkingLot.GetEmptySlot() = %v, want %v", gotSlotNumber, tt.wantSlotNumber)
			}
		})
	}
}

func TestParkingLot_Leave(t *testing.T) {
	type fields struct {
		Slots         []lot.ParkingSlot
		EmptySlots    []int
		Capacity      int
	}
	type args struct {
		slotNumber int
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
			parkingLot := &lot.ParkingLot{
				Slots:         tt.fields.Slots,
				EmptySlots:    tt.fields.EmptySlots,
				Capacity:      tt.fields.Capacity,
			}
			gotStatus, err := parkingLot.Leave(tt.args.slotNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParkingLot.Leave() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotStatus != tt.wantStatus {
				t.Errorf("ParkingLot.Leave() = %v, want %v", gotStatus, tt.wantStatus)
			}
		})
	}
}

func TestParkingLot_Park(t *testing.T) {
	type fields struct {
		Slots         []lot.ParkingSlot
		EmptySlots    []int
		Capacity      int
	}
	type args struct {
		slotNumber int
		vehicle    *vehicle.Vehicle
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantStatus bool
		wantErr    bool
	}{
		{
			name : "Park 1",
			fields:fields{
				Slots:[]lot.ParkingSlot{
					{
						SlotNumber:0,
						Status:lot.NOT_OCCUPIED,
						SlotVehicle:nil,
					},
					{
						SlotNumber:1,
						Status:lot.NOT_OCCUPIED,
						SlotVehicle:nil,
					},
					{
						SlotNumber:2,
						Status:lot.NOT_OCCUPIED,
						SlotVehicle:nil,
					},
					{
						SlotNumber:3,
						Status:lot.NOT_OCCUPIED,
						SlotVehicle:nil,
					},
					{
						SlotNumber:4,
						Status:lot.NOT_OCCUPIED,
						SlotVehicle:nil,
					},
					{
						SlotNumber:5,
						Status:lot.NOT_OCCUPIED,
						SlotVehicle:nil,
					},
				},
				EmptySlots:[]int{0,1,2,3,4,5},
				Capacity:6,
			},
			args:args{
				slotNumber:0,
				vehicle:vehicle.GetVehicle("car","KA-01-HH-1234","White"),
			},
			wantStatus:true,
			wantErr:false,
		},
		{
			name : "Park 2",
			fields:fields{
				Slots:[]lot.ParkingSlot{
					{
						SlotNumber:0,
						Status:lot.NOT_OCCUPIED,
						SlotVehicle:nil,
					},
					{
						SlotNumber:1,
						Status:lot.NOT_OCCUPIED,
						SlotVehicle:nil,
					},
					{
						SlotNumber:2,
						Status:lot.NOT_OCCUPIED,
						SlotVehicle:nil,
					},
					{
						SlotNumber:3,
						Status:lot.NOT_OCCUPIED,
						SlotVehicle:nil,
					},
					{
						SlotNumber:4,
						Status:lot.NOT_OCCUPIED,
						SlotVehicle:nil,
					},
					{
						SlotNumber:5,
						Status:lot.NOT_OCCUPIED,
						SlotVehicle:nil,
					},
				},
				EmptySlots:[]int{1,2,3,4,5},
				Capacity:6,
			},
			args:args{
				slotNumber:1,
				vehicle:vehicle.GetVehicle("car","KA-01-HH-9999","White"),
			},
			wantStatus:true,
			wantErr:false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parkingLot := &lot.ParkingLot{
				Slots:         tt.fields.Slots,
				EmptySlots:    tt.fields.EmptySlots,
				Capacity:      tt.fields.Capacity,
			}
			gotStatus, err := parkingLot.Park(tt.args.slotNumber, tt.args.vehicle)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParkingLot.Park() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotStatus != tt.wantStatus {
				t.Errorf("ParkingLot.Park() = %v, want %v", gotStatus, tt.wantStatus)
			}
		})
	}
}
