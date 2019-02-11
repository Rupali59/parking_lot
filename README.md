### Parking Lot Problem Statement
I own a parking lot that can hold up to 'n' cars at any given point in time. Each slot is
given a number starting at 1 increasing with increasing distance from the entry point
in steps of one. I want to create an automated ticketing system that allows my
customers to use my parking lot without human intervention.

When a car enters my parking lot, I want to have a ticket issued to the driver. The
ticket issuing process includes us documenting the registration number (number
plate) and the colour of the car and allocating an available parking slot to the car
before actually handing over a ticket to the driver (we assume that our customers are
nice enough to always park in the slots allocated to them). The customer should be
allocated a parking slot which is nearest to the entry. At the exit the customer returns
the ticket which then marks the slot they were using as being available.

Due to government regulation, the system should provide me with the ability to find out:
1. Registration numbers of all cars of a particular colour.
2. Slot number in which a car with a given registration number is parked.
3. Slot numbers of all slots where a car of a particular colour is parked.                                                                                    

We interact with the system via a simple set of commands which produce a specific
output. Please take a look at the example below, which includes all the commands
you need to support - they're self explanatory. The system should allow input in two
ways. Just to clarify, the same codebase should support both modes of input - we
don't want two distinct submissions.
1. It should provide us with an interactive command prompt based shell where
commands can be typed in
2. It should accept a filename as a parameter at the command prompt and read the
commands from that file


#### Application build requirements
    go1.11.5 is required. Executable permission enabled for `parking_lot`. 
 Please make sure its in executable mode in your system. Run 
    `chmod 755 parking_lot`
    in command prompt if not.
 Please make sure your  ***GOPATH*** is set to application directory or set as follows.
 
   `cd [solution directory]`
   `export GOPATH=$(pwd)`
   
#### Application Components    
Application includes following components
1. Helpers
2. Processor
3. Parking Lot 
4. Ticketing System
5. Vehicle


##### Helpers
Provide helpers to read commands from the file system or the command line.


##### Processor
The file checks the content of the command and processes it accordingly. Results in an error if the arguments are invalid.
        
##### Vehicle
The basic required object in the system. Object have properties
    1. Registration number {uniq identifier}
    2. Colour

##### Parking Lot
A parking lot is composed of multiple slots. Slots have a sequential ordering with respect to their distance from the parking lot entrance.
    A `Parking Lot` consists of:
        1. Slots in the lot
        2. A list of empty slots
        3. Capacity of the slot
    It can do the following:-
        1. Check for the closest empty slot
        2. Park Car
        3. Leave Car
        
`Parking Slot ` consists of
        1. Slot Number
        2. Vehicle Object
    It can do the following:-
        1. Check if it occupied
        2. Allocate a vehicle
        3. Deallocate the vehicle
 
##### Ticketing System
The ticketing system consists of 3 components:-
1. Valid Tickets
2. Parking Map
3. Parking Lot object

It has the following functionality
    1. Generate Ticket
    2. Invalidate Ticket
    3. Gets the current status of the parking lot 
    4. Get Registration Number of all parked vehicles given color
    5. Gets the slot numbers of all parked vehicles with the given color
    6. Gets the slot number where the car having registration number is parked