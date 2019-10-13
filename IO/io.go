package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// Reading string input from std input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a string>> ")
	inputStr, err := reader.ReadString('\n') // Reads a string till enter is clicked
	if err != nil {
		log.Print(err)
	}
	fmt.Printf("String read = %s\n", inputStr)

	//NOTE: ( It reads only one character what ever may be the input, the first character is read.)
	fmt.Print("Enter a Character/Rune>> ")
	r, s, err := reader.ReadRune() // Reads a Character/Rune till enter is clicked, It returns the rune, size in bytes, err
	if err != nil {
		log.Print(err)
	}
	fmt.Printf("Character/Rune read = %c, Size = %d\n", r, s)

	_, _, _ = reader.ReadLine() //To consume the Enter key pressed for the above code.

	fmt.Print("Enter a Line>> ")
	b, isPrefix, err := reader.ReadLine() // Reads a Line till enter is clicked, It returns the bytes, isPrefix(boolean set to true if long line is read, in this case the first line is returned), err
	if err != nil {
		log.Print(err)
	} else if isPrefix {
		fmt.Printf("Line read read = %b, isPrefix = %t\n", b, isPrefix)
	} else {
		fmt.Printf("Line read = %s, Size = %t\n", b, isPrefix)
	}

	// ReadBytes / ReadString is a better option over ReadSlice(NOTE: Avoid using it.)
	// fmt.Print("Enter a slice>> ")
	// inputSlice, err := reader.ReadSlice('\n') // Reads a string till the first occurence of delimiter is read, it stores each byte in a byte array.
	// if err != nil {
	// 	log.Print(err)
	// }
	// for _, v := range inputSlice {
	// 	fmt.Printf("Slice read = %c\n", v)
	// }
}
