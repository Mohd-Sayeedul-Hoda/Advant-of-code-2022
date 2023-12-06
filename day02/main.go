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
		if opponent == "A"{
				ret += 3
				return ret
		}else if opponent == "B"{
			ret += 1
			return ret
		}else{
			ret += 2
			return ret
		}
		case "Y":
		if opponent == "A"{
			ret += 1+3
			return ret
		}else if opponent == "B"{
			ret += 2+3
			return ret
		}else{
			ret += 3+3
			return ret
		}
		case "Z":
		if opponent == "A"{
				ret += 2+6
				return ret
		}else if opponent == "B"{
			ret += 3+6
			return ret
		}else{
			ret += 1+6
			return ret
		}
	}
	return 1
}


