// Package utils contains function to be used across the application
package utils

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func HandleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func ToBytes(i interface{}) []byte {
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	HandleErr(enc.Encode(i))
	return buffer.Bytes()
}

// FromBytes takes an interface and data and then will encode the data
func FromBytes(i interface{}, data []byte) {
	decoder := gob.NewDecoder(bytes.NewReader(data))
	HandleErr(decoder.Decode(i))
}

// Hash takes an interface, hashes it
func Hash(i interface{}) string {
	str := fmt.Sprint(i)
	return fmt.Sprintf("%x", sha256.Sum256([]byte(str)))
}

func Splitter(s string, sep string, i int) string {
	r := strings.Split(s, sep)
	if len(r)-1 < i {
		return ""
	}
	return r[i]
}

func ToJSON(i interface{}) []byte {
	r, err := json.Marshal(i)
	HandleErr(err)
	return r
}
