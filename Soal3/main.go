package main

import "fmt"

func isValid(input string) bool {
	stack := make([]rune, 0)
	mapping := map[rune]rune{
		'<': '>',
		'{': '}',
		'[': ']',
	}

	for _, char := range input {
		if char == '<' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else if char == '>' || char == '}' || char == ']' {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if mapping[top] != char {
				return false
			}
		} else {
			return false
		}
	}

	return len(stack) == 0
}

func main() {
	var input string

	fmt.Printf("Masukan string: ")
	fmt.Scan(&input)

	valid := isValid(input)
	if valid {
		fmt.Println("TRUE")
	} else {
		fmt.Println("FALSE")
	}
}
