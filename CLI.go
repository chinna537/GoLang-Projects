package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the text")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)

	// this is single line comment

}
