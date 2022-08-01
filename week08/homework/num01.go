package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

//读取密钥文件
func ReadFile(keyFile string) ([]byte, error) {
	res, err := os.Open(keyFile)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Close()
	stat, _ := res.Stat()
	context := make([]byte, stat.Size())
	res.Read(context)
	return context, nil
}

//hash计算
func dHash(str string) []byte {
	h := sha512.New()
	h.Write([]byte(str))
	hValue := h.Sum(nil)
	return hValue
}

//私钥做数字签名，返回数字签名[]byte
func DigitalSinature(str string) ([]byte, error) {
	//读取私钥
	key, err := ReadFile("rsa/rsa_private_key.pem")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	//将内容解析成 pem 格式块
	block, _ := pem.Decode(key)
	//将 pem 块中的 PKCS # 1, ASN.1 DER 格式的字符串解析成 rsa.PrivateKey
	privitekey, err := x509.ParsePKCS1PrivateKey(block.Bytes)

	//计算原始内容的散列值
	hValue := dHash(str)

	//通过 rsa.SignPKCS1v15 使用私钥对散列值进行签名
	digestSign, err := rsa.SignPKCS1v15(rand.Reader, privitekey, crypto.SHA512, hValue)
	return digestSign, nil
}

//公钥做数字签名认证,返回bool
func VerifySinature(str string, signal []byte) bool {
	//读取公钥
	key, err := ReadFile("rsa/rsa_public_key.pem")
	if err != nil {
		fmt.Println(err)
		return false
	}
	//读取公钥文件内容，使用 pem.Decode 将内容解析成 pem 格式块
	block, _ := pem.Decode(key)
	//通过 x509.ParsePKIXPublicKey 将 pem 块中的 DER 格式字符串解析成 rsa.PublicKey
	publickeyInt, err := x509.ParsePKIXPublicKey(block.Bytes)
	publickey := publickeyInt.(*rsa.PublicKey)
	//hash计算
	hValue := dHash(str)

	//通过 rsa.VerifyPKCS1v15 使用公钥对散列值进行认证，如果该返回返回 err == nil，则表示认证成功
	err = rsa.VerifyPKCS1v15(publickey, crypto.SHA512, hValue, signal)

	return err == nil

}

func main() {
	context := "我们都是中国人！！！"
	digestSign, _ := DigitalSinature(context)
	fmt.Println("数字签名：", digestSign)
	res := VerifySinature(context, digestSign)
	fmt.Println("签名认证：", res)

}
