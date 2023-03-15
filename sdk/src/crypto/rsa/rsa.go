package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func main() {
	privateKey, publicKey := genKey()

	rsaEncrypt := RSAEncrypt("hello, world", publicKey)
	rsaDecrypt := RSADecrypt(rsaEncrypt, privateKey)
	fmt.Println(rsaEncrypt, rsaDecrypt)

	key := "key"
	sign := RSASign(key, privateKey)
	verify := RSAVerify(key, sign, publicKey)
	fmt.Println(sign, verify)
}

func genKey() (*rsa.PrivateKey, *rsa.PublicKey) {
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	return privateKey, &privateKey.PublicKey
}

func RSAEncrypt(originData string, publicKey *rsa.PublicKey) string {
	originBytes := []byte(originData)
	encryptBytes, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, originBytes, nil)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(encryptBytes)
}

func RSADecrypt(encryptData string, privateKey *rsa.PrivateKey) string {
	encryptBytes, _ := base64.StdEncoding.DecodeString(encryptData)
	decryptBytes, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, encryptBytes, nil)
	if err != nil {
		panic(err)
	}
	return string(decryptBytes)
}

func RSASign(originData string, privateKey *rsa.PrivateKey) string {
	shaBytes := sha256.Sum256([]byte(originData))
	signBytes, err := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, shaBytes[:], nil)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(signBytes)
}

func RSAVerify(originData, sign string, publicKey *rsa.PublicKey) bool {
	shaBytes := sha256.Sum256([]byte(originData))
	signBytes, err := base64.StdEncoding.DecodeString(sign)
	err = rsa.VerifyPSS(publicKey, crypto.SHA256, shaBytes[:], signBytes, nil)
	if err != nil {
		return false
	}
	return true
}
