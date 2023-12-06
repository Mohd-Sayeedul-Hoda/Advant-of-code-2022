package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	f, err := os.Open(os.Args[1])	
	if err != nil{
		fmt.Fprintf(os.Stderr, "cant'open the file\n")
		return
	}
	fmt.Println(gamePoint(f))
}

func gamePoint(f *os.File)int{
	input := bufio.NewScanner(f)
	ans := 0
	for input.Scan(){
		arr := strings.Split(input.Text(), " ")	
		ans += pointWin(arr[0], arr[1])
	}
	return ans
}

func pointWin(opponent string,player string)int{
	ret :=0
	switch player{
		case "X": 
		ret += 1
		if opponent == "A"{
				ret += 3
				return ret
		}else if opponent == "B"{
			ret += 0
			return ret
		}else{
			ret += 6
			return ret
		}
		case "Y":
		ret += 2
		if opponent == "A"{
			ret += 6
			return ret
		}else if opponent == "B"{
			ret += 3
			return ret
		}else{
			ret += 0
			return ret
		}
		case "Z":
			ret += 3
		if opponent == "A"{
				ret += 0
				return ret
		}else if opponent == "B"{
			ret += 6
			return ret
		}else{
			ret += 3
			return ret
		}
	}
	return 1
}


