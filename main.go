package main

import "fmt"

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
// TODO: implement me
func nest(s string) (string, error) {
	// edge cases: empty string

	// initialize a stack to help us keep track of pairs
	// iterate through string characters

	// if an opening char -> add it to the stack

	// if close char && stack is empty -> append char to the beginning of the result string

	// if close char && matches what's at top -> pop it from the stack
	// otherwise (not match && not empty), return an error because the substring itself is a mismatch

	// after iteration -> for every open char still on the stack, pop them until we have a properly nested result string
}
