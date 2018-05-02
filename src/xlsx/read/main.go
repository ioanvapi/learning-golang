package main

import (
    "fmt"
    "github.com/tealeg/xlsx"
)

func main() {
    excelFileName := "eos_b1.xlsx"
    xlFile, err := xlsx.OpenFile(excelFileName)
    if err != nil {
        fmt.Println(err)
        return
    }
    for _, sheet := range xlFile.Sheets {
        for _, row := range sheet.Rows {
            for _, cell := range row.Cells {
                text := cell.String()
                fmt.Printf("%s\n", text)
            }
            fmt.Println("===============")
        }
    }
}
