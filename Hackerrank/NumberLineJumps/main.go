package main

import "fmt"


func main() {
	var x1,v1,x2,v2 int32
	fmt.Scan(&x1,&v1,&x2,&v2)

	if (v1>v2 && x2>x1 && (x2-x1)%(v1-v2) ==0){
		fmt.Println("YES")
	}else{
		fmt.Println("NO")
	}


}
