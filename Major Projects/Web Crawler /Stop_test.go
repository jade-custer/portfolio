package main

import "testing"

func TestStop(t *testing.T) {
	tests := []struct {
		word   string
		exists bool
	}{
		{
			"Romeo",
			false,
		},
		{
			"Verona",
			false,
		},
		{
			"10",
			true,
		},
		{
			"above",
			true,
		},
		{
			"abroad",
			true,
		},
	}

	for _, test := range tests {
		//checking if the output and wanted are the same
		got := Stop(test.word)

		if test.exists != got {
			t.Errorf("Wanted %v but go %v", test.exists, got)
		}
	}

}
