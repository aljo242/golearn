package golearn

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
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
func Declarations() string {
	fmt.Println("\nShowing basic declarations in Go...")
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
	return "Declarations"
}

// Conversions shows some basic concepts of converting between types
// in Go's strong typed system
func Conversions() string {
	fmt.Println("\nShowing basic conversions in Go's strong type system...")
	// basic float types are 32 and 64 bits
	j := float32(i)
	k := float64(i)

	fmt.Printf("i: %d, j: %f, k: %f\n", i, j, k)

	// string conversions
	var s string
	//s = string(i)
	//fmt.Printf("string(%d) = %v\n", i, s)

	s = strconv.Itoa(i)
	fmt.Printf("strconv.Itoa(%d) = %v\n", i, s)
	return "Conversions"
}

// Primitives details the basic types Go provides
func Primitives() string {
	fmt.Println("\nShowing the basic types in Go...")
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
	return "Primitives"
}

// Constants covers:
// naming convention
// typed constants
// untyped constants
// enumerated constants
// enumeration expressions
func Constants() string {
	fmt.Println("\nShowing off constants in Go...")
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
	fmt.Printf("\t%b\n", roles)
	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin)            //000001 & 100101 = 000001
	fmt.Printf("Is HQ? %v\n", isHeadquarters&roles == isHeadquarters) // 000010 & 100101 = 000000

	// IN SUMMARY
	// constants are immutable, but CAN be shadowed
	// constant values must be replaced by the compiler
	//		AT compile time
	//		therefore, must be calculatable at compile time
	// are named just like variables:
	//	exported 		: PascalCase
	// 	non-exported 	: camelCase
	// typed constants work like immutable variables
	//		interoperate only with own type
	// untyped constants behave like literals in source
	//		can interoperate with similar types
	// 		therefore, are more like #DEFINE in C/C++
	// Enumerated Constants
	// 	special symbol, "iota", allows for related constants to be created
	// 	iota starts at 0 in each unique const block, then increments by 1
	// 	watch out for defining constant values at 0, since variables
	//		are generally 0-initialized in Go
	// Enumerated Operations
	// operations that can be determined at compile time are allowed
	//	use the iota construct paired with operations to create related constants
	//	arithmetic, bitwise operations, bitshifting
	return "Constants"
}

// ArraysAndSlices first details arrays, which are the basis
// for slices, then slices, which allow for dynamic views
// of allocated memory
func ArraysAndSlices() string {
	fmt.Println("\nShowing Array and Slices Basics in Go...")
	// Arrays are declated using:
	// NAME := [SIZE]TYPE{initializer_list}
	// where size is a compile time constant
	// var NAME [SIZE]TYPE
	// NAME := [...]TYPE{intializer_list}
	// ... is ok, because size can just be inferred from the initializer list length
	// we dont have NAME := []TYPE because that is the syntax for SLICES

	// arrays in Go ARE contiguous in memory

	// NAME := [SIZE]TYPE{initializer_list}
	grades := [3]int{93, 45, 59}
	fmt.Println(grades) // can print arrays like this
	for index, grade := range grades {
		fmt.Println(index, grade) // or iterate using range construct
	}

	// NAME := [...]TYPE{initializer_list}
	grades2 := [...]int{93, 45, 59}
	fmt.Println(grades2) // can print arrays like this
	for index, grade := range grades2 {
		fmt.Println(index, grade) // or iterate using range construct
	}

	// var NAME [SIZE]TYPE
	var grades3 [3]int
	grades3 = grades2    // array assigmment here uses a copy
	fmt.Println(grades3) // can print arrays like this
	for index, grade := range grades3 {
		fmt.Println(index, grade) // or iterate using range construct
	}

	// ARRAY ASSIGNMENT IN GO IS ACTUALLY ALWAYS A COPY!!!!
	// WE CAN DO STUFF LIKE "POINTER" assignment in Go
	// when we use SLICES
	// lets prove this
	// grades3 is assigned from grades2, but we know this is a cop
	// so, if we modify grades3[2], only grades3 is modified, not grades2
	grades2[2] = 12
	fmt.Println("Arrays are copied in go, so modifying a copied array will not modify the orginal array:")
	fmt.Println(grades2, grades3)

	// if we use this syntax, we are taking a reference to the data
	// so modifications made on grades3p will affect the underlying grades
	grades3p := &grades2
	grades3p[2] = 12
	fmt.Println("Arrays can be taken by reference with &, and then will  modify the orginal array:")
	fmt.Println(grades2, grades3p)

	if len(grades) == len(grades2) && len(grades) == len(grades3) {
		fmt.Println("Length of arrays are all the same!")
	}

	// multi dim arrays
	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 0, 0}
	identityMatrix[1] = [3]int{0, 1, 0}
	identityMatrix[2] = [3]int{0, 0, 1}
	fmt.Println(identityMatrix)

	// SLICES
	// slices are projections onto an underlying array
	// along with a len() property, they also have a cap()
	// 		len : length of the data SEEN by the slice
	//		cap : length of the underlying array
	// as they are VIEWs of arrays
	// slices are like RANGES concepts in std::ranges and D-ranges
	// ARE THEY LAZILY EVALUATED?

	// main declaration is:
	// NAME := []TYPE{initializer_list}
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// since slices are "views" this assignment does not copy
	// both slices view the same underlying array
	slice2 := slice
	fmt.Println(slice, "Length:", len(slice), "Capacity:", cap(slice))

	fmt.Println("Modifying copied slice of original slice...")
	slice2[2] = 4
	fmt.Println(slice, "Length:", len(slice), "Capacity:", cap(slice))

	// other slice declarations
	a := slice[:]   // slice of all elements
	b := slice[3:]  // slice of index 3 and up			(element 3 to 9)
	c := slice[:6]  // slice up to index 6				(element 1 to 6)
	d := slice[3:6] // slice from index 3 up to index 6	(element 4 to 6)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	// slices can also come from arrays
	// this makes plenty of sense, as they are by definition
	// views of arrays
	// note that making a slice from an array will mean the length and capacity
	// should be the same, since the length of an ARRAY is equivalent to its capacity
	f := [3]int{3, 2, 1}
	fmt.Println("Making a slice from an array...")
	slicef := f[:]
	fmt.Println("Slice:", slicef, "len:", len(slicef), "cap:", cap(slicef))

	// one more way to make a slice
	// is to use the builtin "make" functionality
	// SLICE := make([]TYPE, LENGTH, CAPACITY)
	slice = make([]int, 3, 100)
	fmt.Println("Making a slice using make([]int, 3)")
	fmt.Println("Slice:", slice, "len:", len(slice), "cap:", cap(slice))
	slice = append(slice, 4)
	fmt.Println("Slice:", slice, "len:", len(slice), "cap:", cap(slice))
	slice = append(slice, 5, 6, 7, 8, 89, 190, 4)
	fmt.Println("Slice:", slice, "len:", len(slice), "cap:", cap(slice))

	// concatenating slices
	// the syntax:
	//		SLICE_A = append(SLICE_B, SLICE_C...)
	// 		decomposes SLICE_C into a literals list which can be accepted
	//		by the append() function
	slice = append(slice, slicef...)
	fmt.Println("Slice:", slice, "len:", len(slice), "cap:", cap(slice))

	// using slices like a stack
	// append() is basically push()
	// b = a[1:] shifts the slice, "popping" the first element essentially
	stack := make([]int, 3, 10)
	stack[0] = 1
	stack[1] = 2
	stack[2] = 3
	stack = append(stack, 4)     // push
	stack = stack[1:]            // pop_front
	stack = stack[:len(stack)-1] // pop_back
	stack = append(stack, 5)     // push
	stack = append(stack, 5)     // push

	// what about removing an element from the middle?
	stack = append(stack[:2], stack[3:]...)
	fmt.Println("Slice:", stack, "len:", len(stack), "cap:", cap(stack))

	// SUMMARY
	// Arrays are contigiuous collections of items of the same type
	//		have fixed size (at compile time)
	// Slices are "views" of arrays
	//		are backed by a real contiguous array somewhere in memory
	//		can be thought of as ranges or vectors since they have dynamic size
	//		can use the make([]TYPE, LENGTH, CAPACITY) fuction to create a slice
	//		with more capacity than length so we dont perform too many copies
	//			same idea as creating std::vector<T>
	//		append() can add items, but watch out for unncessary copies
	//		if you keep appending items
	//		slice copies all will refer to the same underlying array
	// 		since they are just VIEWs of a real place in memory
	return "ArraysAndSlices"
}

