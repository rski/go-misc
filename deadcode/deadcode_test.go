package main

import (
	"go/types"
	"testing"
)

func TestDeadcode(t *testing.T) {
	// ignoring "h" because recursive funcs are nut supported
	tests := []struct {
		testFile  string
		names     []string
		withTests bool
	}{{
		testFile: "p1",
		names:    []string{"unused", "g", "H"},
	}, {
		testFile: "p2",
		names:    []string{"main", "unused", "g"},
	}, {
		testFile:  "p3",
		names:     []string{"y"},
		withTests: true,
	},
	}

	for _, tc := range tests {
		t.Run(tc.testFile, func(t *testing.T) {
			ctx := &context{withTests: tc.withTests}
			ctx.load("./testdata/" + tc.testFile)
			objs := ctx.process()
			compare(t, objs, tc.names)
		})
	}

}

func compare(t *testing.T, objs []types.Object, names []string) {
	left := make(map[string]bool)
	right := make(map[string]bool)
	for _, o := range objs {
		left[o.Name()] = true
	}
	for _, n := range names {
		right[n] = true
	}

	for _, o := range objs {
		if !right[o.Name()] {
			t.Errorf("%s should not have been reported as unused", o.Name())
		}
	}
	for _, n := range names {
		if !left[n] {
			t.Errorf("unused %s should not have been reported", n)
		}
	}
}
