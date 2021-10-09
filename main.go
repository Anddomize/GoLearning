package main

import (
	"./alertmodule"
	"fmt"
)

func main() {
	fmt.Println("Hello")
	alertmodule.Alert_send()
}
