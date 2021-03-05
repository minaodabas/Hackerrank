package main

import "fmt"


func main() {
	var n,t,i int32
	fmt.Scan(&n)
	arr := make([]int32,5)
	for i=0 ;i<n;i++{
		fmt.Scan(&t)
		t-=1
		arr[t]++

	}
	max,occ := arr[0],0

	for i,num := range arr{
		if num > max{
			max = num
			occ = i+1
		}
	}
	fmt.Println(occ)
}

