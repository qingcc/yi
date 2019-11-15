package utils

import (
	"encoding/csv"
	"os"
	"runtime"
	"strconv"
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

//endregion

//region Remark: 写入文件 Author:Qing
func Write2File(fileName string, data [][]string) {
	f, err := os.Create(fileName) //创建文件 test.csv
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM

	w := csv.NewWriter(f) //创建一个新的写入文件流
	//data := [][]string{
	//	{"1", "中国", "23"},
	//	{"2", "美国", "23"},
	//	{"3", "bb", "23"},
	//	{"4", "bb", "23"},
	//	{"5", "bb", "23"},
	//}
	w.WriteAll(data) //写入数据
	w.Flush()
}

//endregion
