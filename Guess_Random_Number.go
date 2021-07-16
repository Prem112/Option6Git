// Guessing Game
// File name: Guess_Random_Number.go
// Author: Natchathiram Premkumar
// Date: 16/07/2021, Initial version

package main
 
import(
   "fmt"
   "math/rand"
   "time"
)
 
func xrand(min, max int) int {
   rand.Seed(time.Now().Unix())
   return rand.Intn(max - min) + min
}
 
func main() {
   mynumber := xrand(1, 100)
   fmt.Println("You guessed the number in %d tries\n", mynumber)
   attempts := 0
   left := 10
   var guess int
   fmt.Println("Guess a number between 1 and 100.")
 
   for attempts < 10 {
	  attempts++
	  left--
      fmt.Println("Enter your guess of what the target number is:")
      fmt.Scanf("%d", &guess)

      if guess < mynumber {
         fmt.Println("Oops, Your guess was LOW.")
		 fmt.Println("You have only left", left, "Chances")
      } else if guess > mynumber {
         fmt.Println("Oops, Your guess was HIGH.")
		 fmt.Println("You have only left", left, "Chances")
      } else if guess == mynumber {
         break
      }
   }
   if guess == mynumber {
      fmt.Printf("Good Job! You guessed it!")
   } else {
      fmt.Printf("Sorry, You didn't guess my number. It was %d", mynumber)
   }
}