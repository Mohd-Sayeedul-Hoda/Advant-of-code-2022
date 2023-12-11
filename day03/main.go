package main

import(
	"fmt"
	"os"
	"bufio"
)

func main(){
	f, err := os.Open(os.Args[1])
	if err != nil{
		fmt.Fprintf(os.Stderr, "Can't open the file you provide\n")
	}
	fmt.Println(readFile(f))
}

func readFile(f *os.File)int{
	input := bufio.NewScanner(f)
	ans := 0
	for input.Scan(){
		str := input.Text()// If want i can add two more input.scan and two variable but i want to read 
		//file ans whole and split it with string.split into array of chunck and then want to work on it
		ans += commonItem(str)
	}
	return ans
}

func commonItem(str string)int{
	mid := int(len(str)/2)
	firstRack := str[:mid]
	secondRack := str[mid:]
	itemMap := make(map[byte]int8)
	repeatedItem := make([]int, 0)
	
	for _, v := range firstRack{
		c := byte(v)
		if _, exists := itemMap[c]; exists {
			continue
		}else{
			itemMap[c] = 1
		}
	}
	for _, v := range secondRack{
		c := byte(v)
		if _, exists := itemMap[c]; exists && itemMap[c] <= 1 {
			repeatedItem = append(repeatedItem, int(c))
			itemMap[c]++
		}else{
			continue
		}
	}
	return  getPiority(&repeatedItem)
}

func getPiority(item *[]int)int{
	val := 0
	for _, v := range *item{
		if v >= 97{
			val += (v - 96)
		}else{
			val += (v - 38)	
		}
	}
	return val
}
