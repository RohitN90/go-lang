//In this section we are learn the control flow statement in go language..

//Basically we have only three key words to use the control statement there are "if", "else" and "else if"

package main

import "fmt"

func condition()  {

  //In this example we have taken a variable called number and we are checking that the number is equal to 20 or not..

  var number int = 20

  //This is a normal "If" statement : It is a conditional base operation where we give a certain condition, 
  //it will do the work when the condition become true.
  
  if number == 20 {
    fmt.Println("Same number")
  }

  //We can add "else" if you want to do another operation if the condition failed...

  if number == 10{
    fmt.Println("This is not the number !")
  }else{
    fmt.Println("Yes ! both are same :) ")
  }

  //If you have multipule statement you can use the "else if" key-word...

  if number == 20{
    fmt.Println("Yes ! you got the number")
  }else if number <= 10{
    fmt.Println("The number is lower than the original number")
  }else {
    fmt.Println("The number is higher than the original number")
  }

  //In this way we can use the "if", "else" and "else if" key-word for the control flow statement or conditon based operatinos..
}
