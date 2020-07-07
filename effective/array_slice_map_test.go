package effective

import (
	"fmt"
	"testing"
)

const (
	Enone = iota
	Eio
	Einval
)

func TestPrint(t *testing.T) {
	a := [...]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
	s := []string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
	m := map[int]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}

	for idx, value := range a {
		fmt.Printf("[array] index: %d, value: %s \n", idx, value)
	}

	for idx, value := range s {
		fmt.Printf("[slice] index: %d, value: %s \n", idx, value)
	}

	for key, value := range m {
		fmt.Printf("[map] key: %v, value: %s \n", key, value)
	}
}
