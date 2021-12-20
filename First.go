  package main

  import (
	  "bufio"
	  "fmt"
	  "log"
	  "math/rand"
	  "os"
	  "strconv"
	  "strings"
	  "time"
  )
  func main() {
	  isPlayerWin := false
	  var status string
	  rand.Seed(time.Now().UnixNano())
	  randomNumber := rand.Int63n(100) + 1
	  for i := 1; i <= 10; i++ {
		  fmt.Print("Попробуйте угадать число: ")
		  reader := bufio.NewReader(os.Stdin)
		  inputNumberStr, err := reader.ReadString('\n')

		  if err != nil {
			  log.Fatal(err)
		  }

		  inputNumberStr = strings.TrimSpace(inputNumberStr)
		  inputNumberInt, err := strconv.ParseInt(inputNumberStr, 10 , 64)

		  if err != nil {
			  log.Fatal(err)
		  }

		  if inputNumberInt != randomNumber {
			  status = "low"
			  fmt.Println("Oops. Yoor guess was", status, "You have left", 10-i, "attempt")
		  } else {
			  status = "high"
			  fmt.Println("Oops. Yoor guess was", status)
			  isPlayerWin = true
			  break
		  }
	  }
	  if !isPlayerWin {
		  fmt.Println("Sorry. You didn’t guess my number. It was:", randomNumber)
	  }
  }

