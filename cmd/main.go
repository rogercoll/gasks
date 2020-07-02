package main

import (
	"os"
	"fmt"
	"github.com/rogercoll/gasks"
)


func main() {
        arguments := os.Args
        if len(arguments) == 1 {
                fmt.Println("Please provide a host:port string")
                return
        }
        CONNECT := arguments[1]
	gasks.Start(CONNECT)
}
