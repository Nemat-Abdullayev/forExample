package main1

import "fmt"

func main() {
	fmt.Println("Hello MAN?")
	foo()
	fmt.Println("this is me!")

	for i := 2; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}

	}
	bar()
}
func foo() {
	fmt.Println("I'am in foo")
}

func bar()  {
	fmt.Println("and the we exited")
}