// Doctor is a basic type containing a Doctor Who doctor number, actor name, and slice of companion names
type Doctor struct {
	// Number is the number indicating which doctor it is
	Number     int
	actorName  string
	companions []string
}

// structs are just a collection of fields ^^^^
// as with anything in Go,	PascalCase = export
//							camelCase = no export

// Instead of typical OOP inheritence, Go uses a model called Composition

// Animal is a basic base struct for animal
type Animal struct {
	Name   string `reqired max:"100"`
	Origin string
}

// we can use tags in fields that give some kind of property to
// users of the struct
// basically a rule that has to be followed for a struct field

// Bird has added parameters for speed and ability to fly
type Bird struct {
	Animal // this struct is EMBEDDED into the Bird struct
	// it is not named, so now Bird has two additional fields
	// Name string
	// Origin string
	// that it got from being composed of Animal
	SpeedKPH float32
	CanFly   bool
}

// MapsAndStructs details other basic container primitives in Go
func MapsAndStructs() string {
	fmt.Println("\nShowing Maps and Structs Basics in Go...")

	// Maps behave the way they do in any language
	// syntax
	//		NAME := map[KEY]VALUE
	//		to be a KEY, a type must have equivalency checking
	//		because an equivalency check is performed to see if a key
	//		is in the map

	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	statePopulations["Georgia"] = 10310371 // add a new element

	fmt.Println(statePopulations)

	// can also declare maps using the "make" syntax
	otherMap := make(map[string]int)

	// we can read and write to maps by using their key like an array index
	// this one below CREATES a new key
	otherMap["key"] = 1          // write
	fmt.Println(otherMap["key"]) // read

	// note that maps do not have some kind of ordering
	// if you modify a map and then print it,
	//		the ordering might just be some random shit

	// we can delete items from maps too
	fmt.Println(statePopulations)
	delete(statePopulations, "Georgia") // deletes an item from a map
	fmt.Println(statePopulations)

	// if you query a key that does not exist in a map
	// the return will be the zero-init value for the value type
	// any query to a map also returns a second return value, ok bool
	// which you can use to check if the value was actually in the map
	pop, ok := statePopulations["Georgia"]
	if !ok {
		fmt.Println("Key not in map, returned value is:", pop)
	}

	// len(map) returns the number of elements in a map
	fmt.Println("Map length:", len(statePopulations))

	// maps are reference types, so modifications made to copies
	// will modify the original
	sp := statePopulations
	delete(sp, "Ohio")
	fmt.Println("Map length after delete on a copy:", len(statePopulations))

	// example of initalizing a struct with named fields
	// NAME = STRUCT {
	//	fieldA : value
	//	fieldB : value
	// 	...
	//}

	aDoctor := Doctor{
		Number:    3,
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}
	fmt.Println(aDoctor)

	// can access struct members using the . syntax
	if aDoctor.Number == 3 {
		fmt.Println("Doctor is number", aDoctor.Number)
	}

	// anonymous structs can be created locally and do not have type names
	anon := struct{ name string }{name: "Bob"}
	// structs are VALUE TYPES, so they are value COPIED
	anotherAnon := anon
	anotherAnon.name = "Joe"

	// can use references to copy by refence
	pAnon := &anon
	pAnon.name = "Jimmy"
	fmt.Println("OG:", anon)
	fmt.Println("Value Copy Modified", anotherAnon)
	fmt.Println("Reference Modified:", pAnon)

	// creating a struct with the Composed struct, Bird
	// which is has the Animal struct Embedded in it

	// When using composed types, this syntax is kind of nice
	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false
	fmt.Println(b)

	// because when using the named initializer syntax, its like this:
	// therefor you need to kind of know more about the layout of the struct
	c := Bird{
		Animal:   Animal{Name: "Emu", Origin: "Australia"},
		SpeedKPH: 48,
		CanFly:   false,
	}
	fmt.Println(c)

	// using reflection in Go
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)

	// SUMMARY
	// Maps are collections of value types accessed by keys
	// created by literal syntax or the make() function
	// check for presence using value, ok = map[key] syntax
	// 		ok is a bool
	// maps themselves are refrence types, so copies will modify
	// the underlying values

	// Structs collect all kinds of different types
	// can have names structs or anonymous structs using the same basic syntax
	// structs themselves are value types, so copies will not modify
	// be weary of costly copies with large structs

	// structs do not have typical OOP inheritence
	// use a composition technique called embedding
	// where structs are embedded into structs,
	// allowing sub-struct fields to be accessed like any other named field

	// Can use the "reflect" library
	// to query type, field, and tag info
	// can use these for validation framework
	return "MapsAndStructs"
}

