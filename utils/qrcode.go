package utils

import (
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"image/png"
	"log"
	"os"
)

func NewQrcode(filename, content string, wide int) bool {
	//code, err := qr.Encode(base64, qr.L, qr.Unicode)
	code, err := qr.Encode(content, qr.M, qr.Auto)
	// code, err := code39.Encode(base64)
	if err != nil {
		log.Fatal(err)
		return false
	}

	if content != code.Content() {
		log.Fatal("data differs")
		return false
	}

	code, err = barcode.Scale(code, wide, wide)
	if err != nil {
		log.Fatal(err)
		return false
	}

	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
		return false
	}
	err = png.Encode(file, code)
	// err = jpeg.Encode(file, img, &jpeg.Options{100})      //图像质量值为100，是最好的图像显示
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
