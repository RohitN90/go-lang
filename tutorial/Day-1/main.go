
//This is the Day-One of learning go-lang basic in a 10 Days

//This is the main pacakge where the go compiler look for the main function to run.
package main

//The "fmt" means format lib or module which help to print the text in the console..
//Till now we know that the "fmt" is used to Print the text only but there are so many way we can use the 'format' module or lib

import "fmt"

//This is the main function where the execution start

func main() {

  //It will print "Hello "

  fmt.Println("Hello ")

//----------------------------------------------what i have learn-------------------------------------------------
  
  //To day i have learn variables, different data types and the different between var and const

  //The syntax of writing the variable are two different ways : normal or short sign
  
  //Normal : (var or const) variable-name data-type = value
  var name string = "Hello Vardhan"

  //Short sign : variable-name := value
  age1 := 49

  fmt.Println(name)
  fmt.Println(age1)


  //---------------------------------------------------Data-Types----------------------------------------------------

  //These are different data-type like int, float32, float64, bool, string, uint

  //These is no different between the int and uint.
  //The int can have both positive and negative integer values but in uint it only have the positive integer only
  //The both int and uint contains the 8bit, 16bit, 32bit and 64bit. 
  
  var age int = -20
  var newAge uint = 20

   fmt.Printf("%v %v", age, newAge)


  //This will store the collection of character to form a string
  var fullname string = "ROhit vardhan"
  fmt.Println(fullname)

  //This is a boolearn data type where it return "True" or "False"
  var isLogin bool = false
  fmt.Println(isLogin)

  //This will store the decimal values
  var money float32 = 26.526
  fmt.Println(money)

  //The different between var and const is the var variable can be modified but the const variable can't be change or modified
  // For example :

  var car string = "BMW"
  
  fmt.Println(car)
  car = "Audi"

  fmt.Println(car)

  //but we can't do with const variable

  const pie float32 = 3.14
  fmt.Println(pie)

  //It will tell you that you can't change the variable onces it is initilized as a const variables 
  
  // " pie = 20.90 "

  //At the end of the Day i have a good understanding on declearing a variables and their data-types and different between var and const...

}

