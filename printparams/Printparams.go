package main


import "fmt"
import "os"


func main (){

	for i:= 0; i<len(os.Args);i++{

		fmt.Printf("%v",os.Args[i])

	}
	


}
