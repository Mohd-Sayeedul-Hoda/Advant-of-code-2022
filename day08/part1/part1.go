package part1

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Part1(){
	input, err := os.ReadFile(os.Args[1])
	if err != nil{
		log.Fatal("Can't read the file ")
	}
	treeNodes := DeclaringGrid(input)
	for _, treeNode := range treeNodes{
		fmt.Println(treeNode)
	}
	ans := FindVisible(treeNodes)
	fmt.Println(ans)
}

// I know this will not the most optimised solution but to some degree it will 
// I will try to optimize it more in future but for know have to deal with it
func FindVisible(treeNodes [][]int)int{
	ans := 0
	upSide := make([]int, len(treeNodes[0])) // it will be compare with j in tree nodes
	for i := range upSide{
		upSide[i] = -1             // Loop and assign -1 to each element
	}
	leftSide := make([]int, len(treeNodes)) // it will compare with i in tree nodes
	for i := range leftSide{
		leftSide[i] = -1             // Loop and assign -1 to each element
	}
	for i, trees := range treeNodes{
		for j, tree := range trees{
			if tree > upSide[j]{
				fmt.Println(tree, "from up")
				ans += 1
				upSide[j] = tree
				continue
			}else if tree > leftSide[i]{
				fmt.Println(tree, "from left")
				ans += 1
				leftSide[i] = tree
				continue
			}else{
				// here we have to find visibility from right side and downside 
				// Right side calculating here
				k := j+1 
				l := len(trees)-1
				if j == k || i == len(treeNodes) -1 {
					ans += 1
					continue
				}
				forCheck := false
				for k <= l {
					if trees[k] >= tree || trees[l] >= tree{
						forCheck = true
						break
					}
					k += 1
					l -= 1
				}
				if forCheck{
					// here we going to check down side of it
					m := i+1
					n := len(treeNodes)-1
					newCheck := false
					for m <= n {
						if treeNodes[m][j] >= tree || treeNodes[m][j] >= tree{
							newCheck = true
							break
						}
						m += 1
						n -= 1
					}
					if !newCheck{
						fmt.Println(tree, "from down")
						ans += 1
					}

				}else {
					fmt.Println(tree, "from right")
					ans += 1
					continue
				}
				
			}
		}
	}

	return ans
}

func DeclaringGrid(input []byte)[][]int{
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	tree := make([][]int, len(lines))
	for i, line :=  range lines{
		line = strings.TrimSpace(line)
		tree[i] = make([]int, len(line))
		for j, treeNode := range line{
			tree[i][j], _ = strconv.Atoi(string(treeNode))
		}
	}
	return tree
}
