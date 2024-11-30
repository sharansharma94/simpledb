package db

import (
	"bufio"
	"errors"
	"os"
)

type Database struct {
	filePath string
}

func NewDatabase(filePath string) (*Database, error) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	return &Database{
		filePath: filePath,
	}, nil
}

func (db *Database) Write(key, value string, deleted bool) error {

	record := Record{Key: key, Value: value, Deleted: deleted}
	data, err := EncodeRecord(record)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(db.filePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(append(data, '\n'))
	return err
}

func (db *Database) Read(key string) (string, error) {
	file, err := os.Open(db.filePath)
	if err != nil {
		return "", err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var foundValue string
	for scanner.Scan() {
		line := scanner.Bytes()
		record, err := DecodeRecord(line)
		if err != nil {
			continue
		}

		if record.Key == key {
			if record.Deleted {
				foundValue = ""
			} else {
				foundValue = record.Value
			}
		}
	}
	if foundValue == "" {
		return "", errors.New("key not found")
	}
	return foundValue, nil
}

func (db *Database) Delete(key string) error {
	record := Record{Key: key, Deleted: true}
	data, err := EncodeRecord(record)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(db.filePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(append(data, '\n'))
	return err
}
