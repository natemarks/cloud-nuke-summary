package file

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
)

// ReadFileLines returns a slice of strings, each string is a line from the file
func ReadFileLines(filename string) ([]string, error) {
	// Open the file for reading
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	var lines []string

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// Append each line to the slice
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// CalculateSHA256Sum returns the sha256sum of a file
func CalculateSHA256Sum(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	checksum := hex.EncodeToString(hash.Sum(nil))
	return checksum, nil
}
