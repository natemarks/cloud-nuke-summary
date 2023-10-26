package summary

import (
	"reflect"
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

func TestExtractStringBetweenTwoSubstrings(t *testing.T) {
	type args struct {
		input string
		start string
		end   string
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
		wantErr    bool
	}{{
		name: "sdf",
		args: args{
			input: "fghrhryjytjyjaaahyhyhyhybbbjyutkukuki",
			start: "aaa",
			end:   "bbb"},
		wantResult: "hyhyhyhy",
		wantErr:    false,
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := ExtractStringBetweenTwoSubstrings(tt.args.input, tt.args.start, tt.args.end)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractStringBetweenTwoSubstrings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResult != tt.wantResult {
				t.Errorf("ExtractStringBetweenTwoSubstrings() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestGetMessage(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name        string
		args        args
		wantMessage Message
		wantErr     bool
	}{
		{
			name: "sdf",
			args: args{
				input: "\\x1b[1;mcloudwatch-dashboard\\x1b[0m EcsLoadTesting-CsP9WPSdg2-us-east-2 us-west-2\\n",
			},
			wantMessage: Message{
				Service:      "mcloudwatch-dashboard",
				ResourceName: "EcsLoadTesting-CsP9WPSdg2-us-east-2",
				Region:       "us-west-2",
			},
			wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMessage, err := GetMessage(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotMessage, tt.wantMessage) {
				t.Errorf("GetMessage() gotMessage = %v, want %v", gotMessage, tt.wantMessage)
			}
		})
	}
}
