package golearn

import (
	"fmt"
	"strconv"
)

// BIG CONCEPT
// lowercase symbols (variables, structs, function names)
// exist on the PACKAGE-LEVEL, meaning they are not exported
// Uppercase symbols (variables, structs, function names)
// exist on the GLOBAL-LEVEL, meaning they are exported (included
// if another package imports them)
// NOTE: variables only will be exported if they are declared
// at the package level, as shown:
var poop string = "poop" // NOT exported
// Poop needs a comment since all exported symbols need comments
var Poop string = "p00p" // exported

// global var can only be declared like this
var i int = 69

// Hello returns our hello world string
func Hello() string {
	return "Hello World!"
}

// Declarations shows the ways to declare variables in Go
func Declarations() {
	fmt.Println("Showing basic declarations in Go...")
	// var NAME TYPE
	var a int
	a = 1
	fmt.Println("a is", a)

	// var NAME TYPE = VALUE
	var b int = 22
	fmt.Println("b is", b)

	// NAME := VALUE
	// type inference used here
	c := a + b
	fmt.Printf("c is %d, %T\n", c, c)

	//we can do group declarations like this
	var (
		AA int = 1
		BB int = 2
		CC int = 3
		i  int = 70 // shadowing rule in Go is that inner-most scope always wins
	)
	DD := AA + BB + CC + i
	fmt.Println("DD is: ", DD)
	// if the grouping makes some sense
}

// Conversions shows some basic concepts of converting between types
// in Go's strong typed system
func Conversions() {
	fmt.Println("Showing basic conversions in Go's strong type system...")
	// basic float types are 32 and 64 bits
	j := float32(i)
	k := float64(i)

	fmt.Printf("i: %d, j: %f, k: %f\n", i, j, k)

	// string conversions
	var s string
	s = string(i)
	fmt.Printf("string(%d) = %v\n", i, s)

	s = strconv.Itoa(i)
	fmt.Printf("strconv.Itoa(%d) = %v\n", i, s)
}

