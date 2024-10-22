//Today we are learn about Array and slice in golang...

package main

import "fmt"

func main()  {

  //This is how we declear an array : array-name, [] : length, followed by data-type..
  //The arrays are fixed length and continuous in memory location which help to access the elements very easly..
  var array[3]int32

  //We can add elements into the array by using this syntax.
  array[0] = 12
  array[1] = 132
  array[2] = 90

  //This is how we access the elements in the array..
  fmt.Println(array[0])

  //If you want to select the values in a particula range we use this type of syntax
  fmt.Println(array[0:3])

  //if you want to print array with the loop, this is the syntax
  for _, v := range array {
    fmt.Printf("%v ", v)
  }

  fmt.Println()

  //we are using a special print function "fmt.Printf()", where if we want to combine the text with value we use this "printf()" function.
  fmt.Printf("%v \n", array[2])

  //we can also write in classic way
  for i:=0; i < len(array); i++ {
    fmt.Printf("%v ", array[i])
  }

  //If you want to initial array with respective value we can write as..
  var newArray[5]int16 = [5]int16{1,2,3,4,5}
  fmt.Println(newArray)

  //A slice is also an array with some additional functionalities.
  //By removing the length between the square brackets we can create a slice
  var intSlice[]int32 = []int32{1,2,3}

  fmt.Printf("the length of slice is %v and the capacity of slice is %v \n", len(intSlice), cap(intSlice))
  fmt.Println(intSlice)

  //we can add element into the slice by using "append" fucntion, but i will re-allocate a new slice and copy the old slice with exsisting array..
  //If we see the length and capacity of intSlice before and after adding the element we can see the different..
  intSlice = append(intSlice, 4)
  fmt.Printf("the length of array is %v and capacity of slice is %v \n", len(intSlice), cap(intSlice))
  fmt.Println(intSlice)

  //We can create a slice using "make" function
  var arrSlice []int32 = make([]int32, 5)
  fmt.Println("creating a slice with make function !")
  fmt.Println(arrSlice)

  //This is all about array and slices today..
}
