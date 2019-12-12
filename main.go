package main

import (
	"bufio"
	"io"
	"os"
)

func main()  {
	f, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
		if err != nil || io.EOF == err {
			break
		}
		m[line] = 1
	}

}
var m map[string]int

