package part2

import "os"

type stack struct{
	slice []byte
}

func (s *stack) push(char byte){
	s.slice = append(s.slice, char)
}

func (s *stack) peek() byte{
	return s.slice[len(s.slice)-1]
}

func (s *stack)pop()byte{
	var temp byte = s.peek()
	s.slice = s.slice[0:len(s.slice)-1]
	return temp
}

func part2(f *os.File){
	
}
