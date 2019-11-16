package upload

import (
	"encoding/json"
	"github.com/qingcc/yi/utils"
	tsgutils "github.com/typa01/go-utils"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

//region Remark: 上传 Author:Qing
func WangEditorUpload(w http.ResponseWriter, r *http.Request) {
	wang_editor_upload(w, r, "FileData", "editor", "jpg,png")
}

//endregion

//region Remark: 保存上传的图片 Author:Qing
func wang_editor_upload(w http.ResponseWriter, r *http.Request, fileName string, fileType string, suffix string) {

	err := c.Request.ParseMultipartForm(200000)
	if err != nil {
		io.WriteString(w, "获取失败")
		return
	}
	files := r.MultipartForm.File[fileName]
	dataPath := make([]string, len(files))
	for i, _ := range files {
		file, err := files[i].Open()
		if err != nil {
			io.WriteString(w, "上传失败")
			return
		}
		filename := strings.Split(files[i].Filename, ".")
		filename_suffix := filename[len(filename)-1]
		uid := tsgutils.GUID()
		new_filename := uid + "." + filename_suffix

		//判断文件后缀是否允许上传
		if !strings.Contains(suffix, filename_suffix) {
			io.WriteString(w, "文件后缀不允许")
			return
		}

		//创建文件夹
		path := "uploads/" + fileType + "/" + time.Now().Format("2006/0102/")
		utils.DirectoryMkdir(path)

		//创建文件
		out, err := os.Create(new_filename)
		if err != nil {
			io.WriteString(w, "上传错误：%s"+err.Error())
			return
		}
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			io.WriteString(w, "上传错误：%s"+err.Error())
		}
		dataPath[i] = "/" + path + new_filename
	}
	dataPathB, _ := json.Marshal(dataPath)
	io.WriteString(w, string(dataPathB))
	return
}

//endregion
