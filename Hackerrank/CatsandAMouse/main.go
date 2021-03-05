package main

import (
	"fmt"
	"math"
)

func main() {
	var x,y,z float64
	var n,i int32
	fmt.Scan(&n)
	fmt.Scan(&x,&y,&z)
	for i= 0; i< n; i++{
		if math.Abs(x-z) > math.Abs(y-z){
			fmt.Println("CatB")
		}else if math.Abs(x-z) < math.Abs(y-z){
			fmt.Println("CatA")
		}else if math.Abs(x-z) == math.Abs(y-z){
			fmt.Println("MouseC")
		}

	}


}