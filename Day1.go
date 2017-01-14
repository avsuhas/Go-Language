/* accepting user first and last name */

package main

import "fmt"

func main() {
    fmt.Print("Enter First_Name: ")
    var First_Name string
    fmt.Scanln(&First_Name)
    fmt.Print(First_Name)
    fmt.Print("Enter Last_Name: ")
    var Last_Name string
    fmt.Scanln(&Last_Name)
    fmt.Print(Last_Name)
    fmt.Print("Enter DateOfBirth in mm/dd/yy format : ")
    
