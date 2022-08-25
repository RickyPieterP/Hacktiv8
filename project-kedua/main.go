package main

import "fmt"

type Person struct {
	Name string
}

func main() {

	var person = []*Person{
		{Name: "nama1"},
		{Name: "nama2"},
		{Name: "nama3"},
		{Name: "nama4"},
		{Name: "nama5"},
		{Name: "nama6"},
		{Name: "nama7"},
		{Name: "nama8"},
		{Name: "nama9"},
		{Name: "nama10"},
	}

	printEmployee := func(in []*Person) {
		for _, v := range in {
			fmt.Println(v.Name)
		}
	}

	printEmployee(person)

}
