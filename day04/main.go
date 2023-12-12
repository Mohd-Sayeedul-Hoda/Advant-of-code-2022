package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	file, err := os.Open(os.Args[1])
	if err != nil{
		fmt.Fprintf(os.Stderr, "Can't open the file")
		return
	}

	defer file.Close()

	input := bufio.NewScanner(file)
	ans := 0
	i := 1 
	for input.Scan(){
		str := input.Text()
		strArr := strings.Split(str, ",")
		ans += checkOverlap(strArr, i)
		i++
	}
	fmt.Println(ans)
}

func checkOverlap(strArr []string, line int)int{
	str1 := strings.Split(strArr[0], "-")
	str2 := strings.Split(strArr[1], "-")
	int1 := stringToIntArray(str1)
	int2 := stringToIntArray(str2)
	if int1[0] >= int2[0] && int1[1] <= int2[1]{ 
		fmt.Println(line)
		fmt.Println(str1, str2)
		return 1
	}else if int2[0] >= int1[0] && int2[1] <= int1[1]{
		fmt.Println("condition 2")
		fmt.Println(line)
		fmt.Println(str1, str2)
		return 1
	}
	return 0
}
func stringToIntArray(str []string) ([]int) {
	var intArr []int
	for _, s := range str{
		val, err := strconv.Atoi(s)
		if err != nil{
			fmt.Fprintf(os.Stderr, "cannot convert to int")
		}
		intArr = append(intArr, val)
	}
	return intArr
}
