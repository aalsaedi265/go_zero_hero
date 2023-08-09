
package main

import "fmt"



type contactInfo struct{
	email string
	zipCode int
}
type person struct{
	firstName string
	lastName string
	contact contactInfo
}
func main(){

	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)
	// fmt.Printf("%+v",alex)
	jim:= person{
		firstName: "Jim",
		lastName: "Rage",
		contact: contactInfo{
			email: "jim@alo.com",
			zipCode: 42069,
		},
	}
	// fmt.Printf("%+v", jim)
	// jimPointer :=  &jim
	jim.updateName("jimmy")
	jim.print()
}
func (pointerToPerson *person) updateName( newFirstName string){
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print(){
	fmt.Printf("%+v", p)
}