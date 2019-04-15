package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Write here > ")
	input, _ := reader.ReadString('\n')
	fmt.Println(input)
}
