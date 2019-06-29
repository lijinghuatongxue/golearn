package main

import "fmt"

func main()  {
i := 0
for { //since there are no checks, this is an infinite loop
if i >= 3 {
	break
}
//break out of this for loop when this condition is met
fmt.Println("Value of i is:", i)
i++
}
fmt.Println("A statement just after for loop.")
}