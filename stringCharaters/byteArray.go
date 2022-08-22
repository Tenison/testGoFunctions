package stringCharaters

import "fmt"

func ConvertStringToByteArray (input string){
	var series[] byte = []byte(input)

	singleCharater := string(series[0])
	
	fmt.Println(singleCharater)
}