package main


import "fmt"
import "os"


func main (){

	for i:= 1; i<len(os.Args);i++{

		fmt.Printf("%v",os.Args[i])
		fmt.Printf("\n")

	}
	


}
