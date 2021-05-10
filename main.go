package main

import "fmt"

type Stack []string

// IsEmpty checks if the current stack is empty or not.
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push adds an element to the top of the stack.
func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

// Top returns the top element without popping it.
func (s *Stack) Top(str string) string {
	index := len(*s) - 1
	return (*s)[index]
}

// Pop returns and removes the top element on this stack.
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func main() {
	s1 := ""
	s2 := "]-("

	// test cases from the problem statement
	fmt.Printf("properly nested string: %s", nest(s1))
	fmt.Printf("properly nested string: %s", nest(s2))
	//fmt.Printf("properly nested string: %s", nest(s3))
	//fmt.Printf("properly nested string: %s", nest(s4))
	//fmt.Printf("properly nested string: %s", nest(s5))

	// other cases ?
}

// TODO: initialize helper map of brace characters

// nest returns the shortest properly nested string that contains s as a substring.
func nest(s string) (string, error) {
	// edge cases: empty string
	if s == "" {
		return "", nil
	}

	var st Stack
	// iterate through string characters
	for i, c := range s {
		// if an opening char -> add it to the stack

		// if close char && stack is empty -> append char to the beginning of the result string

		// if close char && matches what's at top -> pop it from the stack
		// otherwise (not match && not empty), return an error because the substring itself is a mismatch
	}

	// after iteration -> for every open char still on the stack, pop them until we have a properly nested result string
}
