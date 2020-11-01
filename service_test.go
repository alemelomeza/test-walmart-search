package main

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	assertMessage := func(t *testing.T, got, want bool) {
		t.Helper()
		if got != want {
			t.Errorf("got %t, want %t", got, want)
		}
	}

	t.Run("aa it should not be palindrome", func(t *testing.T) {
		got := isPalindrome("aa")
		want := false

		assertMessage(t, got, want)
	})

	t.Run("aaa it should be a palindrome", func(t *testing.T) {
		got := isPalindrome("aaa")
		want := true

		assertMessage(t, got, want)
	})

	t.Run("abc it should not be a palindrome", func(t *testing.T) {
		got := isPalindrome("abc")
		want := false

		assertMessage(t, got, want)
	})

	t.Run("12 it should not be a palindrome", func(t *testing.T) {
		got := isPalindrome("aa")
		want := false

		assertMessage(t, got, want)
	})

	t.Run("121 it should be a palindrome", func(t *testing.T) {
		got := isPalindrome("121")
		want := true

		assertMessage(t, got, want)
	})

	t.Run("123 it should not be a palindrome", func(t *testing.T) {
		got := isPalindrome("123")
		want := false

		assertMessage(t, got, want)
	})
}

func TestCampaing(t *testing.T) {
	tests := []struct {
		criteria string
		query    string
		want     int
	}{
		{"aa it should be 0", "aa", 0},
		{"aaa it should be 50", "aaa", 50},
		{"abc it should be 0", "abc", 0},
		{"12 it should be 0", "12", 0},
		{"121 it should be 50", "121", 50},
		{"123 it should be 0", "123", 0},
	}

	for _, tt := range tests {
		t.Run(tt.criteria, func(t *testing.T) {
			got := Campaing(tt.query)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
