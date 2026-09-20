package twosum

import (
	"encoding/json"
	"fmt"
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
	Nums     []int  `json:"nums"`
	Target   int    `json:"target"`
	Expected []int  `json:"expected"`
}

type pairAt string

const (
	pairAtStart pairAt = "start"
	pairAtEnd   pairAt = "end"
)

type scaleSpec struct {
	Name   string `json:"name"`
	N      int    `json:"n"`
	Fill   int    `json:"fill"`
	Left   int    `json:"left"`
	Right  int    `json:"right"`
	Target int    `json:"target"`
	PairAt pairAt `json:"pair_at"`
}

func expandScale(spec scaleSpec) (literalCase, error) {
	nums := make([]int, spec.N)
	for i := range nums {
		nums[i] = spec.Fill
	}

	var expected []int
	switch spec.PairAt {
	case pairAtStart:
		expected = []int{0, 1}
	case pairAtEnd:
		expected = []int{spec.N - 2, spec.N - 1}
	default:
		return literalCase{}, fmt.Errorf("unknown pair_at: %q", spec.PairAt)
	}

	nums[expected[0]] = spec.Left
	nums[expected[1]] = spec.Right
	return literalCase{
		Name:     spec.Name,
		Nums:     nums,
		Target:   spec.Target,
		Expected: expected,
	}, nil
}

func sameIndices(actual, expected []int) bool {
	if len(actual) != 2 || len(expected) != 2 {
		return false
	}
	return (actual[0] == expected[0] && actual[1] == expected[1]) ||
		(actual[0] == expected[1] && actual[1] == expected[0])
}

func TestTwoSum(t *testing.T) {
	var cases []literalCase
	if err := json.Unmarshal(casesJSON, &cases); err != nil {
		t.Fatalf("cases.json: %v", err)
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			got := twoSum(c.Nums, c.Target)
			if !sameIndices(got, c.Expected) {
				t.Fatalf("got %v, want %v (either order)", got, c.Expected)
			}
		})
	}
}

func TestTwoSumScale(t *testing.T) {
	var specs []scaleSpec
	if err := json.Unmarshal(scaleJSON, &specs); err != nil {
		t.Fatalf("scale.json: %v", err)
	}

	for _, spec := range specs {
		c, err := expandScale(spec)
		if err != nil {
			t.Fatal(err)
		}

		t.Run(c.Name, func(t *testing.T) {
			started := time.Now()
			got := twoSum(c.Nums, c.Target)
			elapsed := time.Since(started)

			if !sameIndices(got, c.Expected) {
				t.Fatalf("got %v, want %v (either order)", got, c.Expected)
			}
			t.Logf("%s: %s", c.Name, elapsed)
		})
	}
}
