package putil

import (
	"fmt"
	"testing"
)

type Person struct {
	name string
}

func modifyArray(array *[]Person) {
	//*array = array[0:2]
	(*array)[0].name = "33"
	*array = append(*array, Person{
		"4",
	})
	fmt.Printf("%v\n", array)

}

func TestArrayFunc(t *testing.T) {

	var ss = []Person{
		Person{
			"sdfd",
		}, Person{
			"2",
		}, Person{
			"3",
		}}

	modifyArray(&ss)
	//ss = ss[0:2]
	fmt.Printf("%v\n", ss)
}
