package main

import "testing"

func TestStorage(t *testing.T) {
	storageTests := []struct {
		name     string
		criteria string
		want     int
	}{
		{"ere it should return one result", "ere", 1},
		{"tsf it should return two results", "tsf", 2},
		{"12 it should return one result", "12", 1},
		{"121 it should return one result", "12", 1},
	}

	db, _ := NewStorage(Memory)

	for _, tt := range storageTests {
		t.Run(tt.name, func(t *testing.T) {
			results, _ := db.Find(tt.criteria)
			got := len(results)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
