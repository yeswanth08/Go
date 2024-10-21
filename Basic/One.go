package main

import "fmt"

func main(){

	var a int = 10;
	fmt.Printf("type of a %T",a)
	// a = "hello" 
	// go ensures that the type of a variable is same as the type of the value assigned to it	
	fmt.Printf("type of a %T",a)

	b := 25;

	fmt.Printf("bin of b %b",b);

	/*
		%T	a Go-syntax representation of the type of the value
		%v	the value in a default format
		%q	a double-quoted string safely escaped with Go syntax
		%x	base 16, with lower-case letters for a-f

		untyped decalration is automatically converted to the type of the value assigned to it
		used when we dont know the type of the value
	*/

	var temparr = [3]int{1,2,3}

	var temparr2 = [...]int{1,2,3,4,5} // ... makes the compiler count the number of elements in the array
	
}