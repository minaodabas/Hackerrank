package main
import "fmt"

func main(){
	var h,m,s int
	var n string

	fmt.Scanf("%d:%d:%d%s",&h,&m,&s,&n)

	if n == "AM" && h == 12 {
		h = 0
	}
	if n == "PM" && h != 12 {
		h += 12
	}
	fmt.Printf("%02d:%02d:%02d",h,m,s)

}
