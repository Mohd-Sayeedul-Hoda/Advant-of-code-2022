package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	file, err := os.Open(os.Args[1])
	if err != nil{
		fmt.Fprintf(os.Stderr, "Can't open the file")
		return
	}
	input := bufio.NewScanner(file)
	ans := 0
	for input.Scan(){
		str := input.Text()
		strArr := strings.Split(str, ",")
		ans += checkOverlap(strArr)
	}
	fmt.Println(ans)
}

func checkOverlap(strArr []string)int{
	str1 := strings.Split(strArr[0], "-")
	str2 := strings.Split(strArr[1], "-")
	if str1[0] >= str2[0] && str1[1] <= str2[1]{
		return 1
	}else if str2[0] >= str1[0] && str2[1] <= str1[1]{
		return 1
	}
	return 0
}

