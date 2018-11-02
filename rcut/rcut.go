package main

import "fmt"
import "strings"
import "os"
import "strconv" 

func main() {
    n, err := strconv.Atoi(os.Args[2])
    if (err == nil) {
        fmt.Printf(strings.Join(strings.Split(os.Args[1], "-")[:n], "-"))
    } else {
        fmt.Printf(err.Error())
    }
}

