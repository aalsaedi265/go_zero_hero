package main
import "fmt"

func createSlice() []int{
	var numbers []int
	for i:=0; i<=10; i++{
		numbers= append(numbers, i)
	}
	return numbers
}
func checkEvenOrOdd(num int) {
    if num%2 == 0 {
        fmt.Printf("%d is even\n", num)
    } else {
        fmt.Printf("%d is odd\n", num)
    }
}