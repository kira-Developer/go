package main 

import "fmt"
/*

Boolean types	They are boolean types and consists of the two predefined constants: (a) true (b) false
Numeric types	They are again arithmetic types and they represents a) integer types or b) floating point values throughout the program.
String types	A string type represents the set of string values. Its value is a sequence of bytes. Strings are immutable types that is once created, it is not possible to change the contents of a string. The predeclared string type is string.
Derived types	They include (a) Pointer types, (b) Array types, (c) Structure types, (d) Union types and (e) Function types f) Slice types g) Interface types h) Map types i) Channel Types

uint8 	Unsigned 8-bit integers (0 to 255)
uint16 	Unsigned 16-bit integers (0 to 65535)
uint32 	Unsigned 32-bit integers (0 to 4294967295)
uint64	Unsigned 64-bit integers (0 to 18446744073709551615)

int8	Signed 8-bit integers (-128 to 127)
int16	Signed 16-bit integers (-32768 to 32767)
int32	Signed 32-bit integers (-2147483648 to 2147483647)
int64	Signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32 	IEEE-754 32-bit floating-point numbers
float64 	IEEE-754 64-bit floating-point numbers
complex64	Complex numbers with float32 real and imaginary parts
complex128 	Complex numbers with float64 real and imaginary parts

byte 		same as uint8
rune		same as uint8
uint		32 or 64 bits
int			same size as uint
uintptr		an unsigned integer to store the uninterpreted bits of a pointer value
*/
func main(){
 var intnumber int8 = -6
 var uintnumber uint8 = 6
 var floatnumber float32 = -6.7

 fmt.Println(intnumber)
 fmt.Println(uintnumber)
 fmt.Println(floatnumber)
}