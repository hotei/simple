// jsontest.go  (c) David Rook 2010 - released under Simplified BSD 2-clause License

/*
 * key observation here is that result of a 'new()' is a pointer to object
 * that's exactly what JSON wanted as receiver in Unmarshal()
 *
 * This compiles and runs ok.  Not absolutely sure if it does what's intended
 * Need more testing...
 */

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Message struct {
	Name  string
	Phone string
}

var r = Message{"myname", "myphone"}
var inpt_1 string = "{\"Name\":\"myname\",\"Phone\":\"myphone\"}"
var inpt_2 string = "{\"Name\":\"myname\", \"Phone\":\"myphone\", \"Email\":\"myemail\" }"

func jsonwrite() {
	x, err := json.Marshal(r)
	if err != nil {
		fmt.Printf("error %v\n", err)
		return
	}
	y := string(x)
	fmt.Printf("marshalled r struct is : %s\n", y)
	fmt.Printf("r should be : %s\n", inpt_1)
	if y != inpt_1 {
		fmt.Printf("Error\n")
	}
}

func jsonread() {
	newrec := new(Message)
	newrec.Name = "noname"
	newrec.Phone = "nophone"
	fmt.Printf("input value to JSON: %s\n", inpt_1)
	input_one := []byte(inpt_1) // explicit conversion of string to []byte
	err := json.Unmarshal(input_one, newrec)
	if err != nil {
		fmt.Printf("jsonread failed at 1\n")
		os.Exit(-1)
	}
	fmt.Printf("name -> %v\n", newrec.Name)
	fmt.Printf("phone-> %v\n", newrec.Phone)
	fmt.Printf("result= %v\n", *newrec)
	if newrec.Phone != "myphone" || newrec.Name != "myname" {
		fmt.Printf("JSON read test FAILED\n")
	}
	input_two := []byte(inpt_2)
	fmt.Printf("input value to JSON: %s\n", inpt_2)
	err = json.Unmarshal(input_two, newrec)
	if err != nil {
		fmt.Printf("jsonread failed at 2\n")
		os.Exit(-1)
	}
	fmt.Printf("%v\n", newrec)

	fmt.Printf("JSON read test PASSED\n")
}

func main() {
	fmt.Printf("JSON test <start>\n")
	jsonwrite()
	fmt.Printf("JSON write test fini\n")
	jsonread()
	fmt.Printf("JSON read test fini\n")
	fmt.Printf("\nJSON test <end>\n")
}
