package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

var (
	filename   = flag.String("f", "./tools/readfile2String/aim.log", "file name")
	prefix     = flag.String("prefix", "", "prefix")
	suffix     = flag.String("suffix", "", "suffix")
	delimiter  = flag.String("delimiter", ",", "delimiter") // 分隔符
	quoteMarks = flag.String("quote", "", "delimiter")      // 引号
)

func main() {
	flag.Parse()
	if *filename != "" {
		lines := ReadLineByLine(*filename)
		if *prefix != "" || *suffix != "" {
			tempLines := make([]string, 0, len(lines))
			for _, v := range lines {
				tempLines = append(tempLines, fmt.Sprintf("%s%s%s", *prefix, v, *suffix))
			}
			lines = tempLines
		}
		buf := new(bytes.Buffer)
		for _, line := range lines {
			buf.WriteString(*quoteMarks)
			buf.WriteString(line)
			buf.WriteString(*quoteMarks)
			buf.WriteString(*delimiter)
		}
		logrus.Infof("len:%d, output:%s", len(lines), buf.String())
	}
}

// region read file by line
func ReadLineByLine(filename string) (lines []string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n') // 以'\n'为结束符读入一行
		if err != nil || io.EOF == err {
			break
		}
		lines = append(lines, line[:len(line)-1]) // 去掉 '/n' 字符
	}
	return
}

// endregion
