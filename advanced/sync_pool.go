package main

import (
	"fmt"
	"sync"
)

type person struct {
	name string
	age  int
}

func sync_pool_main() {
	var pool = sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating a new Person.")
			return &person{}
		},
	}

	pool.Put(&person{
		name: "Dushyanth",
		age:  26,
	})

	pool.Put(&person{
		name: "Dhana",
		age:  46,
	})

	pool.Put(&person{
		name: "Chaithanya",
		age:  28,
	})

	person1 := pool.Get().(*person)
	person2 := pool.Get().(*person)
	person3 := pool.Get().(*person)

	fmt.Println(person1)
	fmt.Println(person2)
	fmt.Println(person3)

	person4 := pool.Get().(*person)
	fmt.Println(person4)
}
