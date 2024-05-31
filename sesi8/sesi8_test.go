package main

import "testing"

func Test(t *testing.T) {
	type input struct {
		a, b int
	}

	testCases := []struct {
		desc string
		in   input
		out  int
	}{
		{
			desc: "add negative number",
			in:   input{a: -1, b: -2},
			out:  -3,
		},
		{
			desc: "add negative in positive number",
			in:   input{a: -1, b: 2},
			out:  1,
		},
		{
			desc: "add positive in negative number",
			in:   input{a: 1, b: -2},
			out:  -1,
		},
		{
			desc: "add positive number",
			in:   input{a: 1, b: 2},
			out:  3,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			output := Add(tC.in.a, tC.in.b)
			if output != tC.out {
				t.Errorf("mismatch output, expected %v actual %v", tC.out, output)
			}
		})
	}
}
