package main

import (
	"fmt"
	"time"
	"math"	
  )

// Define eatTacos() here:
func eatTacos() {
  fmt.Println("Yum!")
}

func startGame() {
  instructions := "Press enter to start..." 
  
  fmt.Println(instructions)
  
}


func isItLateInNewYork() string {
	var lateMessage string
	t := time.Now()
	tz, _ := time.LoadLocation("America/New_York")
	nyHour := t.In(tz).Hour()
	if nyHour < 5 {
	  lateMessage = "Goodness it is late"
	} else if nyHour < 16 {
	  lateMessage = "It's not late at all!"
	} else if nyHour < 19 {
	  lateMessage = "I guess it's getting kind of late"
	} else {
	  lateMessage = "It's late"
	}
	
	// Return the string lateMessage
	return lateMessage
	
}


func computeMarsYears(earthYears int) int {

earthDays := earthYears * 365
marsYears := earthDays / 687
return marsYears
}


func specialComputation(x float64) float64 {
	return math.Log2(math.Sqrt(math.Tan(x)))
}


func getLikesAndShares(postId int) (int, int) {
	var likesForPost, sharesForPost int
	switch postId {
	case 1:
	  likesForPost = 5
	  sharesForPost = 7
	case 2:
	  likesForPost = 3
	  sharesForPost = 11
	case 3:
	  likesForPost = 22
	  sharesForPost = 1
	case 4:
	  likesForPost = 7
	  sharesForPost = 9
	}
	fmt.Println("Likes: ", likesForPost, "Shares: ", sharesForPost)
	// Add in a return for likesForPost and sharesForPost
	return likesForPost, sharesForPost
}


func queryDatabase(query string) string {
	var result string
	connectDatabase()
	// Add deferred call to disconnectDatabase() here
	defer disconnectDatabase()
	if query == "SELECT * FROM coolTable;" {
	  result = "NAME|DOB\nVincent Van Gogh|March 30, 1853"
	}  
	fmt.Println(result)
	return result
}

func connectDatabase() {
	fmt.Println("Connecting to the database.")
}

func disconnectDatabase() {
	fmt.Println("Disconnecting from the database.")
}


func main() {
  // Call your function here:
  eatTacos()
  startGame()

  var nyLate string
  nyLate = isItLateInNewYork()
  fmt.Println(nyLate)


  myAge := 43
  myMartianAge := computeMarsYears(myAge)
  fmt.Println(myMartianAge)


  var a, b, c, d float64
  a = .0214
  b = 1.02
  c = 0.312
  d = 4.001
  
  a = specialComputation(a)
  b = specialComputation(b)
  c = specialComputation(c)
  d = specialComputation(d)
  
  fmt.Println(a, b, c, d)


  var likes, shares int
  
  likes, shares = getLikesAndShares(4)
  
  if likes > 5 {
    fmt.Println("Woohoo! We got some likes.")
  }
  if shares > 10 {
    fmt.Println("We went viral!")
  }


  queryDatabase("SELECT * FROM coolTable;")
}