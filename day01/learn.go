package main 

import (
	"fmt"
	"bufio"
	"os"
)

func main(){
	fmt.Println(os.Args[1])
	f, err := os.Open(os.Args[1])
	if err != nil{
		fmt.Fprintf(os.Stderr, "some thing wrong with the file: %s", err)
	}
	input := bufio.NewScanner(f)
	for input.Scan(){
		char := input.Text()
		fmt.Println(char)
	}
}
