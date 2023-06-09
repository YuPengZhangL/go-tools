package json

import (
	"fmt"
	"testing"
)

func TestFlattenJSON(t *testing.T) {
	demo := map[string]interface{}{
		"test_int":   1,
		"test2_str":  "one",
		"test_float": 1.11,
		"test_map":   map[string]interface{}{"test_map1": map[string]interface{}{"test_map2": "test_map2", "test_map3": "test_map3"}},
	}
	test := JSONFlattener{}
	if err := test.FullFlattenJSON("", demo, true, true); err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%v\n%v\n", test, demo)
}
