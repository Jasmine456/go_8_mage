package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)

/*
生成1024位的RSA私钥：
openssl genrsa -out data/rsa_provate_key.pem 1024
根据私钥生成公钥：
openssl rsa -in data/rsa_provate_key.pem -pubout -out data/rsa_public_key.pem

pem是一种标准格式，它通常包含页眉和页脚
 */

var (
	publicKey []byte
	privateKey []byte
)

func ReadFile(keyFile string) ([]byte,error){
	if f,err := os.Open(keyFile);err != nil{
		return nil,err
	} else{
		content := make([]byte,4096)
		if n,err := f.Read(content);err !=nil{
			return nil,err
		} else{
			return content[:n],nil
		}
	}
}

func ReadRSAKey(publicKeyFile,privateKeyFile string) (err error){
	if publicKey,err =ReadFile(publicKeyFile);err !=nil{
		return err
	}
	if privateKey,err = ReadFile(privateKeyFile);err !=nil{
		return err
	}
	return
}

// RSA加密
func RsaEncrypt(origData []byte) ([]byte,error){
	//	解密pem格式的公钥
	block,_ := pem.Decode(publicKey)
	if block == nil{
		return nil,errors.New("public key error")
	}
//	解析公钥
	pubInterface,err := x509.ParsePKIXPublicKey(block.Bytes) //目前的数字证书一般都是基于ITU(国际电信联盟)制定的x.509标准
	if err != nil{
		return nil,err
	}
//	类型断言
	pub:= pubInterface.(*rsa.PublicKey)
//	加密
	return rsa.EncryptPKCS1v15(rand.Reader,pub,origData)
}

//	RSA解密
func RsaDecrypt(ciphertext []byte) ([]byte,error) {
//	解密
	block,_ := pem.Decode(privateKey)
	if block == nil{
		return nil,errors.New("provate key error!")
	}
	// 解析PKCS1格式的私钥
	privInf,err :=x509.ParsePKCS8PrivateKey(block.Bytes)
	if err!=nil{
		return nil,err
	}
	priv := privInf.(*rsa.PrivateKey)
//	解密
	return rsa.DecryptPKCS1v15(rand.Reader,priv,ciphertext)
}

func main(){
	ReadRSAKey("../data/rsa_public_key.pem","../data/rsa_private_key.pem")

	plain := "因为我们没有什么不同"
	cipher,err := RsaEncrypt([]byte(plain))
	if err != nil{
		fmt.Println(err)
	} else{
		fmt.Printf("密文：%v\n",cipher)
		bPlain,err := RsaDecrypt(cipher)
		if err !=nil{
			fmt.Println(err)
		} else{
			fmt.Printf("明文：%s\n",string(bPlain))
		}
	}
}