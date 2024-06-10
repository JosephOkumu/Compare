package main

import "fmt"

func compare(a, b string) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	} else {
		return 0
	}

}

func main() {
	fmt.Println(compare("Hello!", "Hello!"))
	fmt.Println(compare("Salut!", "lut!"))
	fmt.Println(compare("Ola!", "Ol"))
}