func returnTrue() bool {
	fmt.Println("TRUE")
	return true
}

// ControlFlow details common control flow in Go (if, switch)
func ControlFlow() string {
	fmt.Println("\nShowing Control Flow Basics in Go...")

	// a lot of IDIOMATIC GO uses initializers within if statements
	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}

	// here we do
	// if initialize; boolean {}

	// this is nice, because below pop and ok
	// are only local to the if statement, just like (for i = 0; ...)
	if pop, ok := statePopulations["Florida"]; ok {
		fmt.Println(pop)
	}

	// remember that if there are multiple conditionals
	// ORed together, they are executed right-to-left
	num1 := -5
	num2 := 105
	num := num1

	// both of these will return true
	// but returnTrue() will only execute for the 2nd if stmt
	if num < 5 || returnTrue() || num > 105 {
		fmt.Println("Multi-statement is true")
	}

	num = num2
	if num < 5 || returnTrue() || num > 105 {
		fmt.Println("Multi-statement is true")
	}

	// IDIOMATIC Go uses switch statements instead of
	// large chains of if -> else-if -> else-if

	// unlike C/C++, switch statements have the "break;"
	// built into them, so no need to include

	// overlapping cases are NOT allowed
	switch i := 2 + 3; i { // can use initializers just like if
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3, 4, 5: // can have multiple tests as a comma-separated list
		fmt.Println("three, four, five")
		fallthrough
		// since Go has "break" implied in switch statements,
		// to get fallthrough behavior (which is the default in C-likes)
		// you need to add a fallthorough.  This means that the cases
		// of 3, 4, 5 and 20 are combined with fallthrough,
		// but at case 20:, the break is still there
		// so basically its the converse of C-like style
	case 20:
		fmt.Println("also maybe twenty (from fallthrough)")
		if i == 5 {
			break //  we can still insert breaks so we can skip stuff
			// perhaps an error occurs and we want to break to resolve
		}
		fmt.Println("PLEASE DONT PRINT THIS")
	default:
		fmt.Println("default")
	}

	// another unique switch syntax does not use a tag
	// since we are not switching a specific value
	// we can take any value that is in our CONTEXT
	// and use it in conditionals
	// this is where switch-case becomes the defacto if -> else if -> else
	// in Go
	i = i * 3
	switch {
	case i <= 10:
		fmt.Println("LEQ 20")
	case i <= 20:
		fmt.Println("LEQ 20")
	default:
		fmt.Println("Greater than 20")
	}

	// TYPE SWITCHING
	// interface type can be assigned to anything in Go,
	// so we may need to do type switching to find out what
	// it is at runtime

	// TYPE = NAME.(type) syntax
	// pulls the type parameter
	var j interface{} = 1
	switch j.(type) {
	case int:
		fmt.Println("j is an int")
	case float32, float64:
		fmt.Println("j is a float")
	case string:
		fmt.Println("j is a string")
	default:
		fmt.Println("j is another type")
	}
	return "ControlFlow"
}

// Loops details common loop structures in Go (ONLY for)
func Loops() string {
	fmt.Println("\nShowing Loop Basics in Go...")

	// All looping statements in Go use for
	fmt.Println("All looping statements in Go use for")

	// simple loops
	// classic syntax
	// for NAME := VALUE; CONDITION; ITERATION {
	//		body
	// }

	// note Go does not have "++i", only "i++"
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// for loop with multiple values initialized
	// we can use the syntax of
	// VAR1, VAR2 := VALUE1, VALUE2  	initialization
	// VAR1, VAR2 = VALUE1, VALUE2		assigmnent
	for i, j := 0, 5; i < 5; i, j = i+1, j-1 {
		fmt.Println(i, j)
	}

	// using an alread declared variable
	// here, i is scoped to the main function
	i := 0
	for ; i < 5; i++ {
		fmt.Println(i)
	}

	// removing the iteration statement
	// we now just have the Go equivalent of a
	// while loop
	for i == 5 {
		fmt.Println(i)
		i = 6
	}

	// infinite while loop
	for {
		fmt.Println("inside infinite loop, break me out!")
		break // break just exits the entire loop
	}

	// continue statements
	// just skip directly the next loop interation
	for i = 0; i < 10; i++ {
		if i%2 == 0 { // check if even
			continue
		}
		fmt.Println(i)
	}

	// can use labels to label the loop we want to break from in Go
OuterLoop:
	for i = 0; i < 3; i++ {
	InnerLoop:
		for j := 1; j < 3; j++ {
			ij := i * j
			fmt.Println(ij)
			if ij >= 3 {
				fmt.Println("Breaking from Outer Loop")
				break OuterLoop
			} else if ij == 45 {
				fmt.Println("Breaking from Inner Loop")
				break InnerLoop
			}
		}
	}

	// for initializer; test; incrementer {}
	// for test {}
	// for

	return "Loops"
}

func deferredGuy1() {
	fmt.Println("1. I was deferred at the beginning of DeferPanicRecover()")
}

func deferredGuy2() {
	fmt.Println("2. I was deferred at the beginning of DeferPanicRecover()")
}

// this is an interesting trick where we are passing an invocation
// of an anonymous function (A LAMBDA)
// it does a recover() call, which is basicall what catches a thrown panic
// so, thinking about it more, the semantics are pretty similar
// to try-catch with exceptions

func panicker() {
	fmt.Println("About to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error in Recover:", err)
		}
	}() // () here is the actual invocation of the func
	panic("Please Recover Me")
	fmt.Println("End of panicker")
}

