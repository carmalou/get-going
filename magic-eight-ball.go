package main

import(
  "fmt"
  "os"
  "math/rand"
)

func main() {
  if len(os.Args) > 1 {
    // return something from the array
    fortune := getFortune()
    fmt.Println(fortune);
  } else {
    fmt.Println("Ask the question that is in your heart.")
  }
}

func getFortune() string {
  fortunesArr := []string{"It is certain",
    "It is decidedly so",
    "Without a doubt",
    "Yes, definitely",
    "You may rely on it",
    "As I see it, yes",
    "Most likely",
    "Outlook good",
    "Yes",
    "Signs point to yes",
    "Reply hazy try again",
    "Ask again later",
    "Better not tell you now",
    "Cannot predict now",
    "Concentrate and ask again",
    "Don't count on it",
    "My reply is no",
    "My sources say no",
    "Outlook not so good",
    "Very doubtful",
  }
  n := rand.Int() % len(fortunesArr)
  fmt.Println(n);
  return fortunesArr[n]
}
