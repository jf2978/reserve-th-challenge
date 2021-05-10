package main

import (
	"fmt"
)

var (
	opening = map[string]string{
		"[": "]",
		"{": "}",
		"(": ")",
	}

	closing = map[string]string{
		"]": "[",
		"}": "{",
		")": "(",
	}
)

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
func (s *Stack) Top() string {
	if s.IsEmpty() {
		return ""
	}

	index := len(*s) - 1
	return (*s)[index]
}

// Pop returns and removes the top element on this stack.
func (s *Stack) Pop() string {
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]

	return element
}

func main() {
	s1 := ""
	s2 := "]-("
	s3 := "} { [()] (()v()) } ])"
	s4 := "([{]]}"

	// test cases from the problem statement

	res, err := nest(s1)
	if err != nil {
		panic(err)
	}

	fmt.Printf("properly nested string: %s\n", res)

	res, err = nest(s2)
	if err != nil {
		panic(err)
	}

	fmt.Printf("properly nested string: %s\n", res)

	res, err = nest(s3)
	if err != nil {
		panic(err)
	}

	fmt.Printf("properly nested string: %s\n", res)

	res, err = nest(s4)
	if err != nil {
		panic(err)
	}

	fmt.Printf("properly nested string: %s\n", res)
}

// nest returns the shortest properly nested string that contains s as a substring.
func nest(s string) (string, error) {
	// edge cases: empty string
	if s == "" {
		return "", nil
	}

	var st Stack
	res := ""
	// iterate through string characters
	for i, c := range s {
		char := string(c)

		// if an opening char -> add it to the stack
		if _, ok := opening[char]; ok {
			st.Push(char)
		}

		// if close char && stack is empty -> append char to the beginning of the result string
		if val, ok := closing[char]; ok && st.IsEmpty() {
			res = fmt.Sprintf("%s%s", val, res)
		}

		if val, ok := closing[char]; ok && !st.IsEmpty() {
			// if the top nests our current char, we can just pop it
			if st.Top() == val {
				st.Pop()
			} else if st.Top() != val && i == len(s)-1 {
				res = fmt.Sprintf("%s%s", val, res)
			} else { // otherwise return an error because the substring itself is a mismatch
				return "", fmt.Errorf("substring %s can not be properly nested", s)
			}
		}

		res = fmt.Sprintf("%s%s", res, char)
	}

	// after iteration -> for every open char still on the stack, pop them until we have a properly nested result string
	// this will result in a properly nested string because lingering open braces can be matched at the end
	for !st.IsEmpty() {
		res = fmt.Sprintf("%s%s", res, opening[st.Pop()])
	}

	return res, nil
}
