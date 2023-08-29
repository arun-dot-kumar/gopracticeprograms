package main

import "testing"

func TestStruct(t *testing.T) {
	s := Student{
		Name:   "arun",
		RollNo: 2,
	}
	err := s.Validates()
	if err != nil {
		t.Fatal(err)
	}

}
