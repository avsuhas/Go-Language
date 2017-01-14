/* not a clean code */
/* code optimization, code clean up need to be done */
/* complete code for calculating age and writing it on to a file */

package main

import (
	"fmt"
 	"time"
 	"os"
	)

func diff(a, b time.Time) (year, month, day, hour, min, sec int) {

    y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

    h1, m1, s1 := a.Clock()
	h2, m2, s2 := b.Clock()

    year = int(y2 - y1)
	month = int(M2 - M1)
	day = int(d2 - d1)

    hour = int(h2 - h1)
	min = int(m2 - m1)
	sec = int(s2 - s1)

    // Normalize negative values

    if sec < 0 {

        sec += 60

        min--

    }

    if min < 0 {

        min += 60

        hour--

    }

    if hour < 0 {

        hour += 24

        day--

    }

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

fmt.Print("Enter First_Name: \n")
    var First_Name string
    fmt.Scanln(&First_Name)
    fmt.Println("---------------------")
    fmt.Print(First_Name)
    
    fmt.Print("Enter Last_Name: \n")
    var Last_Name string
    fmt.Scanln(&Last_Name)
    fmt.Println("---------------------")
    fmt.Print(Last_Name)

var d1 int

var m1 int

var y1 int

// Your birthday: let's say it's January 2nd, 1980, 3:30 AM

now := time.Now() 
currentYear := now.Year()
currentMonth := now.Month()
currentDay := now.Day()


 fmt.Println("Enter the Birth Year")

 fmt.Scanln(&y1)

for y1 > currentYear {
//DOB should not  be greater than current 
  fmt.Println("Your yet to born - Invalid Year Entered - please enter Correct ")

  fmt.Scanln(&y1)

 }

 fmt.Println("Enter the Birth Month")

 fmt.Scanln(&m1)
	
for y1 == currentYear && time.Month(m1) > currentMonth  {

  fmt.Println("month is greater than current month please enter valid month")
fmt.Scanln(&m1)

	}
 for m1 > 12 {

  fmt.Println("please enter the correct month range in b/w 1 to 12")

  fmt.Scanln(&m1)

 // if currentYear  && (m1 > currentMonth){

 // fmt.Println("Invalid Month entered - Enter correct")

 // }
  // add the line to for user entry

 }

 for m1 < 1 {

  fmt.Println("Month shouldnot be less than 1. please enter the correct day range in b/w 1 to 31")

  fmt.Scanln(&m1)


 }

 fmt.Println("Enter the Birth day")

 fmt.Scanln(&d1)


 for y1 == currentYear && time.Month(m1) == currentMonth && d1 > currentDay  {

  fmt.Println("day is greater than current day please enter valid day")
	fmt.Scanln(&d1)

	}

 for d1 > 31 {

  fmt.Println("please enter the correct day range in b/w 1 to 12")

  fmt.Scanln(&d1)

 }

 for d1 < 1 {

  fmt.Println("Day shouldnot be less than 1. please enter the correct month range in b/w 1 to 12")

  fmt.Scanln(&d1)



 }



//validation check use while loop

//d1 = 32

 fmt.Println(m1)

//end of while

// Converting integer to month time.Month(m1)

birthday := time.Date(y1, time.Month(m1), d1, 3, 30, 0, 0, time.UTC)

year, month, day, hour, min, sec := diff(birthday, now)




fmt.Printf("You are %d years, %d months, %d days, %d hours, %d mins and %d seconds old.",

    year, month, day, hour, min, sec)

// Writing code Starts from Here

	file, fileErr := os.Create("file")
	if fileErr != nil {
    fmt.Println(fileErr)
    return
	}
	//fmt.Fprintf(file)
	fmt.Fprintf(file, "%v years\n %v month\n %v days\n %v hours\n %v mins \n %v seconds \n", year,month,day,hour, min, sec)  



}
