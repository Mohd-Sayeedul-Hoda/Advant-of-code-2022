package part1

import (
	"bufio"
	"fmt"
	"os"
)

type Stack struct{  
	slice []byte
}

func (s *Stack) PushLast(char byte){
	s.slice = append([]byte{char}, s.slice...)
}

func (s *Stack) Push(char byte){
	s.slice = append(s.slice, char)
}

func (s *Stack) Peek() byte{
	return s.slice[0]
}

func (s *Stack)Pop()byte{
	var temp byte = s.Peek()
	s.slice = s.slice[1:len(s.slice)]
	return temp
}

func Part1(file *os.File){
	input := bufio.NewScanner(file)
	arr := MakeStack(file, input)
	fmt.Println(len(arr))
	for _, v := range arr{
		for _, val := range v.slice{
			fmt.Printf("%c", val)
		}
		fmt.Printf("\n")
	}
}

func MakeStack(file *os.File,input *bufio.Scanner)[]*Stack{
	var arr []*Stack
	firstTime := true
	for input.Scan(){
		str := input.Text()
		if firstTime{
			firstTime = false
			length := len(str)
			length = int(length/4+1)
			for i:=1; i <= length; i++{
				x := new(Stack)
				arr = append(arr, x)
			}
		}
		fmt.Println(str)
		if string(str[1]) == "1"{
			return arr
		}
		s := 0 // for stack value
		for i := 0; i < len(str);  {
			if string(str[i+1]) != " "{
				arr[s].PushLast(str[i+1])
			}	
			i +=4
			s +=1
		}
	} 

	return arr
}

