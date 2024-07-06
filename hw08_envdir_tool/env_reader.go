package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
)

type Environment map[string]EnvValue

type EnvValue struct {
	Value      string
	NeedRemove bool
}

var (
	ErrCannotConvertValueToBytes = errors.New("cannot convert value to bytes")
	ErrOffsetExceedsFileSize     = errors.New("offset exceeds file size")
)

func proceedValue(val interface{}) (string, error) {
	var bytesValue []byte
	switch recievedValue := val.(type) {
	case string:
		bytesValue = []byte(recievedValue)
	case []byte:
		bytesValue = recievedValue
	default:
		return "", ErrCannotConvertValueToBytes
	}
	bytesValue = bytes.TrimRight(bytesValue, " \t")
	bytesValue = bytes.ReplaceAll(bytesValue, []byte{0x00}, []byte("\n"))

	return string(bytesValue), nil
}

func ReadDir(dir string) (Environment, error) {
	envs := Environment{}

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("error when trying to read the directory: %w", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		filename := entry.Name()
		filepath := dir + "/" + filename
		file, err := os.Open(filepath)
		if err != nil {
			return nil, fmt.Errorf("error when trying to open the file: %w", err)
		}
		scanner := bufio.NewScanner(file)
		isScanned := scanner.Scan()
		file.Close()

		if isScanned {
			value := scanner.Text()

			value, err = proceedValue(value)
			if err != nil {
				return nil, fmt.Errorf("cannot proceed value in file: %w", err)
			}
			envs[filename] = EnvValue{Value: value, NeedRemove: false}
		} else {
			envs[filename] = EnvValue{Value: "", NeedRemove: true}
		}
	}
	return envs, nil
}
