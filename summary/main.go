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
			msg, _ := ExtractStringBetweenTwoSubstrings(line, "msg=\"", "\"")
			contents.MessageLines = append(contents.MessageLines, msg)
		} else {
			contents.StatusLines = append(contents.StatusLines, line)
		}

	}
	return contents, err
}

// ExtractStringBetweenTwoSubstrings returns a string between two substrings
func ExtractStringBetweenTwoSubstrings(input, start, end string) (result string, found bool) {
	startIndex := strings.Index(input, start)
	if startIndex == -1 {
		return "", false // Beginning substring not found
	}

	endIndex := strings.Index(input[startIndex+len(start):], end)
	if endIndex == -1 {
		return "", false // Ending substring not found
	}

	extracted := input[startIndex+len(start) : startIndex+len(start)+endIndex]
	return extracted, true
}
