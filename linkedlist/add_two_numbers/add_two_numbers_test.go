package addtwonumbers

import (
	"encoding/json"
	"testing"
	"time"

	_ "embed"
)

//go:embed cases.json
var casesJSON []byte

//go:embed scale.json
var scaleJSON []byte

type literalCase struct {
	Name     string `json:"name"`
	L1       []int  `json:"l1"`
	L2       []int  `json:"l2"`
	Expected []int  `json:"expected"`
}

type scaleSpec struct {
	Name   string `json:"name"`
	L1Len  int    `json:"l1_len"`
	L1Fill int    `json:"l1_fill"`
	L2Len  int    `json:"l2_len"`
	L2Fill int    `json:"l2_fill"`
}

func expandScale(spec scaleSpec) literalCase {
	l1 := filled(spec.L1Len, spec.L1Fill)
	l2 := filled(spec.L2Len, spec.L2Fill)
	return literalCase{
		Name:     spec.Name,
		L1:       l1,
		L2:       l2,
		Expected: schoolbookAdd(l1, l2),
	}
}

func filled(n, fill int) []int {
	out := make([]int, n)
	for i := range out {
		out[i] = fill
	}
	return out
}

func schoolbookAdd(l1, l2 []int) []int {
	var out []int
	carry := 0
	i := 0
	for i < len(l1) || i < len(l2) || carry > 0 {
		sum := carry
		if i < len(l1) {
			sum += l1[i]
		}
		if i < len(l2) {
			sum += l2[i]
		}
		out = append(out, sum%10)
		carry = sum / 10
		i++
	}
	return out
}

func toList(digits []int) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for _, d := range digits {
		tail.Next = &ListNode{Val: d}
		tail = tail.Next
	}
	return dummy.Next
}

func fromList(n *ListNode) []int {
	var out []int
	for n != nil {
		out = append(out, n.Val)
		n = n.Next
	}
	return out
}

func sameDigits(actual, expected []int) bool {
	if len(actual) != len(expected) {
		return false
	}
	for i := range actual {
		if actual[i] != expected[i] {
			return false
		}
	}
	return true
}

func TestAddTwoNumbers(t *testing.T) {
	var cases []literalCase
	if err := json.Unmarshal(casesJSON, &cases); err != nil {
		t.Fatalf("cases.json: %v", err)
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			got := fromList(addTwoNumbers(toList(c.L1), toList(c.L2)))
			if !sameDigits(got, c.Expected) {
				t.Fatalf("got %v, want %v", got, c.Expected)
			}
		})
	}
}

func TestAddTwoNumbersScale(t *testing.T) {
	var specs []scaleSpec
	if err := json.Unmarshal(scaleJSON, &specs); err != nil {
		t.Fatalf("scale.json: %v", err)
	}

	for _, spec := range specs {
		c := expandScale(spec)
		t.Run(c.Name, func(t *testing.T) {
			started := time.Now()
			got := fromList(addTwoNumbers(toList(c.L1), toList(c.L2)))
			elapsed := time.Since(started)

			if !sameDigits(got, c.Expected) {
				t.Fatalf("got %v, want %v", got, c.Expected)
			}
			t.Logf("%s: %s", c.Name, elapsed)
		})
	}
}
