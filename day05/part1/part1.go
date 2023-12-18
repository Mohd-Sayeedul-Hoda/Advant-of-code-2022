package part1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	return s.slice[len(s.slice)-1]
}

func (s *Stack)Pop()byte{
	var temp byte = s.Peek()
	s.slice = s.slice[0:len(s.slice)-1]
	return temp
}

func Part1(file *os.File){
	input := bufio.NewScanner(file)
	arr := MakeStack(input)
	for _, v := range arr{
		for _, val := range v.slice{
			fmt.Printf("%c", val)
		}
		fmt.Printf("\n")
	}
	Move(input, arr)
	fmt.Printf("\n")
	for _, v := range arr{
		for _, val := range v.slice{
			fmt.Printf("%c", val)
		}
		fmt.Printf("\n")
	}
}

func MakeStack(input *bufio.Scanner)[]*Stack{
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
		if string(str[1]) == "1"{
			return arr
		}
		s := 0 // for stack value
		for i := 0; i < len(str); i += 4  {
			if string(str[i+1]) != " "{
				arr[s].PushLast(str[i+1])
			}	
			s +=1
		}
	} 

	return arr
}

func Move(input *bufio.Scanner, arr []*Stack){
//	firstTime := true
	for input.Scan(){
		str := input.Text()
		if string(str) == ""{
			continue
		}
		howMuch, err := strconv.Atoi(string(str[5]))
		if err != nil{
			fmt.Println("cannot convert it into string")
			return
		}
		from, err := strconv.Atoi(string(str[12]))
		if err != nil{
			fmt.Println("cannot convert it into string")
			return
		}
		to, err := strconv.Atoi(string(str[17]))
		if err != nil{
			fmt.Println("cannot convert it into string")
			return
		}
		for i := 0; i < howMuch;  i++{
			temp := arr[from-1].Pop()	
			arr[to-1].Push(temp)
		}
	}
}
