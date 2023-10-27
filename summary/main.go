package summary

import (
	"errors"
	"fmt"
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

type Message struct {
	Service      string
	ResourceName string
	Region       string
}

// GetContentsFromFile returns a FileContents struct with the contents of the file
func GetContentsFromFile(filepath string) (contents FileContents, err error) {
	contents.Filepath = filepath
	contents.Sha256sum, err = file.CalculateSHA256Sum(filepath)
	contents.AllLines, err = file.ReadFileLines(filepath)
	for _, line := range contents.AllLines {
		// don't look for messages in lines that talk about enabled rregiong
		// ..level=info msg="Identifying enabled regions"..
		// ..msg="Found enabled region ap-south-1"..
		if strings.Contains(line, "enabled region") {
			contents.StatusLines = append(contents.StatusLines, line)
			continue
		}

		// don't look for mesages in lines that  list service checks
		// ..msg="- acm"..
		if strings.Contains(line, "msg=\"- ") {
			contents.StatusLines = append(contents.StatusLines, line)
			continue
		}

		if strings.Contains(line, "msg=") {

			msg, err := ExtractStringBetweenTwoSubstrings(line, "msg=\"", "\"")
			if err != nil {
				panic(errors.New(fmt.Sprintf("malformed message line: %v", msg)))
			}
			contents.MessageLines = append(contents.MessageLines, msg)
		} else {
			contents.StatusLines = append(contents.StatusLines, line)
		}

	}
	return contents, err
}

// ExtractStringBetweenTwoSubstrings returns a string between two substrings
func ExtractStringBetweenTwoSubstrings(input, start, end string) (result string, err error) {
	startIndex := strings.Index(input, start)
	if startIndex == -1 {
		return "", errors.New("Beginning substring not found")
	}

	endIndex := strings.Index(input[startIndex+len(start):], end)
	if endIndex == -1 {
		return "", errors.New("Ending substring not found")
	}

	extracted := input[startIndex+len(start) : startIndex+len(start)+endIndex]
	return extracted, nil
}

// GetMessage returns a Message struct with the service, resource name, and region
func GetMessage(input string) (message Message, err error) {
	words := strings.Fields(input)
	service := strings.TrimPrefix(words[0], "\\x1b[1;")
	service = strings.TrimSuffix(service, "\\x1b[0m")

	region := words[len(words)-1]

	return Message{
		Service:      service,
		ResourceName: words[len(words)-2], // slice with first and last words removed
		Region:       strings.TrimSuffix(region, "\\n"),
	}, err
}

func PrintResourcesCountByRegion(fileContents FileContents) {
	result := make(map[string]int)
	for _, messageLine := range fileContents.MessageLines {
		msg, _ := GetMessage(messageLine)
		result[msg.Region]++
	}
	for region, count := range result {
		out := fmt.Sprintf("%v: %v", region, count)
		fmt.Println(out)
	}
	fmt.Println()
}

// PrintReport prints a report of the file contents
func PrintReport(fileContents FileContents) {
	PrintResourcesCountByRegion(fileContents)
	for _, messageLine := range fileContents.MessageLines {
		msg, _ := GetMessage(messageLine)
		fmt.Println(msg.Service+" : ", msg.ResourceName+" : ", msg.Region)
	}
}
