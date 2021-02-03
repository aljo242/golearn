package golearn

import (
	"testing"
)

func TestHello(t *testing.T) {
	expected := "Hello World!"
	if ret := Hello(); ret != expected {
		t.Errorf("Hello() = %q, want %q", ret, expected)
	}
}

func TestDeclarations(t *testing.T) {
	expected := "Declarations"
	if ret := Declarations(); ret != expected {
		t.Errorf("Declarations() = %q, want %q", ret, expected)
	}
}

func TestConversions(t *testing.T) {
	expected := "Conversions"
	if ret := Conversions(); ret != expected {
		t.Errorf("Conversions() = %q, want %q", ret, expected)
	}
}
func TestPrimitives(t *testing.T) {
	expected := "Primitives"
	if ret := Primitives(); ret != expected {
		t.Errorf("Primitives() = %q, want %q", ret, expected)
	}
}

func TestConstants(t *testing.T) {
	expected := "Constants"
	if ret := Constants(); ret != expected {
		t.Errorf("Constants() = %q, want %q", ret, expected)
	}
}

func TestArraysAndSlices(t *testing.T) {
	expected := "ArraysAndSlices"
	if ret := ArraysAndSlices(); ret != expected {
		t.Errorf("ArraysAndSlices() = %q, want %q", ret, expected)
	}
}
func TestMapsAndStructs(t *testing.T) {
	expected := "MapsAndStructs"
	if ret := MapsAndStructs(); ret != expected {
		t.Errorf("MapsAndStructs() = %q, want %q", ret, expected)
	}
}

func TestControlFlow(t *testing.T) {
	expected := "ControlFlow"
	if ret := ControlFlow(); ret != expected {
		t.Errorf("ControlFlow() = %q, want %q", ret, expected)
	}
}

func TestLoops(t *testing.T) {
	expected := "Loops"
	if ret := Loops(); ret != expected {
		t.Errorf("Loops() = %q, want %q", ret, expected)
	}
}
func TestDeferPanicRecover(t *testing.T) {
	expected := "DeferPanicRecover"
	if ret := DeferPanicRecover(); ret != expected {
		t.Errorf("DeferPanicRecover() = %q, want %q", ret, expected)
	}
}

func TestPointers(t *testing.T) {
	expected := "Pointers"
	if ret := Pointers(); ret != expected {
		t.Errorf("Pointers() = %q, want %q", ret, expected)
	}
}

func TestFunctions(t *testing.T) {
	expected := "Functions"
	if ret := Functions(); ret != expected {
		t.Errorf("Functions() = %q, want %q", ret, expected)
	}
}

func TestInterfaces(t *testing.T) {
	expected := "Interfaces"
	if ret := Interfaces(); ret != expected {
		t.Errorf("Interfaces() = %q, want %q", ret, expected)
	}
}

func TestGoRoutines(t *testing.T) {
	expected := "GoRoutines"
	if ret := GoRoutines(); ret != expected {
		t.Errorf("GoRoutines() = %q, want %q", ret, expected)
	}
}

func TestChannels(t *testing.T) {
	expected := "Channels"
	if ret := Channels(); ret != expected {
		t.Errorf("Channels() = %q, want %q", ret, expected)
	}
}

func TestFilepath(t *testing.T) {
	expected := "Filepath"
	if ret := Filepath(); ret != expected {
		t.Errorf("Filepath() = %q, want %q", ret, expected)
	}
}
