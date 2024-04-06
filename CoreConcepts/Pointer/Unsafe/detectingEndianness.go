/*
	This function finds out the endianness/ the order in which the bytes are stored in this system.
	Bigendian: Higher order bytes are stored at first.
	Littleendian: Lower order bytes are stored at first.
*/
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var x int
	x = 0x0112
	ptr := unsafe.Pointer(&x)
	if 0x01 == *(*byte)(ptr) {
		fmt.Println("BigEndian")
	} else if 0x12 == *(*byte)(ptr) {
		fmt.Println("LittleEndian")
	}
	fmt.Printf("0x12 as it's little endian in decimal is %d.\n", *(*byte)(ptr))
	fmt.Println(`unsafe.Pointer(&x) returns the pointer to the location where integer x is stored.
**************************************************************************************************
The integer 0x1232, this is written in hexadecimal(base 16)
0x1232 in decimal/base 10 is 0*16^3 + 1*16^2 + 1*16^1 + 2*16^0 = 274
1 byte = 8 bits
0000 0000
---- ----
0-15 0-15
4 bits store a number from 0 to 15, in hexa decimal we represent 0-9 as they are after that A-F
0,1,2,3,4,5,6,7,8,9,A,B,C,D,E,F
12 = 0xC
1 = 0x1
5 = 0x5
**************************************************************************************************
So when we convert *(*byte)(ptr) the pointer to byte pointer and get the value of the byte(i., 8 bits)
we get the 2 numbers two hexadecimal digits from the number.
Big endian:		  01 12  // upper byte is stored first
Little endian:    12 01  // lower byte is stored first
So, if the first byte has 0x01 then it's big endian else if it sotres 0x12 it's big endian.
`)
}
