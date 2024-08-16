package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkSub(b *testing.B) {
	b.Run("Bench", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Bench")
		}
	})

	b.Run("World", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("World")
		}
	})
}

func BenchmarkHelloBench(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Bench")
	}
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("World")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWorld(Eko)",
			request:  "Eko",
			expected: "Hello Eko",
		},
		{
			name:     "HelloWorld(Kurniawan)",
			request:  "Kurniawan",
			expected: "Hello Kurniawan",
		},
		{
			name:     "HelloWorld(Khannedy)",
			request:  "Khannedy",
			expected: "Hello Khannedy",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			message := fmt.Sprintf("Result must be '%s'", test.expected)
			require.Equal(t, test.expected, result, message)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Sub Test 1", func(t *testing.T) {
		result := HelloWorld("Sub Test 1")
		require.Equal(t, "Hello Sub Test 1", result, "Result must be 'Hello Sub Test 1'")
	})

	t.Run("Sub Test 2", func(t *testing.T) {
		result := HelloWorld("Sub Test 2")
		require.Equal(t, "Hello Sub Test 2", result, "Result must be 'Hello Sub Test 2'")
	})
}

func TestMain(m *testing.M) {
	fmt.Println("=============================")
	fmt.Println("BEFORE TEST")
	fmt.Println("=============================")

	m.Run()

	fmt.Println("=============================")
	fmt.Println("AFTER TEST")
	fmt.Println("=============================")
}

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
