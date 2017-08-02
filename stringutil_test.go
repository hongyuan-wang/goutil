package goutil

import "testing"
import "fmt"

func TestPrintStrEle(t *testing.T) {
	cases := []string{
		"string",
		"hello",
	}

	for _, str := range cases {
		count := PrintStrEle(str)
		if count == len(str) {
			fmt.Println("test %s ok", str)
		} else {
			t.Errorf("test %s error for %d, excepted %d ", str, count, len(str))
		}
	}

}
