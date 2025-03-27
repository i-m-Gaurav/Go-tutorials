package main

import (
	"fmt"
	"maps"
)

func main(){
	//create a map
	// its similar to the objects, dictionaries form other languages


	// here the first string is the type of the key and second string is the 
	// type of the value.
	map1 := make(map[string]string)

	// setting an element
	// if key value does not exist in the map
	// then it return the zero value.

	map1["Name"] = "Golang"
	map1["area"] = "backend"	

	fmt.Println(map1)

	// it will return empty value because, it does not exist.
	fmt.Println(map1["hone"])


	map2 := make(map[string]int)

	map2["age"] = 22

	fmt.Println(map2)
	fmt.Println(len(map2))

	// lets delete one element from map

	//delete(map2,"age")

	fmt.Println(map2)

	clear(map2)

	fmt.Println(map2)

	//initialize the value at the time of declaring

	map3 := map[string]int {"price" : 40 , "phones" : 3}

	fmt.Println(map3)

	// check if some value is there,

	 v, ok := map3["price"]

	 if ok {
		fmt.Println(v)
		fmt.Println("its in the map")
	 } else {
		fmt.Println(v)
		fmt.Println("Not there")
	 }

	 map4 := map[string]int {"price" : 40, "phones" : 3}

	// check if two maps are equal

	fmt.Println(maps.Equal(map3,map4))




}