package main

import "fmt"

func main() {
	var n string
	fmt.Scan(&n)
	if n[0] == n[1] && n[1] == n[2] && n[2] == n[3] {
		fmt.Println("SAME")
	} else {
		fmt.Println("DIFFERENT")
	}
}
