// package just collection of code
// main mean this code is going to run on terminal

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var p1 = fmt.Println

func main() {
	p1("what is ur name")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')

	if err == nil {
		p1("Hello", name)
	} else {
		log.Fatal(err)
	}
}
