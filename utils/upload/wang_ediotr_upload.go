package util

import (
	"github.com/gin-gonic/gin"
	"github.com/qingcc/goblog/logic"
	"github.com/qingcc/goblog/util"
	tsgutils "github.com/typa01/go-utils"
	"io"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

//region Remark: 上传 Author:Qing
func WangEditorUpload(c *gin.Context) {
	c.JSON(http.StatusOK, wang_editor_upload(c, "FileData", "editor", "jpg,png"))
}

//endregion

//region Remark: 保存上传的图片 Author:Qing
func wang_editor_upload(c *gin.Context, fileName string, fileType string, suffix string) gin.H {
	objLog := logic.GetLogger(c)
	err := c.Request.ParseMultipartForm(200000)
	if err != nil {
		return gin.H{"errno": 0, "msg": "上传失败"}
	}
	files := c.Request.MultipartForm.File[fileName]
	dataPath := make([]string, len(files))
	for i, _ := range files {
		file, err := files[i].Open()
		if err != nil {
			return gin.H{"errno": 0, "msg": "上传失败"}
			break
		}
		filename := strings.Split(files[i].Filename, ".")
		filename_suffix := filename[len(filename)-1]
		uid := tsgutils.GUID()
		new_filename := uid + "." + filename_suffix

		//判断文件后缀是否允许上传
		if !strings.Contains(suffix, filename_suffix) {
			return gin.H{"errno": 0, "msg": "文件后缀不允许"}
		}

		//创建文件夹
		path := "uploads/" + fileType + "/" + time.Now().Format("2006/0102/")
		util.DirectoryMkdir(path)

		//创建文件
		out, err := os.Create(new_filename)
		if err != nil {
			_, file, line, _ := runtime.Caller(0) //获取错误文件和错误行
			objLog.Errorf(file+":"+strconv.Itoa(line), "上传错误：%s", err)
		}
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			_, file, line, _ := runtime.Caller(0) //获取错误文件和错误行
			objLog.Errorf(file+":"+strconv.Itoa(line), "上传错误：%s", err)
		}
		dataPath[i] = "/" + path + new_filename
	}
	return gin.H{
		"errno": 1,
		"data":  dataPath,
	}
}

//endregion
