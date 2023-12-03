package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func final_write(strs []string) {
	// f, err := os.Create("answer.txt")

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	bigprint := 0
	for i := 0; i < len(strs); i++ {
		newS := string(strs[i])
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
	fmt.Println(output)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	final_write(output)
}
