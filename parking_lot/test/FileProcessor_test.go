package test

import (
	"testing"
	helpers "../internal/app/helpers"
)


func TestFileProcessor_Process(t *testing.T) {
	type fields struct {
		FileName string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "FileTest",
			fields: fields{
				FileName:"/home/rupalib/Downloads/parking-lot-1.4.1/parking_lot/cmd/commands.txt",
			},
			wantErr:false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fileProcessor := &helpers.FileProcessor{tt.fields.FileName}
			if err := fileProcessor.Process(); (err != nil) != tt.wantErr {
				t.Errorf("FileProcessor.Process() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
