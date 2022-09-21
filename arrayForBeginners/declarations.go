package arrayForBeginners

import (
	"fmt"
)

/**
* Arrays, Slices, loop
**/
func ArrayGo(){
	var firstArray [2]int

	for i,j := 0,1 ; i < 2; i,j = i+1, j+1{// i+1 === i++
		firstArray[i] = j
	}

	var secondArray [3]string = [3]string{"Alex", "Osei", "PJ"}

	thirdArray := [...]int{1,2,3}

	var firstNames [3]string = secondArray //creates a copy in a different memory location(not a reference type)
	
	//We can't append to arrays because they have a fixed size that does not change
	//firstNames = append(firstNames, "Yaw")
	
	for _,v := range(firstNames) {
		fmt.Printf("secondArray : %v\n", string(v))
	}

	fmt.Printf("firstArray : %v\n", firstArray)
	fmt.Printf("secondArray : %v\n", secondArray)
	fmt.Printf("thirdArray : %v\n", thirdArray)
}

func SliceGo(){
	var firstSlice []int

	for i,j := 0,1 ; i < 2; i,j = i+1, j+1{// i+1 === i++
		firstSlice[i] = j
	}
	
	//
	var secondSlice []string = []string{"Alex", "Osei", "PJ"}

	var firstNames []string = secondSlice //References the same memory location(any changes also changes the original Slice)
	firstNames = append(firstNames, "Yaw")


	//
	thirdSlice := make([]int, 0, 10)
	thirdSlice = append(thirdSlice, 2,4,6,8,10)

	fmt.Printf("firstArray : %v Capacity : %v Size: %v\n", firstSlice, cap(firstSlice), len(firstSlice))
	fmt.Printf("Ref Types :: secondSlice : %v\n", secondSlice)
	fmt.Printf("Ref Types :: firstNames : %v\n", firstNames)
	fmt.Printf("thirdArray : %v Capacity : %v Size: %v\n", thirdSlice, cap(thirdSlice), len(thirdSlice))
}

func DeclareAndPrint(){
	var newCount [8]int

	newCount[0] = 1

	//using make to arrays
	//(type, size, optional = "capacity")
	newSlice := make([]int, 10)
	//doesn't change the size of the array
	newSlice[0] = 5 

	//using append() and indexing
	newSliceTwo := make([]int, 10)
	//doesn't change the size of the array
	newSliceTwo[0] = 12 
	//newSliceTwo[11] = 12 //THIS WILL RESULT IN AN ERROR, Because array size is 10
	newSliceTwo = append(newSliceTwo, 5) //inceases the size of the array by appending to the end of the array

	//creating an empty array with a capacity of 100 with "make"
	newSliceThree := make([]int, 0, 100)//(type, size, optional = "capacity")
	newSliceThree = append(newSliceThree, 5)


	//declare and initialize
	var names [4]string = [4]string{"Alex", "Osei", "Owusu", "PJ"}
	numberSeries := [...]int{1,2,3,4,5,6,7,8}

	

	fmt.Printf("newCount : %v\n", newCount)
	fmt.Printf("new Slice : %v %v\n", newSlice, cap(newSlice))
	fmt.Printf("new SliceTwo : %v %v\n", newSliceTwo, cap(newSliceTwo))
	fmt.Printf("new SliceThree : %v %v\n", newSliceThree, cap(newSliceThree))
	fmt.Printf("One Name : %v\n", string(names[0]))
	fmt.Printf("numberSeries : %v %v", numberSeries, cap(numberSeries))
}