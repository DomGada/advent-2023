// Starting with the starter code from day1

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func IndexOfSubstring(str, subStr string) []int {
	ret := make([]int, 0, 3)
	for i := 0; i+len(subStr) <= len(str); i++ {
		if str[i:i+len(subStr)] == subStr {
			fmt.Println("Found the substring!")
			ret = append(ret, i)
		}
	}
	return ret
}

func final_write(strs []string) {
	// f, err := os.Create("answer.txt")

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	bigprint := 0
	myNums := [10]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for i := 0; i < len(strs); i++ {
		newS := string(strs[i])
		for k := 0; k < len(myNums); k++ {
			index := IndexOfSubstring(newS, myNums[k])
			if len(index) > 0 {
				for num := 0; num < len(index); num++ {
					out := []rune(newS)
					out[index[num]+1] = rune(k + 48)
					newS = string(out)
				}
			}
		}
		fmt.Println(newS)
		myprint := ""
		for j := 0; j < len(newS); j++ {
			if unicode.IsDigit([]rune(newS)[j]) {
				myprint = myprint + string(newS[j])
			}

		}
		addValue := string(myprint[0]) + string(myprint[len(myprint)-1])
		add, err := strconv.Atoi(addValue)
		if err != nil {
			panic(err)
		}
		bigprint = bigprint + add

		// l, err := f.WriteString(myprint)
		// if i < len(strs)-1 {
		// 	f.WriteString("\n")
		// }
		// if err != nil {
		// 	fmt.Println(err)
		// 	f.Close()
		// 	return
		// }
		// fmt.Println(l, "bytes written")

	}
	// err = f.Close()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	fmt.Println(bigprint)

}

func main() {
	file, err := os.Open("input.txt")

	output := make([]string, 0, 3)
	if err != nil {
		fmt.Println("File reading error.", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		output = append(output, scanner.Text())

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	final_write(output)
}
