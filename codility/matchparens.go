package main

import "fmt"

//A string S consisting of N characters is considered to be properly nested if any of the following conditions is true:
//
//S is empty;
//S has the form "(U)" or "[U]" or "{U}" where U is a properly nested string;
//S has the form "VW" where V and W are properly nested strings.
//For example, the string "{[()()]}" is properly nested but "([)()]" is not.
//
//Write a function:
//
//func Solution(S string) int
//that, given a string S consisting of N characters, returns 1 if S is properly nested and 0 otherwise.
//
//For example, given S = "{[()()]}", the function should return 1 and given S = "([)()]", the function should return 0, as explained above.
//
//Write an efficient algorithm for the following assumptions:
//
//N is an integer within the range [0..200,000];
//string S consists only of the following characters: "(", "{", "[", "]", "}" and/or ")".
type Stack []rune

func (s *Stack) push(str rune) {
	*s = append(*s, str)
}

func (s *Stack) pop() rune {
	if s.empty() {
		return 0
	}
	index := len(*s) - 1
	last := (*s)[index]
	*s = (*s)[:index]
	return last
}

func (s *Stack) peek() rune {
	if s.empty() {
		return 0
	}
	last := (*s)[len(*s)-1]
	return last
}

func (s *Stack) empty() bool {
	return len(*s) == 0
}

func Solution(S string) int {
	if S == "" {
		return 1
	}
	st := Stack{}
	for _, c := range S {
		if c == '(' || c == '{' || c == '[' {
			st.push(c)
			continue
		}
		if st.empty() {
			return 0
		}
		top := st.peek()

		if c == ')' && top == '(' ||
			c == '}' && top == '{' ||
			c == ']' && top == '[' {
			st.pop()
		} else {
			return 0
		}
	}
	if st.empty() {
		return 1
	}
	return 0
}

func main() {
	//
	fmt.Println(Solution("((((()))))"))
	fmt.Println(Solution("(()(())())"))
	fmt.Println(Solution("(()(())())"))
	fmt.Println(Solution("(()(()())"))
	fmt.Println(Solution("{{}}()[()]"))
	fmt.Println(Solution("{][}"))
	fmt.Println(Solution("{]}"))
	fmt.Println(Solution("(])"))
	fmt.Println(Solution("()"))
}
