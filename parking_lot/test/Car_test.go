package test

import (
	"reflect"
	"testing"
	vehicle "../internal/app/vehicle"
)

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
		// TODO: Add test cases.
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
