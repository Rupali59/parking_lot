package parkingLot

import (
	"errors"
	"fmt"
	"../vehicle"
)

var (
	ErrorCreation             = errors.New("Could not create a parking lot.")
	ResponseSlotNotFound = "No slot found with number %d found."
	ResponseCreatedparkingLot = "Created a parking lot with %d slots."
)

type ParkingLot struct {
	Slots    []ParkingSlot;
	EmptySlots []int;
	Capacity int;
}

func CreateParkingLot(n int) (parkingLot *ParkingLot, response string, err error) {
	var slots []ParkingSlot;
	var slotNumbers []int;
	for i := 0; i < n; i++ {
		slot := NewParkingSlot(i)
		slots = append(slots, *slot);
		slotNumbers = append(slotNumbers,i)
	}
	if (len(slots) == n) {
		parkingLot := ParkingLot{
			Slots:    slots,
			Capacity: n,
			EmptySlots:slotNumbers,
		};
		return &parkingLot, fmt.Sprintf(ResponseCreatedparkingLot,n), nil
	}else{
		return nil, "",ErrorCreation
	}
}


func(parkingLot *ParkingLot) GetEmptySlot()(slotNumber int){
	if(len(parkingLot.EmptySlots)>0){
		min := parkingLot.Capacity+1
		for i:=0;i<len(parkingLot.EmptySlots);i++{
			if(parkingLot.EmptySlots[i]<min){
				min = parkingLot.EmptySlots[i]
			}
		}
		return min;
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
				parkingLot.Slots[slotNumber]=slot;
				parkingLot.EmptySlots = append(parkingLot.EmptySlots, slot.SlotNumber)
				return true,nil
			}
		}
	}
	return false,errors.New(fmt.Sprintf(ResponseSlotNotFound,slotNumber))
}

func(parkingLot *ParkingLot)Park(slotNumber int,vehicle *vehicle.Vehicle) (status bool, err error){
	slot := parkingLot.Slots[slotNumber]
	allotStatus,err := slot.AllotVehicle(vehicle)
	parkingLot.Slots[slotNumber] = slot

	if(err!=nil){
		return false,err
	}
	if(allotStatus == true){
		otherSlots := []int{};
		for _,i:=range parkingLot.EmptySlots{
			if(i != slotNumber){
				otherSlots = append(otherSlots, i)
			}
		}
		parkingLot.EmptySlots=otherSlots
		return true,nil
	}
	return false,err
}