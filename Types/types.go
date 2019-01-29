/* *****************************************************************************
 *  Name:      Arbaaz Khan
 *  Language:  go
 *
 *  Description:  In this program we have explored different types of datatypes
                  available in go.
 *
 *  % go run types.go
 *
***************************************************************************** */
package main

import "reflect"

func main() {
	/* Different datatypes available:
	   * bool
	   * Numeric Types
	       * int8, int16, int32, int64, int
	       * uint8, uint16, uint32, uint64, uint
	       * float32, float64
	       * complex64, complex128
	       * byte
	       * rune
	   * string
	*/
	//bool
	println("*******bool*******")
	var b bool //default value is false
	var c = true
	d := b && c //logical and
	e := b || c //logical or
	println("b=", b, "c=", c, "d=", d, "e=", e)
	//int
	println("*******int*******")
	var f = 5
	println("Type of f=", reflect.TypeOf(f).String())
	var g int8 = 6           //Range -128 to 127
	var h int16 = -123       //Range -32768 to 32767
	var i int32 = 5324       //Range -2147483648 to 2147483647
	var j int64 = -351678416 //Range -9223372036854775808 to 9223372036854775807
	println("g=", g, "h=", h, "i=", i, "j=", j)

	//unint
	println("*******unint*******")
	var k uint8 = 6          //Range 0 to 255
	var l uint16 = 123       //Range 0 to 65535
	var m uint32 = 5324      //Range 0 to 4294967295
	var n uint64 = 351678416 //Range 0 to 18446744073709551615
	println("k=", k, "l=", l, "m", m, "n=", n)

	//float
	println("*******float*******")
	var o float32 = 6.54     //Range 0 to 255
	var p float64 = 123.4989 //Range 0 to 65535
	println("o=", o, "p=", p)

	//complex
	println("*******complex*******")
	q1 := 6 + 7i
	q2 := 9 + 6i
	q3 := q1 + q2
	println("q1=", q1, "q2=", q2, "q3=", q3)

	//byte
	println("*******byte*******")
	println("byte is an alias of uint8")

	//rune
	println("*******rune*******")
	println("rune is an alias of int32")

	//string
	println("*******string*******")
	var r string
	r = "hi"
	println("r=", r)

}
