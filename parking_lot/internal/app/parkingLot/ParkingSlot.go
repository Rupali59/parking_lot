package parkingLot

import (
	"../vehicle"
	"errors"
	_ "fmt"
)

var (
	ErrorSlotOccupied    = errors.New("Slot is already occupied")
	ErrorSlotNotOccupied = errors.New("Slot is not occupied")
)

type ParkingSlotStatus int

const (
	NOT_OCCUPIED ParkingSlotStatus = iota
	OCCUPIED
)

type ParkingSlot struct {
	SlotNumber  int
	Status      ParkingSlotStatus
	SlotVehicle *vehicle.Vehicle
}

func NewParkingSlot(slotNumber int) *ParkingSlot {
	return &ParkingSlot{
		SlotNumber:  slotNumber,
		Status:      NOT_OCCUPIED,
		SlotVehicle: nil,
	}
}

func (slot *ParkingSlot) IsOccupied() bool {
	return slot.Status == OCCUPIED
}

func (slot *ParkingSlot) AllotVehicle(vehicle *vehicle.Vehicle) (status bool, err error) {

	if slot.Status == NOT_OCCUPIED {
		slot.Status = OCCUPIED
		slot.SlotVehicle = vehicle
		return true, nil
	}
	return false, ErrorSlotOccupied
}

func (slot *ParkingSlot) Deallocate() (status bool, err error) {
	if slot.Status == NOT_OCCUPIED {
		return true, ErrorSlotNotOccupied
	}
	slot.SlotVehicle = nil
	slot.Status = NOT_OCCUPIED
	return true, nil
}
