package main

import (
	"testing"
)

func TestAscDescCheck(t *testing.T) {
	got := checkReports([][]string{{"7", "6", "4", "2", "1"},
		{"1", "3", "2", "4", "5"},
		{"1", "3", "6", "7", "9"},
		{"8", "6", "4", "4", "1"},
	})
	want := 4
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestDiffMax(t *testing.T) {
	got := checkReports([][]string{{"1", "2", "7", "8", "9"}, {"9", "7", "6", "2", "1"}})
	want := 0
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestFile(t *testing.T) {
	reportList := createListOfReports("test_file.txt")
	got := checkReports(reportList)
	want := 4
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestFullFile(t *testing.T) {
	reportList := createListOfReports("full_input.txt")
	got := checkReports(reportList)
	want := 606
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
