package main

import (
  "os"
  "fmt"
)

func main(){
  f, err := os.OpenFile("dadn.lib", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
  if err != nil {
    panic(err)
  }
  defer f.Close()

  aBytes := makeRange(0,255)
  for _,v1 := range aBytes {
    for _,v2 := range aBytes {
      for _,v3 := range aBytes {
        for _,v4 := range aBytes {
          for _,v5 := range aBytes {
            for _,v6 := range aBytes {
              fmt.Println("Writing", v1, v2, v3, v4, v5, v6)
              write([]byte{byte(v1), byte(v2), byte(v3), byte(v4), byte(v5), byte(v6)}, f)
            }
            f.Sync()
          }
        }
      }
    }
  }
}

func write(bs []byte, f *os.File){
    _, err := f.Write(bs)
    if err != nil {
      panic(err)
    }
}

func makeRange(min, max int) []int {
    a := make([]int, max-min+1)
    for i := range a {
        a[i] = min + i
    }
    return a
}
