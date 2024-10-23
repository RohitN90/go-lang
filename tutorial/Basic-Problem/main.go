//Today we are solve some basic problems with go language..
package main

import (
  "fmt"
  "math"
)


func main()  {
  var n int = 10 

  //Print number from 0 to n, where the n is an integer

  for i:=0; i <= n; i++{
    fmt.Printf("%v ", i)
  }

  fmt.Println()

  //Print number for n to 0, where n is an integer

  for i:=n; i >=0; i--{
    fmt.Printf("%v ", i)
  }

  fmt.Println()

  var result string = evenOrOdd(20)
  fmt.Println(result)

  var res string = isPrime(5)
  fmt.Println(res)

  a:= 1
  b:= 2
  fmt.Printf("Before swaping the numbers variable-1 : %v and Variable-2 :%v \n", a, b)
  res1, res2 := swap(a, b)
  fmt.Printf("After swaping the numbers variable-1 : %v and Variable-2 :%v \n", res1, res2)

  var arr []int32 = []int32{1,2,3,4,5}
  var maxValue int32 = max(arr)
  fmt.Println(maxValue)

  var minValue int32 = min(arr)
  fmt.Println(minValue)

}

//write a function which return wheather a number is even or odd
func evenOrOdd(num int) string {
  if num % 2 == 0{
    return "even"
  }
  return "odd"
}

//write a function which return wheather a number is prime or not
func isPrime(num int) string {
  if num < 2{
    return "Not a Prime number"
  }
  if num % 2 == 0{
    return "Not a Prime number"
  }
  for i:=2; i < int(math.Sqrt(float64(num))); i= i + 2{
    if num % i == 0{
      return "Not a Prime number"
    }
  }
  return "It is a Prime number"
}

//write a function which swap two variable with any type
func swap(a any, b any) (any, any){
  var temp any = a
  a = b
  b = temp

  return a , b
}

//write a function which retuen maximum number in an array
func max(arr []int32) int32 {
  var temp int32 = 0
  for i := range arr {
    if arr[i] > temp{
      temp = arr[i]
    }
  }
  return temp
}

//write a function which return minimum number in an array
func min(arr []int32) int32 {
  var temp int32 = math.MaxInt32
  for i := range arr{
    if arr[i] < temp{
      temp = arr[i]
    }
  }
  return temp
}

//So these are some basic or logic building function, where every programmer should know.
//In the next concept with will focus on string manipulation, etc..
//Thank you for Today...