// DeferPanicRecover shows some advanced control flow constructs in Go
func DeferPanicRecover() string {
	fmt.Println("\nShowing Defer, Panic, Recover Basics in Go...")

	// DEFER
	// deferred functions execute when the context it is called inside
	// returns to the context which called it
	// meaning, when this function returns, the deferred
	// functions are then executed
	defer deferredGuy1()
	defer deferredGuy2()
	// LIFO ordering
	// so think of the deferred functions as being pushed onto a stack
	// think also of closing resources in the opposite order you opened them
	// like using a deque in C++ to manage initialized items

	// lets use a real example to show some of this capability
	// the most common pattern is:
	// open resource with some call:
	//		res, err := GetResource()
	//		if err != nil {
	//			// we in trouble
	//		}
	//		defer res.Close()

	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close() // we will always close this resource no matter
	// what throws us from this function
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", robots)

	// one more thing to note is that if an argument is passed to a deferred
	// function, the argument that is used is the one seen AT DEFERED CALL
	// this means that if the argument changes later, the deferred function
	// does not see it
	// this in effect means that these arguments are put onto some kind of
	// stack by the Go runtime when the defer call is made

	a := "i will be printed"
	defer fmt.Println(a)
	a = "i will not be printed"

	// PANIC
	// Go does not have exceptions because the idea
	// is that most things that exceptions are thrown on in something like
	// C and C++ are not actually "exceptional".  We should just be able to
	// handle errors

	// if something really bad does happen, we can throw PANIC
	// a panic will throw a runtime error and abort
	// and also print a stack trace for some debug info
	// if you for example divide by 0, the Go runtime will
	// throw a panic, and there are other "destructive states"
	// that could be reached where a panic would be thrown
	if a == "i will be printed" {
		panic("WE ARE IN TROUBLE") // luckily this won't be triggered :)
	}

	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("Hello Go!"))
	//})
	//err = http.ListenAndServe(":8080", nil)
	//if err != nil {
	//	panic(err.Error())
	//}

	// this function below will call a panic,
	// but also has a deferred recover()
	// execution of panicker() stops when the panic is encountered,
	// so then the deferred statement is run
	// this will then recover, so the panic will not propogate
	// all the way up into this calling function
	// to this function, execution continues normally
	fmt.Println("start")
	panicker()
	fmt.Println("end")

	// we basically can only use recover() in a deferred context
	// if we want to recover inside (when leaving) the throwing function

	return "DeferPanicRecover"
}

type basicStruct struct {
	foo int
	bar int
}

// Pointers shows how pointers work in go... wow these are getting worse and worse
func Pointers() string {
	fmt.Println("\nShowing Pointers Basics in Go...")

	// pointers are basically the same to C-Like languages
	a := 42
	b := &a         // b is a pointer to the address of a
	var c int = 42  // equivalent to above
	var d *int = &c // equivalent to above
	fmt.Println(c, d)
	fmt.Println(a)

	*b = 21        // "assign value pointed to by b to 21"
	fmt.Println(a) // a will be  modified since b pointed to a

	// what about pointer arithmetic?
	// CANT DO IT...
	// 		unless we import the "unsafe" package :)
	arr := [3]int{1, 2, 3}
	b = &arr[0]
	d = &arr[1]
	fmt.Printf("%v %p %p\n", arr, b, d) // %p is pointer

	var sp *basicStruct
	sp = &basicStruct{foo: 1, bar: 2}
	sp = new(basicStruct) // cannot use the init list syntax
	// will be created with default values
	fmt.Println(sp)
	// to do assignment of fields from a pointer to a struct need to use this
	// (*NAME).field = VALUE syntax
	// because the () operation takes precedence over the . operator
	// without it, you will get an error
	(*sp).foo = 2
	(*sp).bar = 1

	// compiler will just infer this for us tho so
	sp.foo = 2
	sp.bar = 1
	// is equivalent and valid
	// which may be confusing when coming from C/C++
	// need to be careful to check if you are working with a pointer or a stack obj

	// what is the zero-value for a pointer????
	// NIL ... aka nullptr

	return "Pointers"
}

// we can provide arguments in this list style
// but ONLY if they are ALL of the SAME type
func multiArgSameType(a, b, c, d, e, f, g int) int {
	return a + b + c - d*e + f - g
}

// functions can accept pointers
func pointerAcceptor(p *int) {
	*p = 0 // set the value pointed to by p to 0
	// this WILL have side effects
}

type bigBoy struct {
	bigArr  [100]int
	hugeArr [1000]int
}

func takeStructByRef(pBigBoy *bigBoy) {
	fmt.Println("Working on the big boy.")
	for k, v := range pBigBoy.bigArr {
		v = k + k%2
		pBigBoy.hugeArr[k] = v
		if k%10 == 0 {
			fmt.Println(k)
		}
	}
}

// variadic parameters let us do some "templating" of sorts
// when passed, they act as a SLICE!!!
// can only have one variadic parameter, and it has to be LAST in the args list
// nothing stopping you from writing:
// 		func sum2(msg string, otherVal int, values ... int) int {}
func sum(values ...int) int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}

// Go does allow values to be returned as pointers
// this may seem counter-intuitive
// result would be stored on the stack of this function
// so the returned pointer should be pointing to dead memory right?
// Go Compiler will actually move the value to the heap on return
// then return the pointer to that heap space
// Make sure to notes some of the "handy" things the Go Runtime will do
// for you
// alternatively, just allocate pointers to the heap within functions
// so you are never confused by them
func sumReturnPointer(values ...int) *int {
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println("Moving stack variable to the heap")
	return &result
}

// here we are creating the pointer ourselves on the heap with new
// this makes everything that is happening a bit clearer, so i
// will do this
func sumReturnPointer2(values ...int) *int {
	result := new(int)
	for _, v := range values {
		*result += v
	}
	return result
}

// we can also do named return values like this
// the pointer should be stack allocated, but here
// will NOT be moved to the heap
// so this will create a "panic" event
func sumReturnPointer3(values ...int) (result *int) {
	for _, v := range values {
		*result += v
	}
	return // dont need to specify the variable, since it is
	// known to the compiler from the signature
}

func sumReturnPointer4(values ...int) (result int) {
	for _, v := range values {
		result += v
	}
	return // dont need to specify the variable, since it is
	// known to the compiler from the signature
}

// we can return as many values as we want from a function
// we just need to write them out in this (type1, type2, .. typeN) syntax

// A Go idiom for most functions is to return (val, error)
// you can see that the "error" is a return type
// this allows us to use val, err := func()
// and check errors in our code
// this is the typical approach to handling errors as opposed
// to throwing exceptions and catching them higher up the call stack
func multiReturnVal(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Divide by 0")
	}
	return a / b, nil // no error is nil
}

type greeter struct {
	greeting string
	name     string
}

