	package main
	import (
		"fmt"
		"greeting"
	)
	func negate(myBoolean *bool)  {
		*myBoolean = !*myBoolean
		//return *myBoolean
	}
	func main() {
		truth := true
		negate(&truth)
		fmt.Println(truth)
		lies := false
		negate(&lies)
		fmt.Println(lies)
		greeting.Hi()
		greeting.Hello()
	}