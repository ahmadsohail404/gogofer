package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) // default golang date format
	// go be like: hamare yahan yehi date hota hai

	createdDate := time.Date(2020, time.December, 23, 4, 24, 12, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))

}

// for building executable files for OS

// GOOS="windows" go build
// GOOS="darwin" go build
// GOOS="linux" go build
