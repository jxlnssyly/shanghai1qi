package controllers

import (
	"bytes"
	"encoding/gob"
)

func Serialization(buffer *bytes.Buffer,data interface{}) error {
	enc := gob.NewEncoder(buffer)
	return enc.Encode(data)
}

func DeSerialization(data *[]byte,e interface{}) error {
	dec := gob.NewDecoder(bytes.NewBuffer(*data))
	return dec.Decode(e)
}


