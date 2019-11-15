package utils

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
)

/*
RSA 公钥加密私钥解密 私钥签名公钥验证
*/

func test() {
	var theMsg = "the message you want to encode 你好 世界"
	fmt.Println("Source:", theMsg)
	//私钥签名
	sig, _ := RsaSign([]byte(theMsg))
	fmt.Println(string(sig))
	//公钥验证
	fmt.Println(RsaSignVer([]byte(theMsg), sig))
	//公钥加密
	fmt.Println("---------------")
	enc, _ := RsaEncrypt([]byte(theMsg))
	fmt.Println("Encrypted:", string(enc))
	//私钥解密
	decstr, _ := RsaDecrypt(enc)
	fmt.Println("Decrypted:", string(decstr))
}

func NewKey() bool {
	var bits int
	flag.IntVar(&bits, "b", 1024, "秘钥长度，默认为1024")
	if err := GenRsaKey(bits); err != nil {
		log.Fatal("秘钥文件生成失败")
		return false
	}
	log.Println("秘钥文件生成成功")
	return true
}

//生成 私钥和公钥文件
func GenRsaKey(bits int) error {
	//生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	file, err := os.Create("private.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	//生成公钥文件
	publicKey := &privateKey.PublicKey
	defPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	block = &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: defPkix,
	}
	file, err = os.Create("public.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	return nil
}

//公钥加密私钥解密  私钥签名公钥验证

var privateKey, publicKey []byte

//func init() {
//	var err error
//	publicKey, err = ioutil.ReadFile("public.pem")
//	if err != nil {
//		os.Exit(-1)
//	}
//	privateKey, err = ioutil.ReadFile("private.pem")
//	if err != nil {
//		os.Exit(-1)
//	}
//	fmt.Printf("%s\n", publicKey)
//	fmt.Printf("%s\n", privateKey)
//}

//私钥签名
func RsaSign(data []byte) ([]byte, error) {
	h := sha256.New()
	h.Write(data)
	hashed := h.Sum(nil)
	//获取私钥
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.SignPKCS1v15(rand.Reader, priv, crypto.SHA256, hashed)
}

//公钥验证
func RsaSignVer(data []byte, signature []byte) error {
	hashed := sha256.Sum256(data)
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//验证签名
	return rsa.VerifyPKCS1v15(pub, crypto.SHA256, hashed[:], signature)
}

// 公钥加密
func RsaEncrypt(data []byte) ([]byte, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	return rsa.EncryptPKCS1v15(rand.Reader, pub, data)
}

// 私钥解密
func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	//获取私钥
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}
