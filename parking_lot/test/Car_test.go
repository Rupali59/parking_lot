package test

import (
	vehicle "../internal/app/vehicle"
	"reflect"
	"testing"
)

/*
*/
func TestCarFactory_CreateVehicle(t *testing.T) {
	type args struct {
		registrationNumber string
		color              string
	}
	tests := []struct {
		name       string
		carFactory *vehicle.CarFactory
		args       args
		want       *vehicle.Vehicle
	}{
		{
			name:"Create car",
			carFactory:&vehicle.CarFactory{},
			args:args{
				registrationNumber:"KA-01-HH-3141",
				color:"White",
			},
			want:&vehicle.Vehicle{
				RegistrationNumber:"KA-01-HH-3141",
				Color:"White",
				Type:"Car",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			carFactory := &vehicle.CarFactory{}
			if got := carFactory.CreateVehicle(tt.args.registrationNumber, tt.args.color); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CarFactory.CreateVehicle() = %v, want %v", got, tt.want)
			}
		})
	}
}
