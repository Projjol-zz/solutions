package main

import (
    "fmt"
    "os"
    //"hexToBase64"
)

func main(){
    if len(os.Args) > 1 {
        challengeName := os.Args[1]
        switch challengeName {
        case "hexToBase64":
            if len(os.Args[1:]) == 2 {
                fmt.Println(HexToBase64(os.Args[2]))
            } else {
                fmt.Println("Missing mandatory parameter")
            }
        }
    } else {
        fmt.Println("Missing mandatory parameter")
    }
}
