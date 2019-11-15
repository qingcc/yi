package utils

/*
AES CFB加解密
*/
import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

//AES CFB解密
func AesDecrypter(en_str string) string {
	key, _ := hex.DecodeString("6368616e67652074686973207061a373")
	ciphertext, _ := hex.DecodeString(en_str)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)
	return fmt.Sprintf("%s", ciphertext)
}

//AES CFB加密
func AesEncrypter(str string) string {
	key, _ := hex.DecodeString("6368616e67652074686973207061a373")
	plaintext := []byte(str)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	return fmt.Sprintf("%x", ciphertext)
}
