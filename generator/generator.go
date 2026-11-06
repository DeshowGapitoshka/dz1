package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.Create("collisions.in")
	defer f.Close()

	blocks := []string{"B0", "AQ"} // 'B'*33+'0' == 'A'*33+'Q' == 2226
	for i := 0; i < 1000; i++ {
		s := ""
		x := i
		for j := 0; j < 17; j++ { // 17 блоков = 34 символа
			if x&1 == 0 {
				s += blocks[0]
			} else {
				s += blocks[1]
			}
			x >>= 1
		}
		fmt.Fprintf(f, "INSERT %s %d\n", s, i)
	}
}
