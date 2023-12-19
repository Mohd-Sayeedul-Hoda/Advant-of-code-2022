package part1

import(
	"fmt"
	"os"
	"bufio"
)


func Part1(){
	file, err := os.Open(os.Args[1])
	if err != nil{
		fmt.Printf("Cannot open the file")
		return
	}
	input := bufio.NewScanner(file)
	input.Scan()
	str := input.Text()
	fmt.Println(str)
	check(str)
}

func check(str string){
	for f := 0; f < len(str); f++{
		if str[f] != str[f+3] && str[f+1] != str[f+2] && str[f] != str[f+1] && str[f+3] != str[f+1] && str[f+2] != str[f+3] && str[f+2] != str[f]{
			fmt.Println(f+4)
			break
		}	
	}
}


