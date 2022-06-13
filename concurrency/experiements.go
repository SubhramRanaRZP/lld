package main

import "fmt"

func main(){
	a,b :=  1,2
	a , b = a+b, a
	fmt.Println(a, b)
}