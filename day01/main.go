package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main(){
	var f, err = os.Open(os.Args[1]);
	if err != nil{
		fmt.Fprintf(os.Stderr, "File cannot be open: %v", err)
	}
	input := bufio.NewScanner(f)
	var max [3]int 
	var temp = 0
	for input.Scan(){
		char := input.Text();
		if char == ""{
			if max[0] < temp{
				temp1 := max[0]
				temp2 := max[1]
				max[0] = temp
				max[1] = temp1
				max[2] = temp2
				
			}else if max[1] < temp{
				temp1 := max [1]
				max[2] = temp1
				max[1] = temp
			}else if  max[2] < temp{
				max[2] = temp
			}
			temp = 0 
			continue
		}
		num, err  := strconv.Atoi(char)
		if err != nil{
			fmt.Fprintf(os.Stderr, "cannot convert string to int %v %v \n", char,  err)
		}
		temp += num
	}
	fmt.Println(max[0]+max[1]+max[2])
}

