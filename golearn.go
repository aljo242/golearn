package golearn

import (
	"fmt"
	"reflect"
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
}

// Conversions shows some basic concepts of converting between types
// in Go's strong typed system
func Conversions() {
	fmt.Println("\nShowing basic conversions in Go's strong type system...")
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
}

// Constants covers:
// naming convention
// typed constants
// untyped constants
// enumerated constants
// enumeration expressions
func Constants() {
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
}

// ArraysAndSlices first details arrays, which are the basis
// for slices, then slices, which allow for dynamic views
// of allocated memory
func ArraysAndSlices() {
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
func MapsAndStructs() {
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
}
