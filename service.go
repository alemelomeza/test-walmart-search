package main

import (
	"strings"
)

const (
	// % of maximum discount
	limitRateDiscount = 90
	// % of discount
	palindromeCampaing = 50
)

// Campaing return a discount
func Campaing(q string) int {
	var discount int

	// Applay discounts
	if isPalindrome(q) {
		discount += palindromeCampaing
	}
	// Discount limit
	if discount > limitRateDiscount {
		discount = limitRateDiscount
	}

	return discount
}

// normalize func normalizes a string
func normalize(s string) string {
	lc := strings.ToLower(s)
	ts := strings.TrimSpace(lc)
	rw := strings.ReplaceAll(ts, " ", "")

	return rw
}

// reverse func reverse a string
func reverse(s string) string {
	data := []rune(s)
	result := []rune{}

	for i := len(data) - 1; i >= 0; i-- {
		result = append(result, data[i])
	}

	return string(result)
}

// normalizeAndReverse func normalizes and reverse a string
func normalizeAndReverse(s string) string {
	ns := normalize(s)
	rs := reverse(ns)

	return rs
}

// isPalindrome func check if the string is palindrome
func isPalindrome(s string) bool {
	if len(s) < 3 {
		return false
	}

	sn := normalize(s)
	snr := normalizeAndReverse(s)

	return bool(sn == snr)
}
