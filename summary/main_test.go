package summary

import (
	"testing"

	"github.com/natemarks/cloud-nuke-summary/projectpath"
)

func TestGetContentsFromFile(t *testing.T) {
	type args struct {
		filepath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "sdf", args: args{filepath: projectpath.Root + "/cloud-nuke-pipeline-out-small.txt"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotContents, err := GetContentsFromFile(tt.args.filepath)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetContentsFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !(len(gotContents.AllLines) > 50) {
				t.Errorf("GetContentsFromFile() error: expected more than 50 ines. got = %v", len(gotContents.AllLines))
			}
		})
	}
}
