package test

import (
	"../internal/app/helpers"
	"os"
	"testing"
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
				FileName: "/../cmd/commands.txt",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cwd,_:=os.Getwd()
			fileProcessor := &helpers.FileProcessor{cwd+tt.fields.FileName}
			if err := fileProcessor.Process(); (err != nil) != tt.wantErr {
				t.Errorf("FileProcessor.Process() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
