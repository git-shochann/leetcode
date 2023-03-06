package main

import (
	"bufio"
	"coding-test/coding-test/m3"
	"fmt"
	"os"
)

func main() {
	// fmt.Println(valid_anagram.ValidAnagram("dog", "god")) // trueを返す
	// fmt.Println(valid_anagram.ValidAnagram("cat", "god")) // falseを返す
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	fmt.Printf("text: %v\n", text)

	fmt.Printf("solution: %v\n", m3.Solution(text))

	//m3.OtherSolution()
}
