package util

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	"runtime"
)

/*
	Operation with rsa encryption
*/
func init() {
	log.SetFlags(log.Ldate | log.Lshortfile)
}

//RsaEncrypt: RSA加密，传入明文返回密文
func RsaEncrypt(plainText string) (cryptText []byte, err error) {
	block, _ := pem.Decode([]byte(publicKey))
	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case runtime.Error:
				log.Println("runtime err:", err, "Check that the key is correct")
			default:
				log.Println("error:", err)
			}
		}
	}()
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return []byte{}, err
	}
	publicKey := publicKeyInterface.(*rsa.PublicKey)

	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(plainText))
	if err != nil {
		return []byte{}, err
	}

	return cipherText, nil
}

//RsaDecrypt: 解密，传入密文，返回解密后的明文
func RsaDecrypt(cryptText []byte) (plainText string, err error) {
	block, _ := pem.Decode([]byte(privateKey))

	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case runtime.Error:
				log.Println("runtime err:", err, "Check that the key is correct")
			default:
				log.Println("error:", err)
			}
		}
	}()
	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	plain, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey.(*rsa.PrivateKey), cryptText)
	if err != nil {
		return "", err
	}
	plainText = string(plain)
	return plainText, nil
}
func Test() {

	plaintext := "test"
	// 直接传入明文和公钥加密得到密文
	crypttext, err := RsaEncrypt(plaintext)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("密文", hex.EncodeToString(crypttext))
	// 解密操作，直接传入密文和私钥解密操作，得到明文
	plaintext, err = RsaDecrypt(crypttext)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("明文：", plaintext)
}

const (
	privateFileName  = "private.pem"
	publicFileName   = "public.pem"
	privateKeyPrefix = "MICAH RSA PRIVATE KEY "
	publicKeyPrefix  = "MICAH RSA PUBLIC KEY "

	privateKey = `-----BEGIN MICAH RSA PRIVATE KEY -----
MIICeQIBADANBgkqhkiG9w0BAQEFAASCAmMwggJfAgEAAoGBAJnqeWNO/wvF2i3a
El+Gj9uv7mfRvNzg9Y7c6RKZyqLuZVZulSdVgl5CyiWDmORYGr9I6PoH4JL/c0IO
XAcKoASCNGwNOSrKTFgRtMQWTQNe7Gd0aS9OA0DITYsg26dqRyrclI02q+5WXLCg
YNlBSIW0azx9sMsv3Qd6SHRzwd8rAgMBAAECgYEAgjgdgUa5g6VeJQLmHonDNnPP
eWi6qAlv1/HRA1q4VUbq18hxSrkOtl89laWT+kUMRCAUdE2r09JqGk350D9OB5KE
AUAsSqFff+xJZOT/Qe6FJLnwGabU7SfUu+PqsJA94PuA5ZoBDHY1PJbPqZXYwlaa
GWAkT6YXWNjFepiZJMECQQDLJkEo6FXI/M4KDd4L/A76ioJUE1fsKRClbW0Gb55v
9zJ+N8i2Pt7V3a8brOqY01CW+An2LlAVwV/Wj6p/kzXLAkEAwfVE3vsHvQ6BztaJ
Q5lSepN54Iwe9dIP5+qmelCblRdkZVpk2e24j+6lfd5WeCvjaRo8mRMEh8fhF9yL
x9nQIQJBAIagQFhCt+7tEgH5tKgOj0Kbncjn1MFtaVmnzbORLtcWY4DfqQSVX4kt
UwktK37Bs0uI7tkVUf9I2iFPrfbDcMcCQQCOb5AkMDV+oSiyvNoQyofHfIYEYOJb
o+gKaEfZi1i2JHeV9swZPnpLOPZPePgHSDO4+4uE3nJ1RFQJMe101oABAkEAxhW5
lBH8/TYxevLwLJHJzIcoMkhJljgG1tbJJA44WIEwFbdnA92dOszo5HrGuYqF1c1E
XoXmxLfbqYqxGOrN0g==
-----END MICAH RSA PRIVATE KEY -----`
	publicKey = `-----BEGIN MICAH RSA PUBLIC KEY -----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCZ6nljTv8Lxdot2hJfho/br+5n
0bzc4PWO3OkSmcqi7mVWbpUnVYJeQsolg5jkWBq/SOj6B+CS/3NCDlwHCqAEgjRs
DTkqykxYEbTEFk0DXuxndGkvTgNAyE2LINunakcq3JSNNqvuVlywoGDZQUiFtGs8
fbDLL90Hekh0c8HfKwIDAQAB
-----END MICAH RSA PUBLIC KEY -----`
)

func GetRsaKey() error {
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		return err
	}
	x509PrivateKey, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		return err
	}
	privateFile, err := os.Create(privateFileName)
	if err != nil {
		return err
	}
	defer privateFile.Close()
	privateBlock := pem.Block{
		Type:  privateKeyPrefix,
		Bytes: x509PrivateKey,
	}

	if err = pem.Encode(privateFile, &privateBlock); err != nil {
		return err
	}
	publicKey := privateKey.PublicKey
	x509PublicKey, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		panic(err)
	}
	publicFile, _ := os.Create(publicFileName)
	defer publicFile.Close()
	publicBlock := pem.Block{
		Type:  publicKeyPrefix,
		Bytes: x509PublicKey,
	}
	if err = pem.Encode(publicFile, &publicBlock); err != nil {
		return err
	}
	return nil
}
