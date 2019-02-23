package main

import (
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"testing..."},
		{"2testing..."},
		{"3testing..."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if tt.name != "testing..." {
				t.Errorf("Silly test value %s is not the same as %s", tt.name, tt.name)
			}

		})
	}
}
