package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
	"time"
)

var dataFlag = flag.String("data", "", "File to use as data source")
var targetFlag = flag.String("file", "", "File to compress")
var file1 []byte
var file1Size int
var totalGlobalTotal int
var totalGlobal int
var final []Block

type Block struct {
	Index    int
	Position int
	Size     int
}

func main() {
	flag.Parse()
	readToMemory(*dataFlag)
	time.Sleep(1 * time.Second)

	f, size := openFile(*targetFlag)
	defer func() { f.Close() }()

	final = make([]Block, size)
	totalGlobal = size
	totalGlobalTotal = size
	t := time.Now()

	read(f, 3)
	read(f, 2)

	ii := 0
	for _, x := range final {
		if isValid(x) {
			final[ii] = x
			ii++
		}
	}
	final = final[:ii]
	b, _ := json.Marshal(final)
	ioutil.WriteFile("output.json", []byte(b), 0644)
	fmt.Println("End", time.Since(t))
}

func isValid(e Block) bool {
	if e.Size != 0 {
		return true
	}
	return false
}

func readToMemory(file string) {
	t1 := time.Now()
	f, size := openFile(file)
	defer func() { f.Close() }()
	r := bufio.NewReader(f)

	file1Size = size

	for byt, err := r.ReadByte(); err == nil; {
		file1 = append(file1, byt)
		byt, err = r.ReadByte()
	}

	fmt.Println("File loaded to memory\nFile size:", size, "\nMemory size:", len(file1), "\nTime:", time.Since(t1))
}

func read(f *os.File, step int) bool {
	r := bufio.NewReader(f)

	var aBuf []byte

	var wg sync.WaitGroup
	c := 0

	for {
		b, err := r.ReadByte()
		aBuf = append(aBuf, b)
		if err != nil {
			break
		}
		aaux, err := r.Peek(step - 1)
		aBuf = append(aBuf, aaux...)
		if err != nil {
			break
		}

		wg.Add(1)
		go handleStep(aBuf, c, step, &wg, final)
		aBuf = make([]byte, 0)
		c++
	}

	wg.Wait()

	return true
}

func handleStep(aBuf []byte, c int, step int, wg *sync.WaitGroup, final []Block) {
	var aux []byte
	for k, _ := range file1 {
		aux = make([]byte, 0)
		for i := 0; i < step; i++ {
			if k+i < file1Size {
				aux = append(aux, file1[k+i])
			} else {
				break
			}
		}
		if compare(aBuf, aux) {
			jump := false
			for i := 0; i < step; i++ {
				if final[c+i].Size != 0 {
					jump = true
				}
			}
			if jump {
				break
			}
			final[c] = Block{
				Index:    c,
				Position: k,
				Size:     step,
			}
			break
		}
	}
	print()
	wg.Done()
}

func compare(a, b []byte) bool {
	if &a == &b {
		return true
	}
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

func print() {
	fmt.Println(totalGlobal, "of", totalGlobalTotal)
	totalGlobal--
}

func openFile(f string) (*os.File, int) {
	fl, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	st, _ := fl.Stat()
	size := st.Size()
	return fl, int(size)
}

func remove(s []int, i int) []int {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}
