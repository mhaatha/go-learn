package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorldSkip(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("This test only run in any operating system except linux")
	}

	result := HelloWorld("Eko")
	require.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Eko")

	require.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")

	fmt.Println("This line will not executed if error")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Eko")

	assert.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")

	fmt.Println("This line still executed even the test is error")
}

func TestHelloWorldEko(t *testing.T) {
	result := HelloWorld("Eko")

	if result != "Hello Eko" {
		// panic("Result is not 'Hello Eko'") // Bad Error Handling
		t.Error("Result must be 'Hello Eko'")
	}

	fmt.Println("TestHelloWorldEko Done")
}

func TestHelloWorldKhannedy(t *testing.T) {
	result := HelloWorld("Khannedy")

	if result != "Hello Khannedy" {
		// panic("Result is not 'Hello Khannedy'") // Bad Error Handling
		t.Fatal("Result must be 'Hello Khannedy'")
	}

	fmt.Println("TestHelloWorldKhannedy Done")
}
