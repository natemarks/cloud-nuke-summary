package summary

import (
	"strings"

	"github.com/natemarks/cloud-nuke-summary/file"
)

type FileContents struct {
	Filepath     string   // path to the result file
	Sha256sum    string   // sha256sum of the file
	StatusLines  []string // First lines before the message lines
	AllLines     []string // all lines and all content
	MessageLines []string //message lines with only the  message content
}

// GetContentsFromFile returns a FileContents struct with the contents of the file
func GetContentsFromFile(filepath string) (contents FileContents, err error) {
	contents.Filepath = filepath
	contents.Sha256sum, err = file.CalculateSHA256Sum(filepath)
	contents.AllLines, err = file.ReadFileLines(filepath)
	for _, line := range contents.AllLines {
		if strings.Contains(line, "msg=") {
			contents.MessageLines = append(contents.MessageLines, line)
		} else {
			contents.StatusLines = append(contents.StatusLines, line)
		}

	}
	return contents, err
}