// method for the greeter type
// takes by value with this style
// therefore modifications make no side effects
// (g greeter) is called a VALUE RECIEVER
func (g greeter) greetVal() {
	g.greeting = "Goodbye"
	fmt.Println(g.greeting, g.name)
}

// takes by reference (pointer) with this style
// therefore, modifications make side effects
// (g *greeter) is called a REFERENCE RECIEVER
func (g *greeter) greetRef() {
	g.greeting = "Goodbye"
	fmt.Println(g.greeting, g.name)
}

// VAL vs REF recievers
// VAl recievers by definition will not have side effects on objects,
// so that has inherent safety at the cost of COPYING
// REF only take a pointer to objs so the calls are always lightweight
// we will need to be much more careful about the modifications we
// are making to fields

// are we able to specify things as const as a non-modification guarantee
// to  the compiler?
// no.. . :(
// that actually kinda blows

// Functions shows basic syntax, parameters, returns, anonymous funcs, function as types, methods
func Functions() string {
	fmt.Println("\nShowing Functions Basics in Go...")

	// ok we know the basic syntax already
	// func NAME(param1 type, param2 type ...) returntype {}

	// extremely supid example alert
	a := multiArgSameType(1, 1, 1, 1, 1, 1, 1)
	// get the value pointed to by &a back and print
	fmt.Println("Before pass by reference", a)
	pointerAcceptor(&a)
	fmt.Println("Should print zero:", a)

	// big boy is a heavy weight struct with two large arrays in it
	// imaging passing this struct by value and having to copy everything
	// alternatively we can just pass a pointer which is just 8 bytes
	// (64-bit addresses on 64-bit machines)
	pBoy := new(bigBoy)
	takeStructByRef(pBoy)

	sum := sum(1, 2, 2, 3, 55, 11, 6, 2, 2, 52, 3, 52)
	fmt.Println("Sum from variadic function args func is:", sum)
	psum := sumReturnPointer(1, 2, 2, 3, 55, 11, 6, 2, 2, 52, 3, 52)
	fmt.Println("Sum from variadic function args func with pointer on stack moved to the heap is is:", *psum)
	psum = sumReturnPointer2(1, 2, 2, 3, 55, 11, 6, 2, 2, 52, 3, 52)
	fmt.Println("Sum from variadic function args func with heap pointer return is is:", *psum)
	sum = sumReturnPointer4(1, 2, 2, 3, 55, 11, 6, 2, 2, 52, 3, 52)
	fmt.Println("Sum from variadic function args func with named return is is:", sum)

	val, err := multiReturnVal(1.0, 0.0) // this will return an error since it divides by 0
	if err != nil {
		fmt.Println(val, err.Error())
	}

	// anonymous functions
	// anything you can do with any other type, you can do with functions in Go
	// FIRST CLASS CITIZENSHIP
	func() {
		fmt.Println("I am an invoked anonymous function")
	}()

	f := func() {
		fmt.Println("I am a an anonymous function saved to a var")
	}
	// invoke
	f()

	// anonymouse functions capture everything that is in their current
	// scope by default
	// this anonymous func captures "i" since it is part of the current context
	for i := 0; i < 5; i++ {
		func() {
			fmt.Println(i)
		}()
	}

	// it is better practice to actually pass something by value into your function
	// explicitly as shown below:
	for i := 0; i < 5; i++ {
		func(val int) {
			fmt.Println(val)
		}(i)
	}

	// full function signatures are as follows:
	var div func(a, b float64) (float64, error)
	div = func(a, b float64) (float64, error) {
		if b == 0.0 {
			return 0.0, fmt.Errorf("DIVIDE BY 0.0")
		}
		return a / b, nil
	}

	d, err := div(5, 4)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(d)

	// methods
	// are only declared outside of structs in Go
	// we use the syntax:
	// func (s structType) FUNCNAME(ARGS) RETURNS { BODY }
	g := greeter{
		greeting: "Hello",
		name:     "Greeter",
	}
	// invoking the method has a standard syntax
	// whether it is calling by val or reference
	g.greetVal()
	fmt.Println(g) // should be unmodified
	g.greetRef()
	fmt.Println(g) // should be modified

	return "Functions"
}

// interfaces are types just like structs
// so we have this common syntax

// in Go
// 		structs describe data
//		interfaces describe behaviors

// so structs have fields of data and can be composed
// 		of other structs which have their own fields

// and interfaces have methods and can be composed
// 		of other interfaces which have their own methods

// so we have this nice symmetry

// Writer writes stuff
type Writer interface {
	Write([]byte) (int, error)
}

// we generally create structs and then have them satisfy interfaces

// ConsoleWriter is a Writer to the ouput console
type ConsoleWriter struct{}

// TCPWriter is a Writer to a TCP connection
type TCPWriter struct{}

// FileWriter is a Writer to a File
type FileWriter struct{}

// we only implictly satisfy interfaces
// by creating their implementations for our structs
func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

func (tw TCPWriter) Write(data []byte) (int, error) {
	fmt.Println("shhh pretend im writing to a TCP connection")
	n, err := fmt.Println(string(data))
	return n, err
}

func (fw FileWriter) Write(data []byte) (int, error) {
	fmt.Println("shhh pretend im writing to a file")
	n, err := fmt.Println(string(data))
	return n, err
}

// note since all types are treated equally, we can do something like this:

// Incrementer increments things
type Incrementer interface {
	Increment() int
}

// IntCounter is just a special type alias of an int, that satsifies the Incrementer interface
type IntCounter int

// Increment for the IntCounter
func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}

// Composing interfaces together is a key concept in Go

// Closer closes
type Closer interface {
	Close() error
}

// WriterCloser is an interface composed of Writer and Closer
type WriterCloser interface {
	Writer
	Closer
}

// BufferedWriterCloser has a buffer and writes/closes
type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}

	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}

		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}

