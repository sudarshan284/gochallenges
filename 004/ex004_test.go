package main

import (
	"reflect"
	"testing"
)

func TestEx004(t *testing.T) {
	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	got := Ex004("1 2 3 4 5 6 7 8 9")

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Ex004() = %v; want %v", got, want)
	}
}
