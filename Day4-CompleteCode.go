/* working code */
// code before tracking back at 10:23 morning

package main

import (
        "fmt"
        "time"
        "os"
        )

func diff(a, b time.Time) (year, month, day int) {

    y1, M1, d1 := a.Date()
    y2, M2, d2 := b.Date()

    year = int(y2 - y1)
    month = int(M2 - M1)
    day = int(d2 - d1)

    // Normalize negative values
    
    if day < 0 {

    // days in month:
    t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
    day += 32 - t.Day()
    month--
    }

    if month < 0 {
        month += 12
        year--
    }

    return
}

    func main() {

    var First_Name string    // Variable Declaration
    var Last_Name string

    fmt.Print("Enter First_Name: ")    //Taking user input 
    fmt.Scanln(&First_Name)              //Storing in Declared Variables    

    fmt.Print("Enter Last_Name: ")
    fmt.Scanln(&Last_Name)
    
    fmt.Println("Entered First and Last Name are:  ")  // Displaying the user info
    fmt.Println(First_Name)             
    fmt.Println(Last_Name)
    fmt.Println("---------------------")

    // Variable Declaration 
    var d1 int                  // for convinence day is named as d1
    var m1 int                  // for convinence month is named as m1
    var y1 int                  // for convinence year is named as y1

    // using built in types to compare current mm/dd/yyyy values
    now := time.Now() 
    currentYear := now.Year()
    currentMonth := now.Month()
    currentDay := now.Day()

    // Accepting User Date of Birth
    fmt.Println("Enter the Birth Year in yyyy format :\n")
    fmt.Scanln(&y1)
    for y1 < 1000 || y1 > 9999{                     // validating year to accept 4 digit entry
        fmt.Println("Date should be in yyyy fornat\n")
         fmt.Scanln(&y1)

    } 

    for y1 > currentYear {      // validating year input to check 
                                // if it is greater than current year
    fmt.Println("Date of Birth cannot be greater than current year\n")
    fmt.Println("please re-enter the correct year\n")
    fmt.Scanln(&y1)
    }

    // Month input should be between 1 to 12 
    // Validating Month input to handle errors 
    // Validation check use ( go ) while loop
    fmt.Println("Enter the Birth Month\n")
    fmt.Scanln(&m1)
    
    for y1 == currentYear && time.Month(m1) > currentMonth  {

    fmt.Println("Entered month is greater than current month, Please re-enter valid month\n")
    fmt.Scanln(&m1)
    }
     
    for m1 < 1 || m1 > 12{
        if m1 < 1{
    fmt.Println("Month shouldnot be less than 1. please enter the correct month range in b/w 1 to 12\n")
    }else {
    fmt.Println("month is greater than current month please enter valid month\n")
    }
    fmt.Scanln(&m1)
    }

    // Day input should be between 1 to 12 
    // Validating Day input to handle errors 
    // Validation check use ( go ) while loop
    fmt.Println("Enter the Birth day")
    fmt.Scanln(&d1)

    for y1 == currentYear && time.Month(m1) == currentMonth && d1 > currentDay  {

    fmt.Println("Entered day is greater than current day please enter valid day")
    fmt.Scanln(&d1)
    }

    for d1 > 31 || d1 < 1{
        if d1 > 31 {
        fmt.Println("please enter the correct day range in b/w 1 to 31")
        }else{
         fmt.Println("Day shouldnot be less than 1. please enter the correct day range in b/w 1 to 31")   
        }
    fmt.Scanln(&d1)
    }

//end of while

// Converting integer to month time.Month(m1)

    birthday := time.Date(y1, time.Month(m1), d1, 3, 30, 0, 0, time.UTC)

    year, month, day := diff(birthday, now)

    fmt.Printf("You are %d years, %d months, %d days old.\n",year, month, day)
    fmt.Println("Thank you for using Age Caluculator !!\n")

// File Writing code Starts from Here

    file, fileErr := os.Create("file.txt") // created a file name "file.txt"
    if fileErr != nil {
    fmt.Println(fileErr)
    return
    }
    fmt.Fprintf(file,"First_Name : %v \nLast_Name : %v\n",First_Name,Last_Name)
    fmt.Fprintf(file,"The Age of the user is :")
    fmt.Fprintf(file," %v years\n %v month\n %v days \n", year,month,day)  
}
