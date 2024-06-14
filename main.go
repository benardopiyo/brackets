package main

import (
	"os"
	"github.com/01-edu/z01"
)

func isValidBrackets(expr string) bool {
	stack := []rune{}
	bracketPairs := map[rune]rune{'(': ')', '[': ']', '{': '}'}

	for _, char := range expr {
		if closingBracket, ok := bracketPairs[char]; ok {
			// If it's an opening bracket, push its corresponding closing bracket onto the stack
			stack = append(stack, closingBracket)
		} else if char == ')' || char == ']' || char == '}' {
			// It's a closing bracket, check if it matches the top of the stack
			if len(stack) == 0 || stack[len(stack)-1] != char {
				return false
			}
			stack = stack[:len(stack)-1] // pop from stack
		}
	}
	return len(stack) == 0
}

func main() {
	args := os.Args[1:]
	for _, arg := range args {
		if isValidBrackets(arg) {
			for _, r := range "OK" {
				z01.PrintRune(r)
			}
		} else {
			for _, r := range "Error" {
				z01.PrintRune(r)
			}
		}
		z01.PrintRune('\n')
	}
}
