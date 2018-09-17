package day27

import "strings"

// WellBalanced returns true if the brackets in the string are balanced and well-formed.
// Runs in O(N) time where N is the number of runes. Using O(N/2) extra space.
func WellBalanced(brackets string) bool {
	openBrackets := make([]rune, 0, (len(brackets)/2)+1)
	for _, r := range brackets {
		if strings.ContainsRune("([{", r) {
			openBrackets = append(openBrackets, r)
		} else if strings.ContainsRune(")]}", r) {
			if len(openBrackets) == 0 {
				return false
			}
			var pop rune
			pop, openBrackets = openBrackets[len(openBrackets)-1], openBrackets[:len(openBrackets)-1]
			if !IsMatch(pop, r) {
				return false
			}
		}
		if len(openBrackets) > len(brackets)/2 {
			return false
		}
	}
	return len(openBrackets) == 0
}

// IsMatch returns true if the left and right runes form a matching pair.
func IsMatch(left, right rune) bool {
	return (right == ')' && left == '(') ||
		(right == ']' && left == '[') ||
		(right == '}' && left == '{')
}
