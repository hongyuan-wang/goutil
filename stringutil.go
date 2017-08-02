package goutil

import "fmt"

func PrintStrEle(s string) int {

	count := 0
	for _, e := range s {
		fmt.Printf("%c\n", e)
		count++
	}

	return count

}
