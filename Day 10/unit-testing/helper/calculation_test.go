package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFailSum(t *testing.T) {
	result := sum(10, 10)

	require.Equal(t, 40, result, "Result has to be 40")

	fmt.Println("TestFailSum Eksekusi Terhenti")
}

func TestSum(t *testing.T) {
	result := sum(10, 10)

	assert.Equal(t, 20, result, "Result has to be 20")

	if result != 20 {
		panic("Result should be 20")
	}
}

func TestTableSum(t *testing.T) {
	tests := []struct {
		request int
		result  int
		errMsg  string
	}{
		{
			request: sum(10, 10),
			result:  20,
			errMsg:  "Result has to be 20",
		},
		{
			request: sum(29, 20),
			result:  40,
			errMsg:  "Result has to be 40",
		},
		{
			request: sum(30, 30),
			result:  60,
			errMsg:  "Result has to be 60",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			require.Equal(t, test.request, test.result, test.errMsg)
		})
	}
}
