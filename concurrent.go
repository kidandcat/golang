package main

import "fmt"

func main(){
  c := make(chan int, 5)
  n := make(chan int, 5)
  n <- 1

  sum(c, n)
  sum(c, n)
  sum(c, n)
  sum(c, n)
  sum(c, n)
  sum(c, n)
  sum(c, n)
  sum(c, n)
  sum(c, n)
}

func sum(c chan int, n chan int){
  go plus1(n, c)
  res := <-c
  fmt.Println("Res: ", res)
}

func plus1(i chan int, c chan int){
  ii := <- i
  for s := 0 ; s < 1 ; s++ {
    ii++;
  }
  c <- ii
  i <- ii
}
