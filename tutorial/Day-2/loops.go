
//This is the Day-2 of learning Go language.

//Today we are going to learn loops and control statements.

package main

import "fmt"

func main()  {

  //Loops : It is a continues process of a particular taks...
  //We have for loop, while loop...
  // But in Go language we can't use the while but instead we can use "for" key word to do the operation like while loop
  
  for i := 0; i < 10; i++ {
    fmt.Print(i," ")
  }

  fmt.Println()

  //This ^^^^ is the basix version of writing a loop in go. 
  //But most of the time we are using the loop with "range" keyword

  for i := range 10 {
    fmt.Print(i," ")
  }

  //As I told you while loop can be achive by using "for" key word

  fmt.Println()

  var count = 0

  for count < 10 {
    fmt.Print(count)
    count = count + 1
  }

  //This is how the while loop and for loop works in go language..
  
}