// Close goes thru the whole buffer and flushes in 8 byte chunks
func (bwc *BufferedWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

// NewBufferedWriterCloser makes a new buffered writercloser
func NewBufferedWriterCloser() *BufferedWriterCloser {
	fmt.Println("Creating new BufferedWriterCloser object")
	bwc := new(BufferedWriterCloser)
	bwc.buffer = bytes.NewBuffer([]byte{})
	return bwc
}

// Interfaces are contracts that a struct must fulfil (generally in implementing some kind of method)
func Interfaces() string {
	fmt.Println("\nShowing Interfaces Basics in Go...")

	// we can create a variable that is of an interface type
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Using a Writer interface!"))

	// we can create an array of structs
	// that satisfy the same interface
	// and call methods on all of them
	writers := [3]Writer{ConsoleWriter{}, TCPWriter{}, FileWriter{}}

	for _, w := range writers {
		w.Write([]byte("Using a Writer interface!"))
	}

	var ic IntCounter = 0

	for i := 0; i < 20; i++ {
		ic.Increment()
	}

	fmt.Println(int(ic)) // should be 20 right?

	var wc WriterCloser = NewBufferedWriterCloser() // define as an interface
	wc.Write([]byte("What is up boys, please like and subscribe!"))
	wc.Close()

	// type conversion
	// this will work because BufferedWriterCloser
	// fulfulls this interface
	bwc, ok := wc.(*BufferedWriterCloser) // now convert to struct
	if ok {
		fmt.Println(bwc)
	} else {
		fmt.Println("Conversion Failed")
	}

	// we can try to convert the WriterCloser to an io.Reader
	// but this will not work because io.Reader requires a Read interface
	// we can add some error checking so we don't panic out of execution
	r, ok := wc.(io.Reader)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("Conversion Failed")
	}

	// we can use something called the empty interface
	var empty interface{} = NewBufferedWriterCloser()
	// what is the point of this thing?
	// to do anything useful with it, we need to convert it to some other interface
	// as shown below
	// we can do that statically in code as we have here
	// or we can use the reflect package to do some slick
	// type switching to determine what to use
	if conv, ok := empty.(WriterCloser); ok {
		conv.Write([]byte("Writing into interface converted from empty interface"))
		conv.Close()
	}

	// lets look at empty interfaces with type switching

	// so say i recieve some generic interface in a function
	// i may want to do this
	var i interface{} = 0
	switch i.(type) {
	case int:
		fmt.Println("i is an int")
	case string:
		fmt.Println("i is a string")
	default:
		fmt.Println("idk what i is...")
	}

	// Best Practices with Interfaces

	// use many, small interfaces
	// 		compose them together to make larger interfaces
	//		LEGO style

	// Don't export interfaces for types that will be consumed
	// Do export interfaces for types that will be used by package

	// Design functions and methods to recieve interfaces whenever possible

	return "Interfaces"
}

func printMsg(msg string) {
	fmt.Println(msg)
	wg.Done()
}

// wait group is like a list of pending go routines
// we can use it to wait for execution of a spawned goroutine
// from a "main" thread
var wg = sync.WaitGroup{}
var wgCounter = 0
var m = sync.RWMutex{}

func increment() {
	wgCounter++
	wg.Done()
}

func printCounter() {
	fmt.Println("counter:", wgCounter)
	wg.Done()
}

func incrementWithMutex() {
	m.Lock() // lock mutex (read + write)
	wgCounter++
	m.Unlock()
	wg.Done()
}

func printCounterWithMutex() {
	m.RLock() // read lock the mutex
	fmt.Println("counter:", wgCounter)
	m.RUnlock()
	wg.Done()
}

func incrementWithMutex2() {
	wgCounter++
	m.Unlock()
	wg.Done()
}

func printCounterWithMutex2() {
	fmt.Println("counter:", wgCounter)
	m.RUnlock()
	wg.Done()
}

// GoRoutines details Go's lightweight process, the goroutine
func GoRoutines() string {
	fmt.Println("\nShowing GoRoutine Basics in Go...")

	// instead of using OS threads
	// Go uses lightweight processes (user-space threads)
	// so we less startup/teardown costs
	// and theoretically don't need to do things like thread pooling

	// go routines are just abstractions of these user threads
	// the go runtime maps go routines onto the actual OS threads for us

	wg.Add(3)
	go printMsg("I am running as a go routine")

	go func() {
		printMsg("I am a go routine running an anonymous function")
	}()

	msg := "I am a message given to anon func"
	go func() {
		printMsg(msg) // this will be a different thread,
		// but the Go runtime will still know where to access msg at
		// we have introduced a dependency from this master thread
		// and this go routine tho, so it is starting to get spicy
	}()
	msg = "I am a messaged that changed after the go routine call"
	// we might get the inital msg value or this second value
	// no real way to know, undefined behavior
	wg.Wait()

	// so generally, we do not want to play around with this closure stuff
	// we could rewrite the function as one that takes the msg by value
	// and resolve this issue

	fmt.Println("Trying again but passing message to goroutine by value")

	wg.Add(1)
	// reset
	msg = "I am a message given to anon func"
	go func(msg string) {
		printMsg(msg) // passed this goroutine a value, so problem solved
	}(msg)
	msg = "I am a messaged that changed after the go routine call"
	wg.Wait() // waiting for signal

	fmt.Println("Unsyncronized goroutines...")
	// this loop will spawn 20 goroutines
	// but we will have no sync BETWEEN ROUTINES
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go printCounter()
		go increment()
	}
	wg.Wait() // wait until they are all done

	fmt.Println("Syncronized goroutines with mutexes...")
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go printCounterWithMutex()
		go incrementWithMutex()
	}
	wg.Wait() // wait until they are all done

	fmt.Println("Syncronized goroutines with mutexes outside of calls...")
	// now that we are locking the mutexes in the SAME context
	// what happens is guaranteed to be ORDERED
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go printCounterWithMutex2()
		m.Lock()
		go incrementWithMutex2()
	}
	wg.Wait() // wait until they are all done

	// this basically is making everything be single threaded tho...
	// great
	// single threaded + mutex overhead = worse than original

	// the runtime packages lets us query things like max number of threads
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(-1))

	// think of GOMAXPROCS as a tuning parameter for your
	// parallel applications

	// minimum 1 thread per core
	newMaxProcs := 100
	fmt.Println("Setting GOMAXPROCS to", newMaxProcs)
	runtime.GOMAXPROCS(newMaxProcs)
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(-1))

	// this is the # of OS threads, so just creating a TON
	// will add lots of memory overhead
	// remember the scheduler maps goroutines as LWPs
	// ONTO OS threads
	// so you dont need as many OS Threads as goroutines or something

	// BEST PRACTICES
	// 		if making a library, try to limit goroutines
	//		let consumer use them

	// 		when creating goroutines, know how it will end
	//		we want to return resources as soon as we can

	//		check for race conditions at compile time
	//		"go build -race" will check for race conditions
	// 		"go build -msan" will do address sanitizer
	//		"go run -race"  will actually run and print stack trace for us

	return "GoRoutines"
}

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var logCh = make(chan logEntry, 50)
var doneCh = make(chan struct{}) // empty struct
// this is what is known as a signal only channel
// so it is some kind of semaphore stuff

