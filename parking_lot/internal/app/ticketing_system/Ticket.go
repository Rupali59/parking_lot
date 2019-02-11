package ticketing_system

import "time"

var TicketNumber = 0

//A physical entity provided to the vehicle owner
type Ticket struct {
	TicketNumber  int //TicketNumber is the ticket identifier
	SlotNumber    int //SlotNumber at which the vehicle was parked
	VehicleNumber string //Registration number of the vehicle that was parked
	Color         string //Color of the vehicle that was parked
	CreatedOn     time.Time //Time at which the ticket was created
}

//Creates a new ticket for the slot number and vehicle
func NewTicket(slotNumber int, vehicleNumber string, color string) *Ticket {
	TicketNumber = TicketNumber + 1
	return &Ticket{
		TicketNumber:  TicketNumber,
		SlotNumber:    slotNumber,
		VehicleNumber: vehicleNumber,
		CreatedOn:     time.Now(),
	}
}

//Deletes the ticket
func (ticket *Ticket) Inactivate() {
	ticket = nil
}
