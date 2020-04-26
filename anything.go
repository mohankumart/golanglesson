package main
import "fmt"

func main() {
	fmt.Println("Hello world!!!")
	foo()
	fmt.Println("Testing came back")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println("divide by 2")
		}
	}
	fmt.Println("exited")
}

func foo() {
	fmt.Println("I am in foo")
}





