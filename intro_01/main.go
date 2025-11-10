package main

import "fmt"
import "github.com/fatih/color"

func main() {
	fmt.Println("Hello, World")
    color.Red("Hello, World")

    //strings
    text := `open
    \n
    close
    `
    fmt.Println(text)
    isAdult(18)
}


func isAdult (age int) {
    if age >= 18 {
        fmt.Println("Adult")
    }else{
        fmt.Println("Teen")
    }
}