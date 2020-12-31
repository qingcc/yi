package main

import (
	"bufio"
	"github.com/qingcc/yi/utils"
	"github.com/sirupsen/logrus"
	"io"
	"math"
	"os"
	"time"
)

func ReadLineByLine(filename string) {
	f, err := os.Open(filename)
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

func main() {
	columns := []string{"id", "name"}
	data := [][]string{
		[]string{"1", "a"},
		[]string{"2", "b"},
		[]string{"3", "c"},
		[]string{"4", "d"},
	}
	utils.AppendCsv("test.csv", true, columns, data)
	data = [][]string{
		[]string{"5", "e"},
		[]string{"6", "f"},
		[]string{"7", "g"},
		[]string{"8", "h"},
	}
	utils.AppendCsv("test.csv", false, columns, data)
	return
	t := time.Now().AddDate(0, 0, 5)
	can := t.Add(time.Hour * 8)
	logrus.Info(t.Sub(can).Hours())
	nn := Decimal(3.5/100*135.00 + 135.00)
	logrus.Info("nn:", nn, " nn*100:", nn*100, int(nn*100))
	tr := Decimal(nn * 100)
	logrus.Info(int(tr))
}

func Decimal(value float64) float64 {
	// 只去浮点数的小数点后两位
	return math.Trunc(value*1e2+0.5) * 1e-2
}