// Primitives details the basic types Go provides
func Primitives() {
	fmt.Println("Showing the basic types in Go...")
	// boolean
	var n bool = 2 == 2
	fmt.Printf("var = %v, %T\n", n, n)

	// numerics
	var (
		// integral
		signedInt     int    = 42
		unsignedInt   uint   = 42
		signedInt8    int8   = 42
		byteAlias     byte   = 42
		signedInt16   int16  = 42
		signedInt32   int32  = 42
		signedInt64   int64  = 42
		unsignedInt8  uint8  = 42
		unsignedInt16 uint16 = 42
		unsignedInt32 uint32 = 42
		unsignedInt64 uint64 = 42

		// fp gang
		f32     float32    = 42e17
		f64     float64    = 42e18
		comp64  complex64  = 1 + 3i
		comp128 complex128 = 1 + 4i
	)
	fmt.Printf("var = %v, %T\n", signedInt, signedInt)
	fmt.Printf("var = %v, %T\n", unsignedInt, unsignedInt)
	fmt.Printf("var = %v, %T\n", signedInt8, signedInt8)
	fmt.Printf("var = %v, %T\n", unsignedInt8, unsignedInt8)
	fmt.Printf("byte alias = %v, %T\n", byteAlias, byteAlias)
	fmt.Printf("var = %v, %T\n", signedInt16, signedInt16)
	fmt.Printf("var = %v, %T\n", unsignedInt16, unsignedInt16)
	fmt.Printf("var = %v, %T\n", signedInt32, signedInt32)
	fmt.Printf("var = %v, %T\n", unsignedInt32, unsignedInt32)
	fmt.Printf("var = %v, %T\n", signedInt64, signedInt64)
	fmt.Printf("var = %v, %T\n", unsignedInt64, unsignedInt64)
	fmt.Printf("var = %v, %T\n", f32, f32)
	fmt.Printf("var = %v, %T\n", f64, f64)
	fmt.Printf("var = %v, %T\n", comp64, comp64)
	fmt.Printf("var = %v, %T\n", comp128, comp128)
	fmt.Printf("real(comp128) = %v, %T\n", real(comp128), real(comp128))
	fmt.Printf("imag(comp128) = %v, %T\n", imag(comp128), imag(comp128))

	// numeric operations
	fmt.Println("Basic Numeric Type Operations:")
	a := 10
	b := 3
	fmt.Printf("(%d + %d) = %d\n", a, b, a+b)
	fmt.Printf("(%d - %d) = %d\n", a, b, a-b)
	fmt.Printf("(%d * %d) = %d\n", a, b, a*b)
	fmt.Printf("(%d / %d) = %d\n", a, b, a/b)
	fmt.Printf("(%d %% %d) = %d\n", a, b, a%b)

	fmt.Println("Basic Bit Operations:")
	fmt.Printf("(%d & %d) = %d\n", a, b, a&b)   // AND
	fmt.Printf("(%d | %d) = %d\n", a, b, a|b)   // OR
	fmt.Printf("(%d ^ %d) = %d\n", a, b, a^b)   // XOR
	fmt.Printf("(%d &^ %d) = %d\n", a, b, a&^b) // NAND
	fmt.Printf("(%d << %d) = %d\n", a, b, a<<b) // shift left
	fmt.Printf("(%d >> %d) = %d\n", a, b, a>>b) // shift right

	// Text Types
	fmt.Println("Basic Text Types:")
	var (
		// string literals are "string" double quotes
		// rune literals are 'rune' single quotes
		s1        string = "this is a string"
		and       string = " and "
		s2        string = "this is another string"
		byteSlice        = []byte(s1) // slice of byte
		r         rune   = 'a'        // runes are int32
	)
	fmt.Printf("var = %v, %T\n", s1, s1)
	fmt.Printf("var = %v, %T\n", byteSlice, byteSlice)
	fmt.Printf("rune = %v, %T\n", r, r)

	// Text Operations
	fmt.Printf("(s1 + and + s2) = %s\n", s1+and+s2) // string concatenation
}

// Constants covers:
// naming convention
// typed constants
// untyped constants
// enumerated constants
// enumeration expressions
func Constants() {
	fmt.Println("Showing off constants in Go...")
	// constants preceded by "const" keyword
	// do not name constants as MYCONST in Go
	// because the capital first letter will
	// mean that all constants are exported,
	// which may not be your intention

	// coming from C++ "const" in Go = "constexpr" in C++
	// it is a compile-time constant, where the symbol
	// is propogated throughout the code as a literal
	// compile-time evaluation of large expressions
	// (which sort of is possible in C++) may not be
	// a big thing in Go.  Will have to investigate more
	const myConst int = 32 // not exported
	const MyConst int = 32 // exported
	const MYCONST int = 32 // exported

	// typed constants have their type included in the declaration
	const a int = 33 // include type

	// untyped constants have their type not included in declaratoin
	const b = 34

	// enumerated constants
	// use the "iota" keyword, which iterates itself
	// everytime a const is assigned with it
	const (
		c = iota
		d
		e
	)

	// a call to iota is limited within the scope of a
	// single enumerated expression
	const (
		f = iota
		g
	)
	fmt.Printf("const c = iota -> %v, %T\n", c, c)
	fmt.Printf("const d = iota -> %v, %T\n", d, d)
	fmt.Printf("const e = iota -> %v, %T\n", e, e)

	fmt.Printf("const f = iota -> %v, %T\n", f, f)
	fmt.Printf("const g = iota -> %v, %T\n", g, g)

	// we can pack multiple bit flags into a single byte
	const (
		isAdmin = 1 << iota
		isHeadquarters
		canSeeFinancials

		canSeeAfrica
		canSeeAsia
		canSeeEurope
		canSeeNorthAmerica
		canSeeSouthAmericca
	)

	fmt.Println("Bit flag packed byte:")
	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("\n%b\n", roles)
}
