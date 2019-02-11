package ticketing_system

import "time"

var TicketNumber = 0

type Ticket struct {
	TicketNumber  int
	SlotNumber    int
	VehicleNumber string
	Color         string
	CreatedOn     time.Time
}

func NewTicket(slotNumber int, vehicleNumber string, color string) *Ticket {
	TicketNumber = TicketNumber + 1
	return &Ticket{
		TicketNumber:  TicketNumber,
		SlotNumber:    slotNumber,
		VehicleNumber: vehicleNumber,
		CreatedOn:     time.Now(),
	}
}

func (ticket *Ticket) Inactivate() {
	ticket = nil
}
