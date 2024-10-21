//Today we are learning about function in golang with and without parameters...
package main

import "fmt"


func main()  {
   //Function without parameters and without return value
  printInfo()

  //Function with parameters ans without return value
  printName("rohit")

  //Function with parameters ans with return value
  result := highOrLow(20, 10)
  fmt.Println(result)
}

//This funciton will print the message "Hello ROhit" in the console or terminal
//As you can see this funciton is not taking any parameter.
func printInfo()  {
  fmt.Println("Hello Rohit")
}


//This function will print the message "My name is <Your name>" where the name variable is pass into the function.
func printName(name string)  {
  fmt.Printf("My name is %v\n", name)
}


//This is a function which return which number is greater ot not..
//This function takes parameters and return the parameter also...
func highOrLow(num1 int, num2 int) int {
  if num1 > num2 {
    return num1
  }
  return num2
}

//That all about functions with parameter and without parameter.
//Sorry for the two days gap.. i went to home 😅 and i didn't take my laptop there..


