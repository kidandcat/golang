package main

import (
  "fmt"
  "bufio"
  "os"
)

func main(){
  f, err := os.Open("index.html")
  defer func(){
    f.Close()
  }()
  if err != nil {
    panic(err)
  }

  r := bufio.NewReader(f)
  bytes1, _ := r.ReadByte()
  bytes2, _ := r.ReadByte()
  fmt.Println("Bytes: ", string(bytes1), string(bytes2))

  // s := bufio.NewScanner(r)
  //
  // for s.Scan() {
  //   fmt.Println(s.Text())
  // }
  read("index.html", 1)
}

func read(file string, step int){
  f, size := openFile(file)
  defer func(){f.Close()}()
  r := bufio.NewReader(f)


  for bytes, err := r.ReadByte(); err == nil; {
    for i := 0 ; i < step ; i++{
      bytes, err = r.ReadByte()
    }
    fmt.Println("Byte: ", bytes)
  }

  fmt.Println("End ")
}

func openFile(f string) (*os.File, int){
  fl, err := os.Open(f)
  if err != nil {
    panic(err)
  }
  st, _ := fl.Stat()
  size := st.Size()
  return fl, int(size)
}
