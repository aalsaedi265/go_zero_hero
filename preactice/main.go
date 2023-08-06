
package main

func main(){
	numbers := createSlice()
    for _, num := range numbers {
        checkEvenOrOdd(num)
    }
}