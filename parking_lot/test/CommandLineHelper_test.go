package test

import (
	"testing"
	helpers "../internal/app/helpers"
)

func TestCommandLineHelper_Process(t *testing.T) {
	tests := []struct {
		name    string
		helper  *helpers.CommandLineHelper
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helper := &helpers.CommandLineHelper{}
			if err := helper.Process(); (err != nil) != tt.wantErr {
				t.Errorf("CommandLineHelper.Process() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
