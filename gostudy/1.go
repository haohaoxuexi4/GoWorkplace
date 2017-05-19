package main

import (
"fmt"
"os"
)
func main(){
	s:=","
	fmt.Println("helloworkd\n")
	for _,arg:=range os.Args[1:]{
		s+=arg
		fmt.Println(arg)
		arg=""
	}
	fmt.Println(s)
}