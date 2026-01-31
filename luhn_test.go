package main

import "testing"

func TestLuhn(t *testing.T) {
	luhnTests := []struct {
		number string
		result bool
	}{
		{number: "4532015112830366", result: true},
		{number: "1234567890123456", result: false},
	}

	for _, lt := range luhnTests {
		t.Run(lt.number, func(t *testing.T) {
			got := validateLuhn(lt.number)
			if got != lt.result {
				t.Errorf("%s got %v, want %v", lt.number, got, lt.result)
			}
		})
	}
}
