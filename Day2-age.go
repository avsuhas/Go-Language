/* without validation trial code --- Not working completelty */


package main

import (
        "fmt"
        
)
    func findage(d1, m1, y1, d2, m2, y2 int) {
    var y,d,m int
    y=y2-y1
    m=m2-m1
    d=d2-d1
   
    if (m2<m1) || (m1==m2) && (d2<d1){  
    y=y-1
    }
    
    m = (12-m1)+m2
    if d2<d1{
    d=d*-1
    }
    fmt.Print(" %d years,%d months and %d days",y,m,d);
    
    
}



func main() {
    fmt.Print("Enter First_Name: ")
    var First_Name string
    fmt.Scanln(&First_Name)
    fmt.Println("---------------------")
    fmt.Print(First_Name)
    
    fmt.Print("Enter Last_Name: ")
    var Last_Name string
    fmt.Scanln(&Last_Name)
    fmt.Println("---------------------")
    fmt.Print(Last_Name)
    

    fmt.Print("Enter Date of Birth: ")
    fmt.Print("day :")
    var d1 int
    fmt.Scanln(&d1)

    fmt.Print("month :")
    var m1 int
    fmt.Scanln(&m1)

    fmt.Print("year :")
    var y1 int
    fmt.Scanln(&y1)

    fmt.Print("Enter the Present date :")
    fmt.Print("day :")
    var d2 int
    fmt.Scanln(&d2)

    fmt.Print("month :")
    var m2 int
    fmt.Scanln(&m2)

    fmt.Print("year :")
    var y2 int
    fmt.Scanln(&y2)


    findage(d1,m1,y1,d2,m2,y2)
   

}
