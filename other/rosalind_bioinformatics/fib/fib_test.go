package fib

//inputFilePath := "../input/fib_test_input.txt"
import (
	"fmt"
	"testing"
)

func Test_fib(t *testing.T) {
	var cases = []struct {
		fpath string
		total int
	}{
		{
			"../input/fib/test_input_01.txt",
			19,
		},
		{
			"../input/fib/test_input_02.txt",
			875089148811941,
		},
	}
	for _, c := range cases {
		rV, err := fib(c.fpath)
		if err != nil {
			t.Errorf("Error retriving counts: %v", err)
		}
		if rV != c.total {
			t.Errorf("total rabbit values do not match - got %v want: %v\n", rV, c.total)
		}
		fmt.Printf("file = %v, counts = %v\n", c.fpath, rV)
	}
}
