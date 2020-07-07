package json

import (
	jsoniter "github.com/json-iterator/go"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}

func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
