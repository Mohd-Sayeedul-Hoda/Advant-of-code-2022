package part2

import(
	"fmt"
	"os"
	"bufio"
)


func Part2(){
	file, err := os.Open(os.Args[1])
	if err != nil{
		fmt.Printf("Cannot open the file")
		return
	}
	input := bufio.NewScanner(file)
	input.Scan()
	str := input.Text()
	check(str, 14)
}

func check(str string, x int){
	charCount := make(map[rune]int)
	distChar := ""

	for i, char := range str{
		charCount[char]++
		if charCount[char] == 1{
			distChar += string(char)
		}else{
			for i, val := range distChar{
				if string(val) == string(char){
					distChar = distChar[i+1:len(distChar)]
				}
			}
			distChar += string(char)
			charCount = make(map[rune]int)
			for _, char  := range distChar{
				charCount[char] = 1 
			}
		}
		if len(distChar) == 14 {
			fmt.Println(i+1)
			break
		}
		
	}

}


