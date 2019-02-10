package test

import (
	"../internal/app/vehicle"
	"reflect"
	"testing"
)

func TestGetVehicle(t *testing.T) {
	type args struct {
		vehicleType        string
		registrationNumber string
		color              string
	}
	tests := []struct {
		name string
		args args
		want *vehicle.Vehicle
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := vehicle.GetVehicle(tt.args.vehicleType, tt.args.registrationNumber, tt.args.color); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetVehicle() = %v, want %v", got, tt.want)
			}
		})
	}
}
