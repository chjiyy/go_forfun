package main

import (
    "fmt"
	"os"
    "github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	// read parameters from stdin
	if os.Args == nil || len(os.Args) == 0 {
            fmt.Println("need parameters: filepath!")
            return
    }
    filePath := os.Args[1]
    fmt.Println("Excel file is generating:", filePath)
    
    f := excelize.NewFile()
    f.SetCellValue("sheet1","A1","ok")
    f.SetCellValue("sheet1","A2","okok")
    
    err := f.SaveAs(filePath)
    if err != nil {
        fmt.Println(err)
    }
}
