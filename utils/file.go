package utils

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/tealeg/xlsx"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

//region Remark: 创建文件夹 Author; Qing
func DirectoryMkdir(path string) {
	if res, _ := DirectoryExists(path); res == false {
		err := os.MkdirAll(path, os.ModePerm)
		_, file, line, _ := runtime.Caller(0) // 获取错误文件和错误行
		println(file+":"+strconv.Itoa(line), err)
	}
}

//endregion

//region Remark: 判断文件夹是否存在 Author:Qing
func DirectoryExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//endregion

//region Remark: 文件或目录是否存在 Author:Qing
func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

//endregion

//region Remark: 列出指定路径的文件和目录 Author:Qing
func ScanDir(dir string) []string {
	file, err := os.Open(dir)
	if err != nil {
		return []string{}
	}
	names, err := file.Readdirnames(-1)
	if err != nil {
		return []string{}
	}
	return names
}

// 列举一个目录（包括子目录）下的xlsx文件
func ListExcelFile(dir string) (result []string) {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".xlsx") {
			result = append(result, path)
		}
		return nil
	})
	if err != nil {
		log.Println(err)
	}
	return result
}

//endregion

func Tracefile(str_content string, file string) {
	fd, _ := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)

	fd_content := strings.Join([]string{str_content, "\n"}, "")
	buf := []byte(fd_content)
	fd.Write(buf)
	fd.Close()
}

//读取xlsx文件
func readfile(file string) {
	// 打开文件
	xlFile, err := xlsx.OpenFile(file)
	if err != nil {
		log.Println("打开文件失败！" + err.Error())
		return
	}
	if len(xlFile.Sheets) == 0 {
		log.Println("没有找到工作表")
	}

	unMap := make(map[string]bool)
	leftMap := make(map[string]string)
	for k, row := range xlFile.Sheets[0].Rows {
		if k == 0 {
			continue
		}
		countryname := row.Cells[2].String()
		//cen := strings.ToLower(row.Cells[3].String())
		if _, ok := unMap[countryname]; !ok {
			//if _, b := Country2Code[cen]; !b {
			//	leftMap[countryname] = cen
			//	continue
			//}
			//str := "\"" + countryname + "\"" + ":\"" + Country2Code[cen] + "\","
			//utils.Tracefile(str, "log.log")
			unMap[countryname] = true
		}
	}
	for cn, val := range leftMap {
		println(cn, " : ", val)
	}
	fmt.Println("\n\nimport success")

}

//以追加的方式写入文件，当写入文件的数据太多时，可以使用分批追加写入
func AppendCsv(filename string, isFirst bool, colsColumns []string, data [][]string) {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm) //os.O_TRUNC执行之前是否需要清空原有数据
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF")
	w := csv.NewWriter(f)
	if isFirst {
		w.Write(colsColumns)
	}
	w.WriteAll(data)
	w.Flush()
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
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
		if err != nil || io.EOF == err {
			break
		}
		lines = append(lines, line[:len(line)-1])
	}
	return
}

// endregion
