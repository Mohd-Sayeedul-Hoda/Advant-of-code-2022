package part1

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

type Stack struct{
	slice []string
}

func (s *Stack) Push(str string){
	s.slice = append(s.slice, str)
}

func (s *Stack) Peek()string{
	if len(s.slice) == 0{
		return ""
	}
	str := s.slice[len(s.slice)-1]
	return str
}

func (s *Stack) Pop()string{
	temp := s.Peek()
	s.slice = s.slice[0:len(s.slice)-1]
	return temp
}

func Part1(){
	file, err := os.Open(os.Args[1])
	if err != nil{
		fmt.Println("Please specify the right file")
		return
	}
	defer file.Close()
	input := bufio.NewScanner(file)
	dirMap := make(map[string]int)
	PerformTask(input, dirMap)
	// so we are going to make stack and perform operation on that thats my plan lets see what will happen
	fmt.Println()
	for key, val := range dirMap{
		fmt.Println(key, val)
	}
	CalculatingSumLess(dirMap)
}

func PerformTask(input *bufio.Scanner, dirMap map[string]int){
	stack := new(Stack)
	var mostRecentDir string
	for input.Scan(){
		str := input.Text()
		strArr := strings.Split(str, " ")
		if strArr[0] == "$"{
			if strArr[1] == "ls"{
				continue
			}else{
				if strArr[2] == ".."{
					temp := stack.Pop()
					mostRecentDir = stack.Peek()
					dirMap[mostRecentDir] += dirMap[temp]
				}else{
					mostRecentDir = path.Join(stack.Peek(), strArr[2])
					stack.Push(mostRecentDir)	
				}
			}	
		}else{
			if strArr[0] == "dir"{
				continue
			}else{
				val, err := strconv.Atoi(strArr[0])	 
				if err != nil{
					fmt.Println("Error while converting to integer")
					return
				}
				dirMap[mostRecentDir] += val

			}

		}
	}
	fmt.Println("running for loop")
	for true{
		temp := stack.Pop()
		if stack.Peek() == ""{
			break
		}
		mostRecentDir = stack.Peek()
		dirMap[mostRecentDir] += dirMap[temp]
		fmt.Printf("%s %s %d %d\n", mostRecentDir, temp, dirMap[mostRecentDir], dirMap[temp])
		fmt.Println()
	}
}

func CalculatingSumLess(dirMap map[string]int){
	ans := 0
	for _, val := range dirMap{
		if val <= 100000{
			ans += val
		}
	}
	fmt.Println("ans is ", ans)
}
