//Package parkingLot stores the models for the lot and slot
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
	SlotNumber  int //SlotNumber represents the distance from entry of the slot
	Status      ParkingSlotStatus//Status of the slot
	SlotVehicle *vehicle.Vehicle//Vehicle parked in the slot
}

//NewParkingSlot creates a new parking slot given its distance from the entry
func NewParkingSlot(slotNumber int) *ParkingSlot {
	return &ParkingSlot{
		SlotNumber:  slotNumber,
		Status:      NOT_OCCUPIED,
		SlotVehicle: nil,
	}
}

//IsOccupied checks whether the slot has been occupied
func (slot *ParkingSlot) IsOccupied() bool {
	return slot.Status == OCCUPIED
}


//AllotVehicle allots the given vehicle to the slot
func (slot *ParkingSlot) AllotVehicle(vehicle *vehicle.Vehicle) (status bool, err error) {

	if slot.Status == NOT_OCCUPIED {
		slot.Status = OCCUPIED
		slot.SlotVehicle = vehicle
		return true, nil
	}
	return false, ErrorSlotOccupied
}

//Deallocate removes the vehicle from the slot
func (slot *ParkingSlot) Deallocate() (status bool, err error) {
	if slot.Status == NOT_OCCUPIED {
		return true, ErrorSlotNotOccupied
	}
	slot.SlotVehicle = nil
	slot.Status = NOT_OCCUPIED
	return true, nil
}
