package main

import(
	"fmt"
	"os"
	//part  "github.com/Mohd-Sayeedul-Hoda/day5/part1"
	part "github.com/Mohd-Sayeedul-Hoda/day5/part2"
)

func main(){
	file, err := os.Open(os.Args[1])
	defer file.Close()
	if err != nil{
		fmt.Fprintf(os.Stderr, "Can't open the provided file ")
		return 
	}
//	part.Part1(file)
	part.Part2(file)
}
