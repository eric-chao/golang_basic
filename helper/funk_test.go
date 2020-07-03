package helper

import (
	"github.com/thoas/go-funk"
	"testing"
)

func TestToMap(t *testing.T) {
	var elements = []string{"abc", "def", "fgi", "adi"}

	elementsMap := funk.Map(
		funk.Chunk(elements, 2),
		func(x []string) (string, string) { // Slice to Map
			return x[0], x[1]
		},
	)

	t.Log(elementsMap) // map[abc:def fgi:adi]

	var ids = []int64{
		1, 2, 3, 4, 5,
	}
	idMap := funk.Map(
		ids,
		func(x int64) (int64, interface{}) { // Slice to Map
			return x, nil
		},
	).(map[int64]interface{})

	for k, v := range idMap {
		t.Logf("key = %v, value = %v", k, v)
	}

	var gradeIdMap map[int64]interface{}
	if len(gradeIdMap) > 0 {
		t.Logf("Map %v", gradeIdMap)
	}
}