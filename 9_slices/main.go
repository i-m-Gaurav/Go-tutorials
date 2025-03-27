package main

import (
	"fmt"
	
	"slices"
)

func main(){
	// slices are like array, but it automatically expands its size,
	// its like vector in the c++,


	// this is the uninitialized slice, initially it is nill by default.
	// unlike array, which is 0 at start.

	var nums []int

	fmt.Println(nums)
	// initially it will give the size of 0
	fmt.Println(len(nums))

	// now if we don't want our slice initially to become nill 
	// then we can use something called make.

	var num = make([]int ,2)

	fmt.Println(num)
	// in this case we have size 2 for this slice
	fmt.Println(len(num))
	// here one more thing is called
	// capacity 

	// some more make example,

	// so this is not nill, because we have used the make
	// function so now its has size 5, and type int.


	// so by default capacity is same as the length of the array,
	// but if we think our slice is going to grow in the future,
	// then we can pass a capacity parameter in the make function.


	var num2 = make([]int, 0, 5)

	// fmt.Println("Length is",len(num2))
	// fmt.Println("Capacity is",cap(num2))

	num2 = append(num2, 4);
	num2 = append(num2, 4);
	num2 = append(num2, 4);
	num2 = append(num2, 4);
	num2 = append(num2, 4);
	num2 = append(num2, 4);
	num2 = append(num2, 4);
	num2 = append(num2, 4);
	num2 = append(num2, 4);
	num2 = append(num2, 4);
	num2 = append(num2, 4);
	

	fmt.Println("Length is",len(num2))
	fmt.Println("Capacity is",cap(num2))


	fmt.Println(num2)

	// another way to make slice

	num3 := []int{}

	num3 = append(num3, 1)

	fmt.Println(len(num3))
	fmt.Println(cap(num3))
	fmt.Println(num3)

	// since we are inserting the element using the append,
	// now we can insert using the indices too.

	 var num4 = make([]int, 2, 5)

	 num4[0] = 3
	 num4[1] = 5
	
	fmt.Println(num4)

	// we should use the append, to be sane.

	// now lets copy 

	var s1 = make([]int , 0 , 5)

	s1 = append(s1, 6)
	var s2 = make([]int , len(s1))

	copy(s2 , s1)

	fmt.Println(s1);
	fmt.Println(s2)


	// now lets copy a slice of fruits.

	var fruits = make([]string , 0 ,5)

	fruits = append(fruits, "Mango")
	fruits = append(fruits, "Apple")
	fruits = append(fruits, "Banana")
	fruits = append(fruits, "Strawberry")
	fruits = append(fruits, "Watermelon")



	var better_fruits = make([]string, len(fruits))

	copy(better_fruits , fruits)

	fmt.Println("Fruits are copied to the better fruits slice :- ",better_fruits);

	// lets copy some floats slices

	var g1 = make([]float64, 0)

	g1 = append(g1, 5.3)
	g1 = append(g1,4.5)
	g1 = append(g1, 25.4)
	g1 = append(g1, 3.34)
	g1 = append(g1, 54.322)

	var g2 = make([]float64, len(g1))

	copy(g2, g1)
	 
	fmt.Println("This is the g2 :- ",g2);

	// slice operator 

	var h1 = []int{1,2,3,4,5,6}

	fmt.Println(h1[:3])

	// Now lets understand about the slice operator
	// it is just colon sign
	// and we use it in the pace of the index like this 
	// h1[0:2]
	// here 0 means from where it will start, okay 
	// and 2 means to where it will go, so in our case, 
	// it will start from  0 and go to the 2nd index, but but but,
	// it will exclude that index, so it will give result of the index
	// 0 & 1 and exclude the index 2, got it pretty simple right.

	// there is a thing called ignore,
	// we can just skip that initial 0 that we are putting,
	// h1[:2] here we have not written the 0 so that mean it will by
	// default start from the 0

	// but if we skip the 2nd arg like this h1[0:] then in this case it will 
	// capture all the elements present int he slice. 
	// so it will give the result for the h1 slice ,
	// 1 ,2 , 3, 4, 5, 6 i mean all the element got it,
	// so here if we put 2 instead of 0 then it will start from the 2nd index, 
	// and capture rest of the elements.

	fmt.Println(h1[2:])


	// there is one slice package available in the go too.

	// it has some pre-built functions. that we can use.

	var t1 = []int {1 ,2,3}
	var t2 = []int {1 ,2,3,5}

	var isEqual = slices.Equal(t1,t2);
	fmt.Println(isEqual)

	// now lets create 2d slices, i don't know where the hell we will
	// use this.

	var dd = [][]int{{1,2},{2,3}}

	fmt.Println(dd)








}