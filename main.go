package main

import (
	"fmt"
	"testGoFunc/stringCharaters"
)

func main() {
	fmt.Println("<<< Welcome to My Testing world >>>")
	fmt.Println("<<< Your Result will be Printed below >>>")
	
	//Test Example here
	fmt.Println("Result is ",stringCharaters.ReplaceCharater("Hii There!!", 5, "t"))
}
