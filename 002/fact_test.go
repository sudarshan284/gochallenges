package main

import "testing"

func TestEx002(t *testing.T) {
	want := 120
	got := Factorial(5)

	if got != want {
		t.Errorf("Factorial () = %v, want %v", got, want)
	}

	want = 40320
	got = Factorial(8)
	if got != want {
		t.Errorf("Ex002() = %v, want %v", got, want)
	}
}
