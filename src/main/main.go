package main

import (
	"fmt"
	"time"
	"net/http"
	"github.com/360EntSecGroup-Skylar/excelize"
)

func main(){
	fmt.Println("Google URL Checker")
	
  //Excel file here
	xlsx, err := excelize.OpenFile("google-url-checker.xlsx")

	if err != nil {
		fmt.Println(err)
		return
	}

	rows := xlsx.GetRows("Sheet1")
	t1 := time.Now()
	fmt.Println(t1)
	for _, row := range rows {
		fmt.Println(row[0])
		if(row[0]=="Top pages"){
			continue
		}
		resp, err := http.Get(row[0])
	    if err != nil {
	        fmt.Println(err)
	    }
	    // Print the HTTP Status Code and Status Name
	    fmt.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))
		
	}
	t2 := time.Now()
	fmt.Println(t2)
	
	diff := t2.Sub(t1)
	fmt.Println(diff)
	
	out := time.Time{}.Add(diff)
	fmt.Println(out.Format("15:04:05"))
	
	//http network GET with HTTP Status Code Return
}
