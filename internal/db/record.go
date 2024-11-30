package db

import (
	"bytes"
	"encoding/gob"
)

type Record struct {
	Key     string
	Value   string
	Deleted bool
}

func EncodeRecord(record Record) ([]byte, error) {
	var buffer bytes.Buffer

	encoder := gob.NewEncoder(&buffer)
	if err := encoder.Encode(record); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func DecodeRecord(data []byte) (Record, error) {
	var record Record
	buffer := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buffer)
	if err := decoder.Decode(&record); err != nil {
		return record, err
	}
	return record, nil
}
