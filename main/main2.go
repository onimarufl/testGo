package main

import "fmt"

func main() {

	str := "abcabc"
	chars := []rune(str)
	charsArr := []string{}
	for i := 0; i < len(chars); i++ {
		char := string(chars[i])
		charsArr = append(charsArr, char)
	}
	// cSet = []string{"a", "b", "c"}
	repeat := true
	i := 0
	for repeat {
		repeat = false
		for i < len(charsArr)-2 {
			aa := []string{}
			aa = append(aa, charsArr[i])
			aa = append(aa, charsArr[i+1])

			// bb := list(cSet-set([str[i],str[i+1]]))

			if charsArr[i] != charsArr[i+1] {
				fmt.Println(aa)
				repeat = true
			}
			i++
		}
		fmt.Println(len(charsArr))

	}
}
