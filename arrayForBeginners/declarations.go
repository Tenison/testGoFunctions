package arrayForBeginners

import (
	"fmt"
)


func DeclareAndPrint(){
	var newCount [8]int

	newCount[0] = 1

	//using make to arrays
	newSlice := make([]int, 10)//(type, size, optional = "capacity")
	newSlice[0] = 5 //doesn't change the size of the array

	newSliceTwo := make([]int, 10)//using append() and indexing
	newSliceTwo[0] = 12 
	newSliceTwo = append(newSliceTwo, 5) //inceases the size of the array

	newSliceThree := make([]int, 10, 100)//(type, size, optional = "capacity")
	newSliceThree = append(newSliceThree, 5)


	//declare and initialize
	var names [4]string = [4]string{"Alex", "Osei", "Owusu", "PJ"}
	numberSeries := [...]int{1,2,3,4,5,6,7,8}

	

	fmt.Printf("newCount : %v\n", newCount)
	fmt.Printf("new Slice : %v %v\n", newSlice, cap(newSlice))
	fmt.Printf("new SliceTwo : %v %v\n", newSliceTwo, cap(newSliceTwo))
	fmt.Printf("new SliceThree : %v %v\n", newSliceThree, cap(newSliceThree))
	fmt.Printf("One Name : %v\n", string(names[0]))
	fmt.Printf("numberSeries : %v", numberSeries)
}