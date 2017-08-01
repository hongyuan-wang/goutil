package goutil

import "fmt"

func PrintStrEle(s string) {

	for _, e := range s {
		fmt.Printf("%c\n", e)
	}

}
