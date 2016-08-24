package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go is running now  on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	case "freebsd":
		fmt.Println("OpenBSD")
	default:
		fmt.Printf("%s.", os)
	}

}
