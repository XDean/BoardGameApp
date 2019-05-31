package util

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type J map[string]interface{}

func TestStructContain(t *testing.T) {
	ok, err := StructContain(J{
		"a": "a",
		"b": J{
			"b1": []interface{}{"b1"},
			"b2": "b2",
		},
		"c": "c",
	}, J{
		"a": "a",
		"b": J{
			"b1": []interface{}{"b1"},
		},
	})
	assert.True(t, ok)
	assert.NoError(t, err)
}

func ExampleStructContain() {
	fmt.Println(StructContain(J{
		"b": J{
			"b1": []interface{}{},
		},
	}, J{
		"b": J{
			"b1": []interface{}{"b1"},
		},
	}))

	fmt.Println(StructContain(J{
		"b": J{
			"b1": []interface{}{"b1"},
		},
	}, J{
		"b": J{
			"b1": []interface{}{"b2"},
		},
	}))

	fmt.Println(StructContain(1, 2))
	fmt.Println(StructContain(1, 1))

	fmt.Println(StructContain(nil, nil))
	fmt.Println(StructContain(nil, 1))
	fmt.Println(StructContain(1, nil))

	fmt.Println(StructContain(1, ""))

	type (
		A struct {
			A int
		}
		AB struct {
			A int
			B int
		}
		B struct {
			B int
		}
	)

	fmt.Println(StructContain(AB{1, 2}, A{1}))
	fmt.Println(StructContain(B{1}, A{1}))
	fmt.Println(StructContain(AB{1, 2}, B{1}))

	fmt.Println(StructContain(&AB{1, 2}, &A{1}))
	fmt.Println(StructContain(&B{1}, &A{1}))

	fmt.Println(StructContain(J{"a": 1}, J{"a": 1}))
	fmt.Println(StructContain(J{"a": 1}, J{"b": 1}))
	fmt.Println(StructContain(J{}, J{"b": 1}))

	fmt.Println(StructContain([]int{1}, []int{1}))
	fmt.Println(StructContain([]int{2}, []int{1}))
	fmt.Println(StructContain([]int{}, []int{1}))

	// Output:
	// false b->b1->Expect len 1 but 0
	// false b->b1->0->Expect b2 but b1
	// false Expect 2 but 1
	// true <nil>
	// true <nil>
	// false Expect non-null but actual null
	// false Expect null but actual not
	// false Expect string but int
	// true <nil>
	// false Expect field A but not present
	// false B->Expect 1 but 2
	// true <nil>
	// false Expect field A but not present
	// true <nil>
	// false Expect key b but not present
	// false Expect key b but not present
	// true <nil>
	// false 0->Expect 1 but 2
	// false Expect len 1 but 0
}
