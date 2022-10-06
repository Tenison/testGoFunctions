package main

import (
	"fmt"
	"testGoFunc/loopsForBeginners"
)

func main() {
	fmt.Println("<<< Welcome to My Testing world >>>")
	fmt.Println("<<< Your Result will be Printed below >>>")
	
	//Test Example here
	//fmt.Println("Result is ",stringCharaters.ReplaceCharater("Hii There!!", 5, "t"))
	
	//Test arrayForBeginners
	//arrayForBeginners.DeclareAndPrint()

	//testing basic loops
	for _, value := range(loopsForBeginners.BasicLoop()){
		fmt.Println("count : ", value)
	}
	
	fmt.Println("Values greater than five: ", loopsForBeginners.WhileLoops())


}
