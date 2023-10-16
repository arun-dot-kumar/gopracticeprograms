package main

import "fmt"

type stack struct {
	list []string
}

func (s *stack) push(data string) {
	s.list = append(s.list, data)
}

func (s *stack) pop() {
	s.list = s.list[:len(s.list)-1]
}

func (s *stack) getTopElement() string {
	return s.list[len(s.list)-1]
}

func main() {
	var s = "{[[{]"
	var st stack
	var valid = true
	for _, char := range s {
		if string(char) == "{" || string(char) == "[" || string(char) == "(" {
			st.push(string(char))
		} else {
			if string(char) == "}" {
				if st.getTopElement() != "{" {
					valid = false
					break
				} else {
					st.pop()
				}
			}
			if string(char) == "]" {
				if st.getTopElement() != "[" {
					valid = false
					break
				} else {
					st.pop()
				}
			}
			if string(char) == ")" {
				if st.getTopElement() != "(" {
					valid = false
					break
				} else {
					st.pop()
				}
			}
		}
	}
	fmt.Println(valid)
}
