package main

import "fmt"

var moduler = 2019

func main() {
	var l, r int
	fmt.Scan(&l, &r)
	if r-l >= 2019 {
		fmt.Println(0)
		return
	}
	ans := l * (l + 1) % moduler
	for i := l; i < r; i++ {
		for j := l + 1; j <= r; j++ {
			mod := i * j % moduler
			if mod < ans {
				ans = mod
			}
		}
	}
	fmt.Println(ans)
}
