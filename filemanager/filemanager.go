package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)
type FileManager struct{
	Inputmanager string
	Outputmanager string
}

func (fm FileManager)FileRead() ([]string, error) {
	file, err := os.Open(fm.Inputmanager)
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

func (fm FileManager) WriteResult( data any) error {
	file, err := os.Create(fm.Outputmanager)
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
func New(inputpath,outputpath string)FileManager{
return FileManager{
	Inputmanager: inputpath,
	Outputmanager: outputpath,
}
}
