package main

import "fmt"

func reverseString(s []byte) {
	l := len(s)
	for i := 0; i < l/2; i++ {
		s[i], s[l-i-1] = s[l-i-1], s[i]
	}

	fmt.Println(string(s))
}

func main() {
	reverseString([]byte(`hello`))
	reverseString([]byte(`hello there`))
}