// we now have a select{} control block in this function
// select multiplexes incoming signals
func logger() {
	for {
		select {
		case entry := <-logCh:
			fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)

		case <-doneCh:
			break
		}
	}
}

// Channels are how we can pass data between threads in go
func Channels() string {
	fmt.Println("\nShowing Channels Basics in Go...")
	// channels are pretty much always going to be used in the context
	// of goroutines, since it is for passing data between threads
	// so, they build off of goroutines, and are what truly make them
	// flexible

	// we have to use the make() function to make a channel
	// this sort of makes sense, since make will put something on
	// the heap, and we will need to place something in "shared memory"
	// for it to be seen by other threads
	ch := make(chan int) // make a channel for an int
	// here we made a strongly typed channel, assuming that means we can make
	// generic channels

	// SEND ON CHANNEL
	//		ch <- value

	// RECIEVE FROM CHANNEL
	//	 	var := <- ch

	// so if arrow is pointing TO channel, send
	// if arrow is pointing FROM channel, recieve from

	// ch <- 		SEND
	// <- ch		RECV

	wg.Add(2)
	go func() {
		i := <-ch // recieve from channel
		fmt.Println(i)
		wg.Done()
	}()

	go func() {
		ch <- 42 // send on channel
		wg.Done()
	}()
	wg.Wait()

	// we need a RECV for every SEND
	// otherwise we will get a deadlock and crash

	// from this, we can discern that SENDs and RECVs over channels
	// are BLOCKING

	// most of the time (mainly for good design's sake)
	// we will want to have a function only SEND or RECV
	// from a channel
	// we can do that with the following syntax:

	//	func(NAME <-chan type)	RECV
	//	func(NAME chan<- type)	SEND

	// if we tried to do a SEND in the func that took a RECV chan
	// we would get a compile time error
	// this gives us more type safety, and encourages modularized designs
	fmt.Println("Unidirectional Channels in a loop:")
	for j := 0; j < 5; j++ {
		wg.Add(2)
		go func(ch <-chan int, j int) { // takes a RECV channel only
			i := <-ch // recieve from channel
			fmt.Println(i)
			wg.Done()
		}(ch, j) // note we are inputting a bidirectional chan
		// and the compiler will just note that it is RECV only

		go func(ch chan<- int, j int) { // takes a SEND channel only
			ch <- j // send on channel
			wg.Done()
		}(ch, j)
	}
	wg.Wait()

	// BUFFERED CHANNELS
	// with buffered channels, we now have non-blocking behavior!
	// we still need to make sure we process all of the data though
	// below, we have  buffer of size 2, so both messages can be
	// SENT without blocking, but we only read one out then leave
	// the goroutine, essentially losing the 45 value on the channel
	fmt.Println("Sending 2 values to buffered channel (size 2)")
	ch = make(chan int, 2)
	wg.Add(2)
	go func(ch <-chan int) {
		i := <-ch // recieve from channel
		fmt.Println(i)
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 42 // send on channel
		ch <- 45 // send on channel
		wg.Done()
	}(ch)
	wg.Wait()

	// iterate over buffered channel to process all data
	chanSize := 50
	sendSize := chanSize - 5
	ch = make(chan int, chanSize)
	fmt.Printf("Sending %d values to buffered channel (size %d)\n", sendSize, chanSize)

	wg.Add(2)
	go func(ch <-chan int) {
		for {
			if val, ok := <-ch; ok {
				fmt.Println(val)
			} else {
				break
			}
		}
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		for i := 0; i < sendSize; i++ {
			ch <- 42 + i%3
		}
		close(ch) // since we sent less than the full channel size
		//	the channel will deadlock in the for range loop above
		// unless we explicitly close the channel, like we do here

		// this creates a NEW problem, where we now have closed the channel
		// if we try to use it again for anyting, we in trouble
		// we have to make a new channel now
		wg.Done()
	}(ch)
	wg.Wait()

	//fmt.Println("Starting Logger....")
	//go logger()

	//logCh <- logEntry{time.Now(), logInfo, "App is starting"}
	//time.Sleep(100 * time.Millisecond)

	//logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}
	//time.Sleep(100 * time.Millisecond)

	//doneCh <- struct{}{} // this syntax is kind of jank
	// but this is how you send a blank semaphore in Go

	return "Channels"
}

func prepaterTestDirTree(tree string) (string, error) {
	tmpDir, err := ioutil.TempDir("", "")
	if err != nil {
		return "", fmt.Errorf("error creating temp directory: %v\n", err)
	}

	err = os.MkdirAll(filepath.Join(tmpDir, tree), 0755)
	if err != nil {
		os.RemoveAll(tmpDir)
		return "", nil
	}

	fmt.Printf("Temp dir to walk: %s\n", tree)
	return tmpDir, nil
}

