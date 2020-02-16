package main

import (
	"fmt"
	"time"
)

func main() {
	m := time.Now().Format("01-JAN-2006 15:04")
	time.Sleep(time.Minute)
	if m < time.Now().Format("01-JAN-2006 15:04") {
		fmt.Println("Successful")
	}
	fmt.Println(m)
}
