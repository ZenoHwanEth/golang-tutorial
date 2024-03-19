// package just collection of code
// main mean this code is going to run on terminal

package main

import (
	"fmt"
	"reflect"
)

var p1 = fmt.Println

func main() {
	// 1.
	// p1("what is ur name")
	// reader := bufio.NewReader(os.Stdin)
	// name, err := reader.ReadString('\n')

	// if err == nil {
	// 	p1("Hello", name)
	// } else {
	// 	log.Fatal(err)
	// }

	// 2. variable
	// var name type
	// var vName string = "Derek"
	// var v1, v2 = 1.2, 3.4
	// var v3 = "Hello"
	// // variable is mutable
	// v4 := 2.4 // auto define data type
	// v4 = 5.4

	// 3. data type
	// int, float64, bool, string, rune
	// Default type 0, 0.0, false, ""
	p1(reflect.TypeOf(25))
	p1(reflect.TypeOf(3.14))
	p1(reflect.TypeOf(true))
	p1(reflect.TypeOf("jhello"))
	p1(reflect.TypeOf('😍'))
	cV1 := 1.5
	cV2 := int(cV1)
	p1(cV2)
}
