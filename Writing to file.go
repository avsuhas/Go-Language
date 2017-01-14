/*  writing a file to disk */

package main

import (
        "fmt"
        "time"
      )


	file, fileErr := os.Create("file")
	if fileErr != nil {
    fmt.Println(fileErr)
    return
	}
	//fmt.Fprintf(file)
	fmt.Fprintf(file, "%v years\n %v month\n %v days\n %v hours\n %v mins \n %v seconds \n", year,month,day,hour, min, sec)  
