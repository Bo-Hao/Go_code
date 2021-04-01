// Golang program to show how
// to use command-line arguments
package main
  
import (
    "fmt"
    "os"
)
  
func main() {
  
    // The first argument
    // is always program name
    myProgramName := os.Args[1]
      
    // it will display 
    // the program name
    fmt.Println(myProgramName)

	
}