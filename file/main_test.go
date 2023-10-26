package file

import (
	"testing"

	"github.com/natemarks/cloud-nuke-summary/projectpath"
)

func TestReadFileLines(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "sdf", args: args{filename: projectpath.Root + "/cloud-nuke-pipeline-out-small.txt"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadFileLines(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFileLines() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !(len(got) > 50) {
				t.Errorf("ReadFileLines() error: expected more than 50 ines. got = %v", len(got))
			}
		})
	}
}