// Filepath shows functionality of the "path/filepath" Go library package
func Filepath() string {
	fmt.Println("\nShowing path/filepath Basics in Go...")

	fmt.Printf("os.PathSeparator:\t%s\n", string(os.PathSeparator))
	fmt.Printf("os.PathListSeparator:\t%s\n", string(os.PathListSeparator))

	myPath := "."
	myAbsPath, err := filepath.Abs(myPath)
	if err != nil {
		fmt.Errorf("Could not get absolute path")
	}

	fmt.Printf("My relative path:\t%s\n", myPath)
	fmt.Printf("My absolute path:\t%s\n", myAbsPath)

	// Base takes the bottom (right-most)
	// for example, if my path is /foo/bar
	// the base is bar
	testBase := filepath.Base(myAbsPath)
	fmt.Printf("My base directory:\t%s\n", testBase)

	// we can use Join() to combine path elements with the PathSeparator
	// Join()ed paths are auto-cleaned
	wonkyPath := filepath.Join(myAbsPath, "..")
	fmt.Printf("New Path:\t%s\n", wonkyPath)
	// Clean() returns the shorted path equivalent
	// for example, if my path is /foo/bar/../bar/..
	// the path will be /foo
	dirtyPath := myAbsPath + "/.."
	fmt.Printf("Dirty Path:\t%s\n", dirtyPath)
	cleanedPath := filepath.Clean(dirtyPath)
	fmt.Printf("Cleaned Path:\t%s\n", cleanedPath)

	relativePath, err := filepath.Rel(cleanedPath, myAbsPath)
	if err != nil {
		fmt.Errorf("Could not get relative path")
	}
	fmt.Printf("Relative path of cleaned path to this dir:\t%s\n", relativePath)

	fileToGet := "go.mod"
	ext := filepath.Ext(fileToGet)
	fmt.Printf("Extension of file %s is %s\n", fileToGet, ext)

	// Walk() lets us "walk" a file tree rooted at root
	// we pass it a walkFn for each file or dir in the tree
	tmpDir, err := prepaterTestDirTree("dir/to/walk/skip")
	if err != nil {
		fmt.Printf("Unable to create test dir tree: %v\n", err)
		return ""
	}

	defer os.RemoveAll(tmpDir)
	os.Chdir(tmpDir)

	subDirToSkip := "skip"

	// the walkFn is called for every file or dir
	// that the Walk() function interacts with
	// Walk() will pass the path string
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		// if it is a directory and the name is
		// the skip directory, we skip
		if info.IsDir() && info.Name() == subDirToSkip {
			fmt.Printf("Skipping a dir without errors: %+v\n", info.Name())
			return filepath.SkipDir
		}

		fmt.Printf("Visited file or dir: %q\n", path)
		return nil
	}

	fmt.Printf("On Unix:\n")
	err = filepath.Walk(".", walkFn)
	if err != nil {
		fmt.Printf("Error walking the path %q: %v\n", tmpDir, err)
		return ""
	}

	return "Filepath"
}

func pwd() string {
	pwd := os.Getenv("PWD")
	if pwd == "" {
		fmt.Printf("Unable to get pwd\n")
		return ""
	}
	fmt.Printf("%s = %s\n", "PWD", pwd)

	fileinfo, err := os.Lstat(pwd) // return fileinfo struct of dir
	if err != nil {
		fmt.Printf("Could not get dir info: %s\n", err)
		return ""
	}

	if fileinfo.Mode()&os.ModeSymlink != 0 { // AND bitmasks
		realpath, err := filepath.EvalSymlinks(pwd)
		if err != nil {
			fmt.Printf("Error getting real path: %v\n", err)
			return ""
		}

		fmt.Printf("PWD: %s\n", realpath)
		return realpath

	}
	return ""
}

func ensureBaseDir(fpath string) error {
	baseDir := filepath.Dir(fpath)
	info, err := os.Stat(baseDir)
	if err == nil && info.IsDir() {
		return nil
	}
	return os.MkdirAll(baseDir, 0755)
}

func permissions() {
	//pwd := pwd()
	err := ensureBaseDir("./tmp")
	if err != nil {
		log.Fatal(err)
	}

	emptyFile, err := os.Create("./tmp/dummyFile.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer emptyFile.Close()

	log.Printf("%v\n", emptyFile)
}

// OS covers what is inside th OS Go packages
func OS() string {
	fmt.Println("\nShowing os Basics in Go...")

	// os.Chdir(dir) - cd dir
	// os.Chmod(name, mode) - cmod mode file

	// os.Environ() prints all environment variables
	env := os.Environ()
	fmt.Printf("Checking if you have Chapel installed...\n")
	for _, e := range env {
		if strings.Contains(e, "CHPL_HOME") {
			fmt.Printf("YES!\t%s\n", e)
			break
		}
	}

	// Executable returns our executable location
	exe, err := os.Executable()
	if err != nil {
		fmt.Errorf("Error getting executable: %v", err)
	}
	fmt.Printf("Executable path: %s\n", exe)

	// os.Exit will forcibly terminate the program
	// if you uncomment what is below, the deferred
	// print will not be completed, as the process
	// will just completely exit
	//defer fmt.Printf("I WONT BE PRINTED")
	//os.Exit(1)

	// os.Expand performs mapping to strings similar to
	// ${var} expansion in a CLI in Linux
	// echo ${CC} -> clang++
	mapper := func(placeHolderName string) string {
		switch placeHolderName {
		case "DAY_PART":
			return "morning"
		case "NAME":
			return "Gopher"
		}

		return ""
	}

	fmt.Printf("Expanding: Good ${DAY_PART}, $NAME! to...\n")
	fmt.Printf("%v\n", os.Expand("Good ${DAY_PART}, $NAME!", mapper))

	// os env are basically stored in a map
	// we use a key to os.Getenv and get
	// the value back out
	chapelPath := os.Getenv("CHPL_HOME")
	fmt.Printf("%s = %s\n", "CHPL_HOME", chapelPath)

	// Getpagesize could be useful to get some
	pageSize := os.Getpagesize()
	fmt.Printf("Page Size: %d bytes\n", pageSize)

	pid := os.Getpid()   // process id of caller
	ppid := os.Getppid() // process id of callers parent
	uid := os.Getuid()   // user id of caller
	fmt.Printf("pid\t\t%d\nppid\t\t%d\nuid\t\t%d\n", pid, ppid, uid)

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Errorf("Error getting Hostname: %v", err)
		return ""
	}
	fmt.Printf("hostname\t%s\n", hostname)

	/* FILEMODES
	   // The single letters are the abbreviations
	   // used by the String method's formatting.
	   ModeDir        FileMode = 1 << (32 - 1 - iota) // d: is a directory
	   ModeAppend                                     // a: append-only
	   ModeExclusive                                  // l: exclusive use
	   ModeTemporary                                  // T: temporary file; Plan 9 only
	   ModeSymlink                                    // L: symbolic link
	   ModeDevice                                     // D: device file
	   ModeNamedPipe                                  // p: named pipe (FIFO)
	   ModeSocket                                     // S: Unix domain socket
	   ModeSetuid                                     // u: setuid
	   ModeSetgid                                     // g: setgid
	   ModeCharDevice                                 // c: Unix character device, when ModeDevice is set
	   ModeSticky                                     // t: sticky
	   ModeIrregular                                  // ?: non-regular file; nothing else is known about this file

	   // Mask for the type bits. For regular files, none will be set.
	   ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice | ModeCharDevice | ModeIrregular

	   ModePerm FileMode = 0777 // Unix permission bits
	*/

	pwd()
	permissions()

	return "OS"
}
