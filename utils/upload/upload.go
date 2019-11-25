package upload

import (
	"github.com/qingcc/yi/utils"
	tsgutils "github.com/typa01/go-utils"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

//region Remark: 上传 Author:Qing
func UploadImage(w http.ResponseWriter, r *http.Request) {
	upload(w, r, "images", "bmp,gif,jpg,jpeg,jpe,png")
}

//endregion

//region Remark: 上传 Author:Qing
func UploadFile(w http.ResponseWriter, r *http.Request) {
	upload(w, r, "file", "zip,rar,pdf,apk")
}

//endregion

//region Remark: 上传 Author:Qing
func UploadVideo(w http.ResponseWriter, r *http.Request) {
	upload(w, r, "video", "mp4")
}

//endregion

//region Remark: 保存上传的文件 Author:Qing
func upload(w http.ResponseWriter, r *http.Request, fileType string, suffix string) {

	//得到上传的文件
	file, header, err := r.FormFile("FileData") //image这个是uplaodify参数定义中的   'fileObjName':'image'
	if err != nil {
		io.WriteString(w, "upload failed")
		return
	}

	filename := strings.Split(header.Filename, ".")
	filename_suffix := filename[len(filename)-1]
	uid := tsgutils.GUID()
	new_filename := uid + "." + filename_suffix

	//判断文件后缀是否允许上传
	if !strings.Contains(suffix, filename_suffix) {
		io.WriteString(w, "上传格式不允许，只允许上传上传："+suffix)
		return
	}

	//创建文件夹
	path := "uploads/" + fileType + "/" + time.Now().Format("2006/0102/")
	utils.DirectoryMkdir(path)

	//创建文件
	out, err := os.Create(path + new_filename)
	if err != nil {
		_, file, line, _ := runtime.Caller(0) //获取错误文件和错误行
		log.Printf(file+":"+strconv.Itoa(line), "上传错误：%s", err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		_, file, line, _ := runtime.Caller(0) //获取错误文件和错误行
		log.Printf(file+":"+strconv.Itoa(line), "上传错误：%s", err)
	}
	//imgHost := "http://" + c.Request.Host
	//返回值
	io.WriteString(w, "上传成功")
	return
	//{
	//	"status": config.HttpError,
	//	"name":   header.Filename,
	//	"msg":    "上传成功",
	//	"size":   header.Size,
	//	"data":   "/" + path + new_filename,
	//	"url":    imgHost + "/" + path + new_filename,
	//	"path":   "/" + path + new_filename,
	//}
}

//endregion
