package util

import (
	"fmt"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/qingcc/goblog/util"
	"image/png"
	"os"
	"time"
)

type GoogleAuthInfo struct {
	Account string //账户
	Secret  string //密钥
	Title   string //厂商
	Path    string
}

func CreateGoogleAuthQr(info GoogleAuthInfo) GoogleAuthInfo {
	info.Path = "./uploads/google_authenticator/" + time.Now().Format("2006/0102/")
	util.DirectoryMkdir(info.Path)
	info.Path = info.Path + time.Now().Format("20060102_") + info.Secret + ".png"
	//如果存在这个文件对于的图片
	if res, _ := util.DirectoryExists(info.Path); res == true {
		return info
	}

	url := "otpauth://totp/" + info.Account + "?secret=" + info.Secret + "&issuer=" + info.Title
	// Create the barcode
	qrCode, _ := qr.Encode(url, qr.M, qr.Auto)

	// Scale the barcode to 200x200 pixels
	qrCode, _ = barcode.Scale(qrCode, 200, 200)

	// create the output file
	file, _ := os.Create(info.Path)
	defer file.Close()
	png.Encode(file, qrCode)
	info.Path = "/" + info.Path
	return info
}
func VerificationGoogleAuthCode(randkey string, code string) bool {
	//判断google验证码与code是否一致
	result := ReturnCode(randkey)
	//result需要与code类型一致
	if code == fmt.Sprint(result) {
		return true
	} else {
		return false
	}
}
