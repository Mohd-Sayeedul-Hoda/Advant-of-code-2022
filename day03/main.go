package main

import (
	"fmt"
	"os"
	"bufio"
)

func main(){
	file, err := os.Open(os.Args[1])
	if err != nil{
		fmt.Fprintf(os.Stderr, "File cannot be located")
	}

	defer file.Close()

	input := bufio.NewScanner(file)
	input.Split(bufio.ScanLines)

	var lines []string

	for input.Scan(){
		lines = append(lines, input.Text())
	}
	ans := 0
	i := 0
	for i < len(lines){
		temp := commonBadge(lines[i:i+3])	
//		fmt.Println(temp)
		ans += temp
		i +=3
	}
	fmt.Println(ans)
}

func commonBadge(str []string) int{
	myMap := make(map[byte]int8)

	for i, valArr := range str{
		for _, val := range valArr{
			c := byte(val)
			if _, exist := myMap[c]; exist{
				if i == 0{
					continue
				}else if i == 1{
					if myMap[c] == 1{
						myMap[c] += 1
					}else{
						continue
					}	
				}else{
					if myMap[c] == 2{
						myMap[c]+=1
					}else{
						continue
					}
					if myMap[c] == 3{
						fmt.Printf("%c\n", c)
						return getPiority(int(c))
						}
					}
			}else{
				if i ==0 {
					myMap[c] = 1
				}else{
					continue
				}
			}
		}
	}

	return 1	
}

func getPiority(c int)int{
	val := 0
	if c >= 97{
		val += (c - 96)
	}else{
		val += (c - 38)	
	}
	return val
}
