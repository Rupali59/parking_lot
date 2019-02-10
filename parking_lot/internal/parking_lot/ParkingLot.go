package parking_lot

import (
	"errors"
	"fmt"
	"../vehicle"
)

var (
	ErrorCreation             = errors.New("Could not create a parking lot.")
	ErrorSlotNotFound = errors.New("No slot found with number %d found.")
	ResponseCreatedparkingLot = "Created a parking lot with %d slots."
)

type ParkingLot struct {
	Slots    []ParkingSlot;
	EmptySlots []int;
	OccupiedSlots []int;
	Capacity int;
}

func CreateParkingLot(n int) (parkingLot *ParkingLot, err error, response string) {
	var slots []ParkingSlot;
	var slotNumbers []int;
	for i := 0; i <= n; i++ {
		slot := NewParkingSlot()
		slots = append(slots, *slot);
		slotNumbers = append(slotNumbers,i)
	}
	if (len(slots) == n) {
		parkingLot := ParkingLot{
			Slots:    slots,
			Capacity: n,
			EmptySlots:slotNumbers,
		};
		return &parkingLot, nil, fmt.Sprintf(ResponseCreatedparkingLot,n)
	}
	return nil, ErrorCreation, ""
}


func(parkingLot *ParkingLot) GetEmptySlot()(slotNumber int){
	if(len(parkingLot.EmptySlots)>0){
		min := parkingLot.Capacity+1
		for _,i:=range parkingLot.EmptySlots{
			if(min < i){
				min = i
			}
		}
		return parkingLot.EmptySlots[min];
	}
	return -1
}


func(parkingLot *ParkingLot) Leave(slotNumber int) (status bool, err error){
	for _,slot:= range parkingLot.Slots{
		if(slot.SlotNumber == slotNumber){
			status,err := slot.Deallocate()
			if(err !=nil){
				return false, err
			}
			if(status){
				parkingLot.EmptySlots = append(parkingLot.EmptySlots, slot.SlotNumber)
				return true,nil
			}
		}
	}
	return false,ErrorSlotNotFound
}

func(parkingLot *ParkingLot)Park(slotNumber int,vehicle *vehicle.Vehicle) (status bool, err error){
	slot := parkingLot.Slots[slotNumber]
	allotStatus,err := slot.AllotVehicle(vehicle)
	if(err!=nil){
		return false,err
	}
	if(allotStatus == true){
		var empty []int
		for i := range parkingLot.EmptySlots{
			if(slotNumber != i){
				empty= append(empty, i)
			}
		}
		parkingLot.EmptySlots = empty
		return true,nil
	}
}