package main

import (
	"bufio"
	"fmt"
	"os"
)

type stack struct {
	chars []string
	count int
}

func (s *stack) push(char string) {

	if char == "[" || char == "]" || char == "(" || char == ")" || char == "{" || char == "}" {
		s.chars = append(s.chars, char)
		s.count++
	}

	fmt.Printf("s count is now %d after push \n", s.count)

}

func (s *stack) pop() string {
	var lastChar string
	fmt.Printf("s count is now %d\n", s.count)
	if s.count >= 1 {

		lastChar = s.chars[s.count-1]
		fmt.Printf("char to return %s\n", lastChar)
		s.count--
	}
	return lastChar
}

func (s *stack) checkBalancedBrackets(input string) bool {

	for i := 0; i < len(input)-1; i++ {
		s.push(string(input[i]))
	}

	var lastBracket string

	for lastBracket == '['
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')
		fmt.Println(text)
		s := stack{}
		result := s.checkBalancedBrackets(text)
		if !result {
			fmt.Print("balanced\n")
		} else {
			fmt.Print("unbalanced\n")
		}
	}
}
