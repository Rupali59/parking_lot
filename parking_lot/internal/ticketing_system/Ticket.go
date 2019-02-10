package ticketing_system

import "time"

var TicketNumber = 0

type Ticket struct {
	TicketNumber int
	SlotNumber int
	VehicleNumber string
	Color string
	CreatedOn time.Time
}

func NewTicket(slotNumber int, vehicleNumber string, color string) *Ticket {
	return &Ticket{
		TicketNumber: TicketNumber,
		SlotNumber: slotNumber,
		VehicleNumber: vehicleNumber,
		CreatedOn:time.Now(),
	}
	TicketNumber = TicketNumber+1;
}


func(ticket *Ticket) delete(){
	ticket.delete()
}