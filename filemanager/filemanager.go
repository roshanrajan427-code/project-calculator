package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

func FileRead(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, errors.New("failed to open file")
	}
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("failed to read file")
	}
	file.Close()
	return lines, err
}

func WriteJson(path string, data any) error {
	file, err := os.Create(path)
	if err != nil {
		return errors.New("can't create a file")
	}
	encoder:=json.NewEncoder(file)
	err=encoder.Encode(data)
	if err != nil {
		file.Close()
		return errors.New("can't conver to json")
	}
	file.Close()
	return nil
}
